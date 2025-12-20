package executor

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

const (
	// Default compilation image for building Go binaries
	defaultCompileImage = "golang:1.25-alpine"
	// Default execution image for running compiled binaries
	defaultExecImage = "alpine:latest"
	// Workspace directory inside containers
	containerWorkspace = "/workspace"
	// Docker connection timeout
	dockerConnectionTimeout = 5 * time.Second
	// CPU period in microseconds (100ms)
	cpuPeriodMicroseconds = 100000
	// CPU percentage denominator
	cpuPercentDenominator = 100
	// Bytes per KB
	bytesPerKB = 1024
	// Binary file permissions
	binaryFileMode = 0o755
)

// dockerExecutor handles execution of Go code using Docker containers.
type dockerExecutor struct {
	client        *client.Client
	compileImage  string
	execImage     string
	maxMemoryMB   int
	maxCPUPercent int
	maxOutput     int
	timeout       time.Duration
	logger        *slog.Logger
}

// newDockerExecutor creates a new Docker-based executor.
func newDockerExecutor(
	compileImage, execImage string,
	maxMemoryMB, maxCPUPercent, maxOutput int,
	timeout time.Duration,
	logger *slog.Logger,
) (*dockerExecutor, error) {
	// Initialize Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDockerNotAvailable, err)
	}

	// Test Docker connection
	pingCtx, pingCancel := context.WithTimeout(context.Background(), dockerConnectionTimeout)
	defer pingCancel()

	if _, pingErr := cli.Ping(pingCtx); pingErr != nil {
		return nil, fmt.Errorf("%w: %w", ErrDockerNotAvailable, pingErr)
	}

	executor := &dockerExecutor{
		client:        cli,
		compileImage:  compileImage,
		execImage:     execImage,
		maxMemoryMB:   maxMemoryMB,
		maxCPUPercent: maxCPUPercent,
		maxOutput:     maxOutput,
		timeout:       timeout,
		logger:        logger,
	}

	return executor, nil
}

// execute runs Go code in a Docker container using two-stage execution.
func (de *dockerExecutor) execute(ctx context.Context, code string) (*ExecutionResult, error) {
	startTime := time.Now()

	// Create a temporary directory for compilation artifacts
	tempDir, err := os.MkdirTemp("", "docker-exec-*")
	if err != nil {
		return nil, fmt.Errorf("create temp directory: %w", err)
	}
	defer func() {
		if cleanupErr := os.RemoveAll(tempDir); cleanupErr != nil {
			de.logger.Error("failed to cleanup temp directory", "error", cleanupErr, "dir", tempDir)
		}
	}()

	// Write code to temporary file
	codeFile := filepath.Join(tempDir, "code.go")
	if writeErr := os.WriteFile(codeFile, []byte(code), 0o600); writeErr != nil {
		return nil, fmt.Errorf("write code file: %w", writeErr)
	}

	// Stage 1: Compile the code
	binaryPath, err := de.compileCode(ctx, tempDir)
	if err != nil {
		return &ExecutionResult{
			Output:   "",
			Error:    err.Error(),
			ExitCode: -1,
			Duration: time.Since(startTime).String(),
		}, nil
	}

	// Stage 2: Execute the compiled binary
	result, execErr := de.executeBinary(ctx, binaryPath)
	if execErr != nil {
		return nil, execErr
	}

	result.Duration = time.Since(startTime).String()
	return result, nil
}

