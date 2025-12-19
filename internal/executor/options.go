package executor

import (
	"log/slog"
	"time"
)

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
