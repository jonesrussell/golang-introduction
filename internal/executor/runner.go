package executor

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// ExecutionResult represents the result of code execution
type ExecutionResult struct {
	Output   string `json:"output"`
	Error    string `json:"error,omitempty"`
	ExitCode int    `json:"exitCode"`
	Duration string `json:"duration"`
}

// CodeExecutor handles execution of Go code
type CodeExecutor struct {
	tempDir       string
	timeout       time.Duration
	maxOutput     int  // Maximum output size in bytes
	allowNetwork  bool // Whether to allow network calls
	allowFileIO   bool // Whether to allow file I/O (restricted)
	maxMemoryMB   int  // Maximum memory in MB (0 = no limit, requires platform support)
	maxCPUPercent int  // Maximum CPU percentage (0 = no limit, requires platform support)
}

// NewCodeExecutor creates a new code executor
func NewCodeExecutor() (*CodeExecutor, error) {
	tempDir, err := os.MkdirTemp("", "go-tutorial-exec-")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}

	return &CodeExecutor{
		tempDir:       tempDir,
		timeout:       10 * time.Second,
		maxOutput:     10000, // 10KB max output
		allowNetwork:  false, // Disable network by default for security
		allowFileIO:   false, // Disable file I/O by default
		maxMemoryMB:   128,   // 128MB memory limit
		maxCPUPercent: 50,    // 50% CPU limit
	}, nil
}

// validateCode checks for potentially dangerous code patterns
func (e *CodeExecutor) validateCode(code string) error {
	// Check for dangerous imports
	dangerousImports := []string{
		"os/exec",
		"syscall",
		"unsafe",
		"runtime",
		"net/http",
		"net",
		"os",
		"io/ioutil",
		"path/filepath",
	}

	// Allow os if file I/O is enabled, but still restrict exec
	if !e.allowFileIO {
		for _, imp := range dangerousImports {
			if strings.Contains(code, `"`+imp+`"`) || strings.Contains(code, "`"+imp+"`") {
				return fmt.Errorf("import of %s is not allowed for security reasons", imp)
			}
		}
	} else {
		// Still block exec and syscall even if file I/O is allowed
		blockedImports := []string{"os/exec", "syscall", "unsafe"}
		for _, imp := range blockedImports {
			if strings.Contains(code, `"`+imp+`"`) || strings.Contains(code, "`"+imp+"`") {
				return fmt.Errorf("import of %s is not allowed for security reasons", imp)
			}
		}
	}

	// Check for network calls if not allowed
	if !e.allowNetwork {
		networkPatterns := []string{
			"http.Get",
			"http.Post",
			"net.Dial",
			"net.Listen",
			"http.Client",
		}
		for _, pattern := range networkPatterns {
			if strings.Contains(code, pattern) {
				return fmt.Errorf("network operations are not allowed for security reasons")
			}
		}
	}

	// Check for file system operations if not allowed
	if !e.allowFileIO {
		filePatterns := []string{
			"os.Create",
			"os.Open",
			"os.WriteFile",
			"os.ReadFile",
			"ioutil.ReadFile",
			"ioutil.WriteFile",
		}
		for _, pattern := range filePatterns {
			if strings.Contains(code, pattern) {
				return fmt.Errorf("file I/O operations are not allowed for security reasons")
			}
		}
	}

	// Check for exec.Command or similar
	if strings.Contains(code, "exec.Command") || strings.Contains(code, "exec.Run") {
		return fmt.Errorf("executing external commands is not allowed for security reasons")
	}

	// Check for system calls
	if matched, _ := regexp.MatchString(`syscall\.`, code); matched {
		return fmt.Errorf("system calls are not allowed for security reasons")
	}

	return nil
}

