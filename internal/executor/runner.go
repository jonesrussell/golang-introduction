package executor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// ExecutionResult represents the result of code execution.
type ExecutionResult struct {
	Output   string `json:"output"`
	Error    string `json:"error,omitempty"`
	ExitCode int    `json:"exitCode"`
	Duration string `json:"duration"`
}

// CodeExecutor handles execution of Go code with security restrictions.
type CodeExecutor struct {
	tempDir       string
	timeout       time.Duration
	maxOutput     int
	allowNetwork  bool
	allowFileIO   bool
	maxMemoryMB   int
	maxCPUPercent int
	logger        *slog.Logger
}

// ExecutorOption is a functional option for configuring CodeExecutor.
type ExecutorOption func(*CodeExecutor)

// WithTimeout sets the execution timeout.
func WithTimeout(timeout time.Duration) ExecutorOption {
	return func(e *CodeExecutor) {
		e.timeout = timeout
	}
}

// WithMaxOutput sets the maximum output size in bytes.
func WithMaxOutput(maxBytes int) ExecutorOption {
	return func(e *CodeExecutor) {
		e.maxOutput = maxBytes
	}
}

// WithNetworkAccess enables or disables network access.
func WithNetworkAccess(allow bool) ExecutorOption {
	return func(e *CodeExecutor) {
		e.allowNetwork = allow
	}
}

// WithFileIO enables or disables file I/O operations.
func WithFileIO(allow bool) ExecutorOption {
	return func(e *CodeExecutor) {
		e.allowFileIO = allow
	}
}

// WithLogger sets a custom logger.
func WithLogger(logger *slog.Logger) ExecutorOption {
	return func(e *CodeExecutor) {
		e.logger = logger
	}
}

// NewCodeExecutor creates a new code executor with security defaults.
func NewCodeExecutor(opts ...ExecutorOption) (*CodeExecutor, error) {
	tempDir, err := os.MkdirTemp("", "go-tutorial-exec-")
	if err != nil {
		return nil, fmt.Errorf("create temp directory: %w", err)
	}

	executor := &CodeExecutor{
		tempDir:       tempDir,
		timeout:       10 * time.Second,
		maxOutput:     10000, // 10KB
		allowNetwork:  false,
		allowFileIO:   false,
		maxMemoryMB:   128,
		maxCPUPercent: 50,
		logger:        slog.Default(),
	}

	// Apply options
	for _, opt := range opts {
		opt(executor)
	}

	executor.logger.Info("code executor initialized",
		"temp_dir", tempDir,
		"timeout", executor.timeout,
		"allow_network", executor.allowNetwork,
		"allow_file_io", executor.allowFileIO,
	)

	return executor, nil
}

var (
	// ErrDangerousImport is returned when code contains blocked imports.
	ErrDangerousImport = errors.New("dangerous import detected")
	// ErrNetworkNotAllowed is returned when network operations are blocked.
	ErrNetworkNotAllowed = errors.New("network operations not allowed")
	// ErrFileIONotAllowed is returned when file I/O operations are blocked.
	ErrFileIONotAllowed = errors.New("file I/O operations not allowed")
	// ErrCommandExecution is returned when command execution is attempted.
	ErrCommandExecution = errors.New("command execution not allowed")
	// ErrSystemCall is returned when system calls are attempted.
	ErrSystemCall = errors.New("system calls not allowed")
)