// compileCode compiles Go code to a binary using a Docker container.
func (de *dockerExecutor) compileCode(ctx context.Context, tempDir string) (string, error) {
	// Use parent context directly (timeout already applied)
	compileCtx := ctx

	// Create container for compilation
	containerConfig := &container.Config{
		Image: de.compileImage,
		Env:   []string{"CGO_ENABLED=0"}, // Disable CGO for static binary
		Cmd: []string{
			"sh", "-c",
			fmt.Sprintf("go build -o %s %s",
				filepath.Join(containerWorkspace, "binary"),
				filepath.Join(containerWorkspace, "code.go")),
		},
		WorkingDir: containerWorkspace,
	}

	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: tempDir,
				Target: containerWorkspace,
			},
		},
		AutoRemove: true,
	}

	resp, createErr := de.client.ContainerCreate(compileCtx, containerConfig, hostConfig, nil, nil, "")
	if createErr != nil {
		return "", fmt.Errorf("%w: %w", ErrCompilationFailed, createErr)
	}

	containerID := resp.ID

	// Ensure container is removed
	defer func() {
		removeCtx, removeCancel := context.WithTimeout(context.Background(), dockerConnectionTimeout)
		defer removeCancel()
		if removeErr := de.client.ContainerRemove(removeCtx, containerID, container.RemoveOptions{Force: true}); removeErr != nil {
			de.logger.WarnContext(removeCtx, "failed to remove compilation container", "error", removeErr, "container", containerID)
		}
	}()

	// Start container
	if startErr := de.client.ContainerStart(compileCtx, containerID, container.StartOptions{}); startErr != nil {
		return "", fmt.Errorf("%w: start container: %w", ErrCompilationFailed, startErr)
	}

	// Wait for container to finish
	statusCh, errCh := de.client.ContainerWait(compileCtx, containerID, container.WaitConditionNotRunning)

	select {
	case waitErr := <-errCh:
		if waitErr != nil {
			return "", fmt.Errorf("%w: wait container: %w", ErrCompilationFailed, waitErr)
		}
	case status := <-statusCh:
		if status.StatusCode != 0 {
			// Get stderr from container to return compilation errors
			logs, logErr := de.getContainerLogs(compileCtx, containerID)
			if logErr != nil {
				return "", fmt.Errorf("%w: exit code %d", ErrCompilationFailed, status.StatusCode)
			}
			return "", fmt.Errorf("%w: %s", ErrCompilationFailed, logs)
		}
	case <-compileCtx.Done():
		return "", fmt.Errorf("%w: compilation timeout", ErrTimeout)
	}

	// Binary should now exist in tempDir
	binaryPath := filepath.Join(tempDir, "binary")
	if _, statErr := os.Stat(binaryPath); statErr != nil {
		return "", fmt.Errorf("%w: binary not found after compilation", ErrCompilationFailed)
	}

	return binaryPath, nil
}

