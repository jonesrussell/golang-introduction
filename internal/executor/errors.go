package executor

import "errors"

// Sentinel errors for Docker execution failures.
var (
	// ErrDockerNotAvailable is returned when Docker daemon is not available.
	ErrDockerNotAvailable = errors.New("docker daemon not available")
	// ErrContainerExecution is returned when container execution fails.
	ErrContainerExecution = errors.New("container execution failed")
	// ErrCompilationFailed is returned when Go code compilation fails.
	ErrCompilationFailed = errors.New("compilation failed")
	// ErrTimeout is returned when execution exceeds the timeout.
	ErrTimeout = errors.New("execution timeout exceeded")
)
