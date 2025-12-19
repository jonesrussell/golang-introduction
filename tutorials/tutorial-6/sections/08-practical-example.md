# Practical Example: File Processor

**Duration:** 8-10 minutes

## Build Together

A robust file processing system demonstrating all error handling concepts.

```go runnable
package main

import (
    "errors"
    "fmt"
    "strings"
)

// ========================================
// Custom Error Types
// ========================================

// Sentinel errors
var (
    ErrEmptyFile     = errors.New("file is empty")
    ErrInvalidFormat = errors.New("invalid file format")
)

// ValidationError for data validation failures
type ValidationError struct {
    Line    int
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("line %d: %s - %s", e.Line, e.Field, e.Message)
}

// FileError wraps file operation errors
type FileError struct {
    Path      string
    Operation string
    Cause     error
}

func (e *FileError) Error() string {
    return fmt.Sprintf("%s %s: %v", e.Operation, e.Path, e.Cause)
}

func (e *FileError) Unwrap() error {
    return e.Cause
}

// ProcessingError collects multiple errors
type ProcessingError struct {
    File   string
    Errors []error
}

func (e *ProcessingError) Error() string {
    var sb strings.Builder
    sb.WriteString(fmt.Sprintf("processing %s failed with %d errors:\n", e.File, len(e.Errors)))
    for i, err := range e.Errors {
        sb.WriteString(fmt.Sprintf("  %d. %v\n", i+1, err))
    }
    return sb.String()
}

// ========================================
// Data Structures
// ========================================

type Record struct {
    ID    string
    Name  string
    Email string
    Age   int
}

// ========================================
// Record Validation
// ========================================

func validateRecord(r *Record, lineNum int) error {
    if r.ID == "" {
        return &ValidationError{Line: lineNum, Field: "id", Message: "required"}
    }
    if r.Name == "" {
        return &ValidationError{Line: lineNum, Field: "name", Message: "required"}
    }
    if !strings.Contains(r.Email, "@") {
        return &ValidationError{Line: lineNum, Field: "email", Message: "invalid format"}
    }
    if r.Age < 0 || r.Age > 150 {
        return &ValidationError{Line: lineNum, Field: "age", Message: "must be 0-150"}
    }
    return nil
}

// ========================================
// Processing Functions
// ========================================

func processRecords(records []Record, filename string, strict bool) ([]Record, error) {
    var errs []error
    var valid []Record

    for i, r := range records {
        if err := validateRecord(&r, i+1); err != nil {
            if strict {
                return nil, fmt.Errorf("record %d: %w", i+1, err)
            }
            errs = append(errs, err)
            continue
        }
        valid = append(valid, r)
    }

    if len(errs) > 0 {
        return valid, &ProcessingError{File: filename, Errors: errs}
    }

    return valid, nil
}

// ========================================
// Error Handling Demo
// ========================================

func handleError(err error) {
    // Check for specific error types
    var fileErr *FileError
    var valErr *ValidationError
    var procErr *ProcessingError

    switch {
    case errors.Is(err, ErrEmptyFile):
        fmt.Println("ERROR: File is empty")

    case errors.Is(err, ErrInvalidFormat):
        fmt.Println("ERROR: Invalid file format")

    case errors.As(err, &fileErr):
        fmt.Printf("FILE ERROR: %s failed on %s\n", fileErr.Operation, fileErr.Path)
        if fileErr.Cause != nil {
            fmt.Printf("  Cause: %v\n", fileErr.Cause)
        }

    case errors.As(err, &valErr):
        fmt.Printf("VALIDATION ERROR: Line %d, Field '%s': %s\n",
            valErr.Line, valErr.Field, valErr.Message)

    case errors.As(err, &procErr):
        fmt.Printf("PROCESSING WARNING: %d errors in %s\n",
            len(procErr.Errors), procErr.File)
        for _, e := range procErr.Errors {
            fmt.Printf("  - %v\n", e)
        }

    default:
        fmt.Printf("ERROR: %v\n", err)
    }
}

func main() {
    fmt.Println("=== Error Handling Demo ===\n")

    // Test data with some invalid records
    records := []Record{
        {ID: "1", Name: "Alice", Email: "alice@example.com", Age: 30},
        {ID: "2", Name: "Bob", Email: "bob@example.com", Age: 25},
        {ID: "", Name: "Invalid", Email: "no-email", Age: 999},  // Multiple errors
        {ID: "4", Name: "Charlie", Email: "charlie@example.com", Age: 35},
    }

    fmt.Println("Processing with strict=false (collect all errors):")
    valid, err := processRecords(records, "test.csv", false)
    if err != nil {
        handleError(err)
    }
    fmt.Printf("Valid records: %d\n", len(valid))

    fmt.Println("\n" + strings.Repeat("-", 40))
    fmt.Println("\nProcessing with strict=true (fail on first error):")
    _, err = processRecords(records, "test.csv", true)
    if err != nil {
        handleError(err)
    }

    fmt.Println("\n" + strings.Repeat("-", 40))
    fmt.Println("\nDemonstrating error wrapping:")
    
    // Simulate a wrapped error
    innerErr := errors.New("connection refused")
    wrappedErr := &FileError{
        Path:      "/data/users.json",
        Operation: "reading",
        Cause:     innerErr,
    }
    outerErr := fmt.Errorf("loading user data: %w", wrappedErr)
    
    fmt.Printf("Full error: %v\n", outerErr)
    handleError(outerErr)
}
```

## Walk Through:
- Custom error types for different failure modes
- Error wrapping with context
- Sentinel errors for common cases
- Strict vs non-strict processing modes
- Error handler uses Is/As to respond appropriately
- Multiple errors collected and returned together