// Security validation patterns
var (
	dangerousImports = []string{
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

	blockedImportsAlways = []string{
		"os/exec",
		"syscall",
		"unsafe",
	}

	networkPatterns = []string{
		"http.Get",
		"http.Post",
		"http.Put",
		"http.Delete",
		"http.Head",
		"net.Dial",
		"net.Listen",
		"http.Client",
		"http.NewRequest",
	}

	filePatterns = []string{
		"os.Create",
		"os.Open",
		"os.OpenFile",
		"os.WriteFile",
		"os.ReadFile",
		"ioutil.ReadFile",
		"ioutil.WriteFile",
		"os.Remove",
		"os.RemoveAll",
		"os.Mkdir",
		"os.MkdirAll",
	}

	syscallRegex = regexp.MustCompile(`syscall\.`)
)

// validateCode checks for potentially dangerous code patterns.
func (e *CodeExecutor) validateCode(code string) error {
	// Always block certain imports regardless of settings
	for _, imp := range blockedImportsAlways {
		if containsImport(code, imp) {
			e.logger.Warn("blocked import detected",
				"import", imp,
				"reason", "always_blocked")
			return fmt.Errorf("%w: %s", ErrDangerousImport, imp)
		}
	}

	// Check other dangerous imports based on settings
	if !e.allowFileIO {
		for _, imp := range dangerousImports {
			if containsImport(code, imp) {
				e.logger.Warn("blocked import detected",
					"import", imp,
					"reason", "file_io_disabled")
				return fmt.Errorf("%w: %s (file I/O disabled)", ErrDangerousImport, imp)
			}
		}
	}

	// Check for network operations
	if !e.allowNetwork {
		for _, pattern := range networkPatterns {
			if strings.Contains(code, pattern) {
				e.logger.Warn("blocked network operation",
					"pattern", pattern)
				return fmt.Errorf("%w: %s detected", ErrNetworkNotAllowed, pattern)
			}
		}
	}

	// Check for file system operations
	if !e.allowFileIO {
		for _, pattern := range filePatterns {
			if strings.Contains(code, pattern) {
				e.logger.Warn("blocked file operation",
					"pattern", pattern)
				return fmt.Errorf("%w: %s detected", ErrFileIONotAllowed, pattern)
			}
		}
	}

	// Check for command execution
	if strings.Contains(code, "exec.Command") || strings.Contains(code, "exec.Run") {
		e.logger.Warn("blocked command execution attempt")
		return ErrCommandExecution
	}

	// Check for system calls
	if syscallRegex.MatchString(code) {
		e.logger.Warn("blocked system call attempt")
		return ErrSystemCall
	}

	return nil
}

// containsImport checks if code contains a specific import.
func containsImport(code, importPath string) bool {
	// Check double-quoted imports
	if strings.Contains(code, `"`+importPath+`"`) {
		return true
	}
	// Check backtick imports (less common but possible)
	if strings.Contains(code, "`"+importPath+"`") {
		return true
	}
	return false
}

// Execute runs Go code and returns the result.
func (e *CodeExecutor) Execute(ctx context.Context, code string) (*ExecutionResult, error) {
	return e.ExecuteWithOptions(ctx, code, false)
}

// ExecuteSnippet runs a code snippet, auto-wrapping it first.
func (e *CodeExecutor) ExecuteSnippet(ctx context.Context, code string) (*ExecutionResult, error) {
	return e.ExecuteWithOptions(ctx, code, true)
}

// ExecuteWithOptions runs Go code with options for snippet handling.
func (e *CodeExecutor) ExecuteWithOptions(ctx context.Context, code string, isSnippet bool) (*ExecutionResult, error) {
	startTime := time.Now()

	// Prepare code for execution (wrap if needed)
	executableCode := PrepareForExecution(code, isSnippet)

	// Validate code for security
	if err := e.validateCode(executableCode); err != nil {
		return &ExecutionResult{
			Output:   "",
			Error:    err.Error(),
			ExitCode: -1,
			Duration: time.Since(startTime).String(),
		}, nil
	}

	// Create execution context with timeout
	execCtx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	// Execute in isolated directory
	result, err := e.executeInIsolation(execCtx, executableCode)
	if err != nil {
		return nil, fmt.Errorf("execute code: %w", err)
	}

	result.Duration = time.Since(startTime).String()

	e.logger.Debug("code execution completed",
		"duration", result.Duration,
		"exit_code", result.ExitCode,
		"output_length", len(result.Output),
	)

	return result, nil
}

// executeInIsolation runs code in an isolated temporary directory.
func (e *CodeExecutor) executeInIsolation(ctx context.Context, code string) (*ExecutionResult, error) {
	// Create unique execution directory
	execDir, err := os.MkdirTemp(e.tempDir, "exec-*")
	if err != nil {
		return nil, fmt.Errorf("create exec directory: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(execDir); err != nil {
			e.logger.Error("failed to cleanup exec directory",
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
	if err := e.writeCodeToFile(tmpFile, code); err != nil {
		return nil, fmt.Errorf("write code to file: %w", err)
	}

	// Execute the code
	return e.runGoCode(ctx, execDir, tmpFileName)
}

// writeCodeToFile writes code to a file and closes it.
func (e *CodeExecutor) writeCodeToFile(file *os.File, code string) error {
	defer file.Close()

	if _, err := file.WriteString(code); err != nil {
		return fmt.Errorf("write string: %w", err)
	}

	return nil
}

// runGoCode executes Go code using go run.
func (e *CodeExecutor) runGoCode(ctx context.Context, workDir, filename string) (*ExecutionResult, error) {
	cmd := exec.CommandContext(ctx, "go", "run", filename)
	cmd.Dir = workDir

	// Set minimal environment for security
	cmd.Env = e.buildSecureEnvironment(workDir)

	// Capture output
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	// Build result
	result := &ExecutionResult{
		ExitCode: 0,
		Output:   e.truncateOutput(stdout.String()),
	}

	if err != nil {
		result.ExitCode = e.getExitCode(err)
		result.Error = e.buildErrorMessage(ctx, stderr.String(), err)
	}

	return result, nil
}

// buildSecureEnvironment creates a minimal environment for code execution.
func (e *CodeExecutor) buildSecureEnvironment(workDir string) []string {
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
func (e *CodeExecutor) truncateOutput(output string) string {
	if len(output) <= e.maxOutput {
		return output
	}
	return output[:e.maxOutput] + "\n... (output truncated)"
}

// getExitCode extracts the exit code from an error.
func (e *CodeExecutor) getExitCode(err error) int {
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return exitErr.ExitCode()
	}
	return -1
}

// buildErrorMessage constructs an appropriate error message.
func (e *CodeExecutor) buildErrorMessage(ctx context.Context, stderr string, err error) string {
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return fmt.Sprintf("Execution timeout: code took longer than %v", e.timeout)
	}

	stderr = e.truncateOutput(stderr)
	if stderr != "" {
		return stderr
	}

	return err.Error()
}

// Cleanup removes all temporary files created by the executor.
func (e *CodeExecutor) Cleanup() error {
	if err := os.RemoveAll(e.tempDir); err != nil {
		e.logger.Error("cleanup failed",
			"temp_dir", e.tempDir,
			"error", err)
		return fmt.Errorf("remove temp directory: %w", err)
	}

	e.logger.Info("code executor cleaned up",
		"temp_dir", e.tempDir)
	return nil
}

// Deprecated setter methods - kept for backward compatibility
// Consider using functional options in NewCodeExecutor instead

// SetTimeout sets the execution timeout.
// Deprecated: Use WithTimeout option in NewCodeExecutor.
func (e *CodeExecutor) SetTimeout(timeout time.Duration) {
	e.timeout = timeout
}

// SetMaxOutput sets the maximum output size in bytes.
// Deprecated: Use WithMaxOutput option in NewCodeExecutor.
func (e *CodeExecutor) SetMaxOutput(maxBytes int) {
	e.maxOutput = maxBytes
}

// SetAllowNetwork enables or disables network access.
// Deprecated: Use WithNetworkAccess option in NewCodeExecutor.
func (e *CodeExecutor) SetAllowNetwork(allow bool) {
	e.allowNetwork = allow
}

// SetAllowFileIO enables or disables file I/O (with restrictions).
// Deprecated: Use WithFileIO option in NewCodeExecutor.
func (e *CodeExecutor) SetAllowFileIO(allow bool) {
	e.allowFileIO = allow
}

// SetMaxMemory sets the maximum memory limit in MB (requires platform support).
// Deprecated: Not currently implemented.
func (e *CodeExecutor) SetMaxMemory(mb int) {
	e.maxMemoryMB = mb
}

// SetMaxCPU sets the maximum CPU percentage (requires platform support).
// Deprecated: Not currently implemented.
func (e *CodeExecutor) SetMaxCPU(percent int) {
	e.maxCPUPercent = percent
}