// executeBinary executes a compiled binary in a minimal Docker container.
func (de *dockerExecutor) executeBinary(ctx context.Context, binaryPath string) (*ExecutionResult, error) {
	// Use parent context directly (timeout already applied)
	execCtx := ctx

	// Read binary into memory to copy into container
	binaryData, err := os.ReadFile(binaryPath)
	if err != nil {
		return nil, fmt.Errorf("read binary: %w", err)
	}

	// Create container for execution
	containerConfig := &container.Config{
		Image:      de.execImage,
		Cmd:        []string{"/binary"},
		WorkingDir: "/",
	}

	// Calculate CPU quota (CPUPercent * CPUPeriod / 100)
	cpuPeriod := int64(cpuPeriodMicroseconds)
	cpuQuota := int64(de.maxCPUPercent) * cpuPeriod / cpuPercentDenominator
	memoryBytes := int64(de.maxMemoryMB) * bytesPerKB * bytesPerKB

	hostConfig := &container.HostConfig{
		Resources: container.Resources{
			Memory:    memoryBytes,
			CPUQuota:  cpuQuota,
			CPUPeriod: cpuPeriod,
		},
		AutoRemove:  true,
		NetworkMode: container.NetworkMode("none"), // No network access
	}

	resp, err := de.client.ContainerCreate(execCtx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("%w: create container: %w", ErrContainerExecution, err)
	}

	containerID := resp.ID

	// Ensure container is removed
	defer func() {
		removeCtx, removeCancel := context.WithTimeout(context.Background(), dockerConnectionTimeout)
		defer removeCancel()
		if removeErr := de.client.ContainerRemove(removeCtx, containerID, container.RemoveOptions{Force: true}); removeErr != nil {
			de.logger.WarnContext(removeCtx, "failed to remove execution container", "error", removeErr, "container", containerID)
		}
	}()

	// Copy binary into container
	if copyErr := de.copyToContainer(execCtx, containerID, binaryData); copyErr != nil {
		return nil, fmt.Errorf("%w: copy binary: %w", ErrContainerExecution, copyErr)
	}

	// Start container
	if startErr := de.client.ContainerStart(execCtx, containerID, container.StartOptions{}); startErr != nil {
		return nil, fmt.Errorf("%w: start container: %w", ErrContainerExecution, startErr)
	}

	// Wait for container to finish
	statusCh, errCh := de.client.ContainerWait(execCtx, containerID, container.WaitConditionNotRunning)

	var exitCode int
	select {
	case waitErr := <-errCh:
		if waitErr != nil {
			if execCtx.Err() == context.DeadlineExceeded {
				return nil, fmt.Errorf("%w", ErrTimeout)
			}
			return nil, fmt.Errorf("%w: wait container: %w", ErrContainerExecution, waitErr)
		}
	case status := <-statusCh:
		exitCode = int(status.StatusCode)
	case <-execCtx.Done():
		// Timeout - kill container
		killCtx, killCancel := context.WithTimeout(context.Background(), dockerConnectionTimeout)
		defer killCancel()
		_ = de.client.ContainerKill(killCtx, containerID, "SIGKILL")
		return nil, fmt.Errorf("%w", ErrTimeout)
	}

	// Get container logs (stdout + stderr)
	output, logErr := de.getContainerLogs(execCtx, containerID)
	if logErr != nil {
		de.logger.WarnContext(execCtx, "failed to get container logs", "error", logErr)
		output = ""
	}

	// Truncate output if needed
	output = de.truncateOutput(output)

	result := &ExecutionResult{
		ExitCode: exitCode,
		Output:   output,
	}

	if exitCode != 0 {
		result.Error = output
		result.Output = ""
	}

	return result, nil
}

// copyToContainer copies binary data into a container.
func (de *dockerExecutor) copyToContainer(ctx context.Context, containerID string, data []byte) error {
	// Create a tar archive containing the binary
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	header := &tar.Header{
		Name: "/binary",
		Mode: binaryFileMode,
		Size: int64(len(data)),
	}

	if err := tw.WriteHeader(header); err != nil {
		return fmt.Errorf("write tar header: %w", err)
	}

	if _, err := tw.Write(data); err != nil {
		return fmt.Errorf("write tar data: %w", err)
	}

	if err := tw.Close(); err != nil {
		return fmt.Errorf("close tar writer: %w", err)
	}

	// Copy tar archive into container
	return de.client.CopyToContainer(ctx, containerID, "/", &buf, container.CopyToContainerOptions{})
}

// getContainerLogs retrieves stdout and stderr from a container.
func (de *dockerExecutor) getContainerLogs(ctx context.Context, containerID string) (string, error) {
	reader, logErr := de.client.ContainerLogs(ctx, containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if logErr != nil {
		return "", logErr
	}
	defer reader.Close()

	var output bytes.Buffer
	if _, copyErr := io.Copy(&output, reader); copyErr != nil {
		return "", copyErr
	}

	// Docker logs are prefixed with [0-3] bytes indicating stream type (stdout/stderr)
	// We need to strip these prefixes
	outputBytes := output.Bytes()
	if len(outputBytes) > 0 {
		// Skip the stream type byte(s) if present
		start := 0
		for start < len(outputBytes) && (outputBytes[start] == 0 || outputBytes[start] == 1 || outputBytes[start] == 2) {
			start++
		}
		return string(outputBytes[start:]), nil
	}

	return "", nil
}

// truncateOutput truncates output if it exceeds maximum size.
func (de *dockerExecutor) truncateOutput(output string) string {
	if len(output) <= de.maxOutput {
		return output
	}
	return output[:de.maxOutput] + "\n... (output truncated)"
}
