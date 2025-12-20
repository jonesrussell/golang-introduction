package executor

import (
	"context"
	"fmt"
	"log/slog"
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

// CodeExecutor handles execution of Go code with security restrictions via Docker.
type CodeExecutor struct {
	timeout       time.Duration
	maxOutput     int
	maxMemoryMB   int
	maxCPUPercent int
	compileImage  string
	execImage     string
	logger        *slog.Logger
	dockerExec    *dockerExecutor
}

// NewCodeExecutor creates a new code executor with security defaults using Docker.
func NewCodeExecutor(opts ...ExecutorOption) (*CodeExecutor, error) {
	executor := &CodeExecutor{
		timeout:       defaultTimeout,
		maxOutput:     defaultMaxOutput,
		maxMemoryMB:   defaultMaxMemoryMB,
		maxCPUPercent: defaultMaxCPUPercent,
		compileImage:  defaultCompileImage,
		execImage:     defaultExecImage,
		logger:        slog.Default(),
	}

	// Apply options
	for _, opt := range opts {
		opt(executor)
	}

	// Initialize Docker executor
	dockerExec, err := newDockerExecutor(
		executor.compileImage,
		executor.execImage,
		executor.maxMemoryMB,
		executor.maxCPUPercent,
		executor.maxOutput,
		executor.timeout,
		executor.logger,
	)
	if err != nil {
		return nil, fmt.Errorf("initialize Docker executor: %w", err)
	}

	executor.dockerExec = dockerExec

	executor.logger.Info("code executor initialized",
		"timeout", executor.timeout,
		"max_memory_mb", executor.maxMemoryMB,
		"max_cpu_percent", executor.maxCPUPercent,
		"compile_image", executor.compileImage,
		"exec_image", executor.execImage,
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
	// Prepare code for execution (wrap if needed)
	executableCode := PrepareForExecution(code, isSnippet)

	// Create execution context with timeout
	execCtx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	// Execute in Docker container (includes compilation and execution)
	result, err := e.dockerExec.execute(execCtx, executableCode)
	if err != nil {
		return nil, fmt.Errorf("execute code: %w", err)
	}

	e.logger.DebugContext(ctx, "code execution completed",
		"duration", result.Duration,
		"exit_code", result.ExitCode,
		"output_length", len(result.Output),
	)

	return result, nil
}

// Cleanup cleans up resources (containers auto-remove, so this is mostly a no-op for compatibility).
func (e *CodeExecutor) Cleanup() error {
	// Docker containers use AutoRemove: true, so cleanup is automatic
	// This method is kept for backward compatibility
	e.logger.Info("code executor cleaned up")
	return nil
}
