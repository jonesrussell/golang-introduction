package executor

import "errors"

// Sentinel errors for validation failures.
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
