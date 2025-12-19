package executor

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"
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

// codeValidator handles security validation of Go code.
type codeValidator struct {
	allowNetwork bool
	allowFileIO  bool
	logger       *slog.Logger
}

// newCodeValidator creates a new code validator.
func newCodeValidator(allowNetwork, allowFileIO bool, logger *slog.Logger) *codeValidator {
	return &codeValidator{
		allowNetwork: allowNetwork,
		allowFileIO:  allowFileIO,
		logger:       logger,
	}
}

// validate checks code for potentially dangerous patterns.
func (v *codeValidator) validate(code string) error {
	// Run all validation checks
	checks := []func(string) error{
		v.validateBlockedImports,
		v.validateDangerousImports,
		v.validateNetworkOperations,
		v.validateFileOperations,
		v.validateCommandExecution,
		v.validateSystemCalls,
	}

	for _, check := range checks {
		if err := check(code); err != nil {
			return err
		}
	}

	return nil
}

// validateBlockedImports checks for always-blocked imports.
func (v *codeValidator) validateBlockedImports(code string) error {
	for _, imp := range blockedImportsAlways {
		if containsImport(code, imp) {
			v.logger.Warn("blocked import detected",
				"import", imp,
				"reason", "always_blocked")
			return fmt.Errorf("%w: %s", ErrDangerousImport, imp)
		}
	}
	return nil
}

// validateDangerousImports checks for conditionally dangerous imports.
func (v *codeValidator) validateDangerousImports(code string) error {
	if v.allowFileIO {
		return nil // Skip if file I/O is allowed
	}

	for _, imp := range dangerousImports {
		if containsImport(code, imp) {
			v.logger.Warn("blocked import detected",
				"import", imp,
				"reason", "file_io_disabled")
			return fmt.Errorf("%w: %s (file I/O disabled)", ErrDangerousImport, imp)
		}
	}
	return nil
}

// validateNetworkOperations checks for network-related code.
func (v *codeValidator) validateNetworkOperations(code string) error {
	if v.allowNetwork {
		return nil // Skip if network is allowed
	}

	for _, pattern := range networkPatterns {
		if strings.Contains(code, pattern) {
			v.logger.Warn("blocked network operation",
				"pattern", pattern)
			return fmt.Errorf("%w: %s detected", ErrNetworkNotAllowed, pattern)
		}
	}
	return nil
}

// validateFileOperations checks for file system operations.
func (v *codeValidator) validateFileOperations(code string) error {
	if v.allowFileIO {
		return nil // Skip if file I/O is allowed
	}

	for _, pattern := range filePatterns {
		if strings.Contains(code, pattern) {
			v.logger.Warn("blocked file operation",
				"pattern", pattern)
			return fmt.Errorf("%w: %s detected", ErrFileIONotAllowed, pattern)
		}
	}
	return nil
}

// validateCommandExecution checks for command execution attempts.
func (v *codeValidator) validateCommandExecution(code string) error {
	if strings.Contains(code, "exec.Command") || strings.Contains(code, "exec.Run") {
		v.logger.Warn("blocked command execution attempt")
		return ErrCommandExecution
	}
	return nil
}

// validateSystemCalls checks for system call attempts.
func (v *codeValidator) validateSystemCalls(code string) error {
	if syscallRegex.MatchString(code) {
		v.logger.Warn("blocked system call attempt")
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
