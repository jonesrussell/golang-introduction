package executor

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

// Default executor configuration values.
const (
	defaultTimeout       = 10 * time.Second
	defaultMaxOutput     = 10000 // 10KB
	defaultMaxMemoryMB   = 128
	defaultMaxCPUPercent = 50
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
	validator     *codeValidator
	isolator      *isolatedExecutor
}

// NewCodeExecutor creates a new code executor with security defaults.
func NewCodeExecutor(opts ...ExecutorOption) (*CodeExecutor, error) {
	tempDir, err := os.MkdirTemp("", "go-tutorial-exec-")
	if err != nil {
		return nil, fmt.Errorf("create temp directory: %w", err)
	}

	executor := &CodeExecutor{
		tempDir:       tempDir,
		timeout:       defaultTimeout,
		maxOutput:     defaultMaxOutput,
		allowNetwork:  false,
		allowFileIO:   false,
		maxMemoryMB:   defaultMaxMemoryMB,
		maxCPUPercent: defaultMaxCPUPercent,
		logger:        slog.Default(),
	}

	// Apply options
	for _, opt := range opts {
		opt(executor)
	}

	// Initialize validator and isolator with current settings
	executor.validator = newCodeValidator(
		executor.allowNetwork,
		executor.allowFileIO,
		executor.logger,
	)

	executor.isolator = newIsolatedExecutor(
		executor.tempDir,
		executor.maxOutput,
		executor.timeout.String(),
		executor.logger,
	)

	executor.logger.Info("code executor initialized",
		"temp_dir", tempDir,
		"timeout", executor.timeout,
		"allow_network", executor.allowNetwork,
		"allow_file_io", executor.allowFileIO,
	)

	return executor, nil
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
	if err := e.validator.validate(executableCode); err != nil {
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

	// Execute in isolated environment
	result, err := e.isolator.execute(execCtx, executableCode)
	if err != nil {
		return nil, fmt.Errorf("execute code: %w", err)
	}

	result.Duration = time.Since(startTime).String()

	e.logger.DebugContext(ctx, "code execution completed",
		"duration", result.Duration,
		"exit_code", result.ExitCode,
		"output_length", len(result.Output),
	)

	return result, nil
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
