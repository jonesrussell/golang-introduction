package executor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

// isolatedExecutor handles running code in an isolated environment.
type isolatedExecutor struct {
	tempDir   string
	maxOutput int
	timeout   string // For error messages
	logger    *slog.Logger
}

// newIsolatedExecutor creates a new isolated executor.
func newIsolatedExecutor(tempDir string, maxOutput int, timeout string, logger *slog.Logger) *isolatedExecutor {
	return &isolatedExecutor{
		tempDir:   tempDir,
		maxOutput: maxOutput,
		timeout:   timeout,
		logger:    logger,
	}
}

// execute runs code in an isolated temporary directory.
func (ie *isolatedExecutor) execute(ctx context.Context, code string) (*ExecutionResult, error) {
	// Create unique execution directory
	execDir, err := os.MkdirTemp(ie.tempDir, "exec-*")
	if err != nil {
		return nil, fmt.Errorf("create exec directory: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(execDir); err != nil {
			ie.logger.Error("failed to cleanup exec directory",
				"dir", execDir,
				"error", err)
		}
	}()

	// Create temporary Go file
	tmpFile, err := os.CreateTemp(execDir, "code-*.go")
	if err != nil {
		return nil, fmt.Errorf("create temp file: %w", err)
	}
	tmpFileName := tmpFile.Name()

	// Write and close file
	if err := writeCodeToFile(tmpFile, code); err != nil {
		return nil, fmt.Errorf("write code to file: %w", err)
	}

	// Execute the code
	return ie.runGoCode(ctx, execDir, tmpFileName)
}

// writeCodeToFile writes code to a file and closes it.
func writeCodeToFile(file *os.File, code string) error {
	defer file.Close()

	if _, err := file.WriteString(code); err != nil {
		return fmt.Errorf("write string: %w", err)
	}

	return nil
}

// runGoCode executes Go code using go run.
func (ie *isolatedExecutor) runGoCode(ctx context.Context, workDir, filename string) (*ExecutionResult, error) {
	cmd := exec.CommandContext(ctx, "go", "run", filename)
	cmd.Dir = workDir

	// Set minimal environment for security
	cmd.Env = buildSecureEnvironment(workDir)

	// Capture output
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	// Build result
	result := &ExecutionResult{
		ExitCode: 0,
		Output:   ie.truncateOutput(stdout.String()),
	}

	if err != nil {
		result.ExitCode = getExitCode(err)
		result.Error = ie.buildErrorMessage(ctx, stderr.String(), err)
	}

	return result, nil
}

// buildSecureEnvironment creates a minimal environment for code execution.
func buildSecureEnvironment(workDir string) []string {
	return []string{
		"PATH=/usr/local/bin:/usr/bin:/bin",
		"HOME=" + workDir,
		"TMPDIR=" + workDir,
		"TMP=" + workDir,
		"TEMP=" + workDir,
		"GOCACHE=" + workDir + "/.cache",
		"GOMODCACHE=" + workDir + "/.modcache",
	}
}

// truncateOutput truncates output if it exceeds maximum size.
func (ie *isolatedExecutor) truncateOutput(output string) string {
	if len(output) <= ie.maxOutput {
		return output
	}
	return output[:ie.maxOutput] + "\n... (output truncated)"
}

// getExitCode extracts the exit code from an error.
func getExitCode(err error) int {
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return exitErr.ExitCode()
	}
	return -1
}

// buildErrorMessage constructs an appropriate error message.
func (ie *isolatedExecutor) buildErrorMessage(ctx context.Context, stderr string, err error) string {
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return fmt.Sprintf("Execution timeout: code took longer than %v", ie.timeout)
	}

	stderr = ie.truncateOutput(stderr)
	if stderr != "" {
		return stderr
	}

	return err.Error()
}
