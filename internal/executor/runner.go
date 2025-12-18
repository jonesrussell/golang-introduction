package executor

import (
	"context"
	"fmt"
	"os"
	"os/exec"
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
	tempDir    string
	timeout    time.Duration
	maxOutput  int // Maximum output size in bytes
}

// NewCodeExecutor creates a new code executor
func NewCodeExecutor() (*CodeExecutor, error) {
	tempDir, err := os.MkdirTemp("", "go-tutorial-exec-")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}

	return &CodeExecutor{
		tempDir:   tempDir,
		timeout:   10 * time.Second,
		maxOutput: 10000, // 10KB max output
	}, nil
}

// Execute runs Go code and returns the result
func (e *CodeExecutor) Execute(ctx context.Context, code string) (*ExecutionResult, error) {
	// Create a context with timeout
	execCtx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	// Create temporary Go file
	tmpFile, err := os.CreateTemp(e.tempDir, "code-*.go")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write code to file
	if _, err := tmpFile.WriteString(code); err != nil {
		return nil, fmt.Errorf("failed to write code: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return nil, fmt.Errorf("failed to close file: %w", err)
	}

	startTime := time.Now()

	// Execute with go run
	cmd := exec.CommandContext(execCtx, "go", "run", tmpFile.Name())
	
	// Capture output
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

