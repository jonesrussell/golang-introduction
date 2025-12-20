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

// WithMaxMemory sets the maximum memory limit in MB (enforced via Docker).
func WithMaxMemory(mb int) ExecutorOption {
	return func(e *CodeExecutor) {
		e.maxMemoryMB = mb
	}
}

// WithMaxCPU sets the maximum CPU usage as a percentage (enforced via Docker).
func WithMaxCPU(percent int) ExecutorOption {
	return func(e *CodeExecutor) {
		e.maxCPUPercent = percent
	}
}

// WithDockerImage sets the Docker image for Go compilation.
func WithDockerImage(image string) ExecutorOption {
	return func(e *CodeExecutor) {
		e.compileImage = image
	}
}

// WithLogger sets a custom logger.
func WithLogger(logger *slog.Logger) ExecutorOption {
	return func(e *CodeExecutor) {
		e.logger = logger
	}
}