// Execute runs Go code and returns the result
func (e *CodeExecutor) Execute(ctx context.Context, code string) (*ExecutionResult, error) {
	// Validate code for security
	if err := e.validateCode(code); err != nil {
		return &ExecutionResult{
			Output:   "",
			Error:    err.Error(),
			ExitCode: -1,
			Duration: "0s",
		}, nil
	}

	// Create a context with timeout
	execCtx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	// Create a unique subdirectory for this execution
	execDir, err := os.MkdirTemp(e.tempDir, "exec-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create exec directory: %w", err)
	}

	// Ensure cleanup happens even if execution fails
	defer func() {
		if cleanupErr := os.RemoveAll(execDir); cleanupErr != nil {
			// Log cleanup error but don't fail the execution
			fmt.Fprintf(os.Stderr, "Warning: failed to cleanup exec directory: %v\n", cleanupErr)
		}
	}()

	// Create temporary Go file in the isolated directory
	tmpFile, err := os.CreateTemp(execDir, "code-*.go")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpFileName := tmpFile.Name()

	// Ensure file cleanup
	defer func() {
		if cleanupErr := os.Remove(tmpFileName); cleanupErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to cleanup temp file: %v\n", cleanupErr)
		}
	}()

	// Write code to file
	if _, err := tmpFile.WriteString(code); err != nil {
		return nil, fmt.Errorf("failed to write code: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return nil, fmt.Errorf("failed to close file: %w", err)
	}

	startTime := time.Now()

	// Execute with go run
	// Set working directory to the isolated exec directory
	cmd := exec.CommandContext(execCtx, "go", "run", tmpFileName)
	cmd.Dir = execDir

	// Set resource limits if supported (Linux)
	// Note: This requires platform-specific code or external tools
	// For now, we rely on timeout and validation

	// Set environment variables - minimal environment for security
	cmd.Env = []string{
		"PATH=/usr/local/bin:/usr/bin:/bin",
		"HOME=" + execDir,
		"TMPDIR=" + execDir,
		"TMP=" + execDir,
		"TEMP=" + execDir,
	}

	// Capture output with size limits
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err = cmd.Run()
	duration := time.Since(startTime)

	result := &ExecutionResult{
		Duration: duration.String(),
		ExitCode: 0,
	}

	// Get output
	output := stdout.String()
	if len(output) > e.maxOutput {
		output = output[:e.maxOutput] + "\n... (output truncated)"
	}
	result.Output = output

	// Get error output
	errorOutput := stderr.String()
	if len(errorOutput) > e.maxOutput {
		errorOutput = errorOutput[:e.maxOutput] + "\n... (error output truncated)"
	}

	// Check for execution errors
	if err != nil {
		if execCtx.Err() == context.DeadlineExceeded {
			result.Error = "Execution timeout: code took too long to run"
			result.ExitCode = -1
		} else {
			result.Error = errorOutput
			if result.Error == "" {
				result.Error = err.Error()
			}
			if exitError, ok := err.(*exec.ExitError); ok {
				result.ExitCode = exitError.ExitCode()
			} else {
				result.ExitCode = -1
			}
		}
	}

	return result, nil
}

// Cleanup removes temporary files
func (e *CodeExecutor) Cleanup() error {
	return os.RemoveAll(e.tempDir)
}

// SetTimeout sets the execution timeout
func (e *CodeExecutor) SetTimeout(timeout time.Duration) {
	e.timeout = timeout
}

// SetMaxOutput sets the maximum output size in bytes
func (e *CodeExecutor) SetMaxOutput(maxBytes int) {
	e.maxOutput = maxBytes
}

// SetAllowNetwork enables or disables network access
func (e *CodeExecutor) SetAllowNetwork(allow bool) {
	e.allowNetwork = allow
}

// SetAllowFileIO enables or disables file I/O (with restrictions)
func (e *CodeExecutor) SetAllowFileIO(allow bool) {
	e.allowFileIO = allow
}

// SetMaxMemory sets the maximum memory limit in MB (requires platform support)
func (e *CodeExecutor) SetMaxMemory(mb int) {
	e.maxMemoryMB = mb
}

// SetMaxCPU sets the maximum CPU percentage (requires platform support)
func (e *CodeExecutor) SetMaxCPU(percent int) {
	e.maxCPUPercent = percent
}
