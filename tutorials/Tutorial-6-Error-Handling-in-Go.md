## **Video Tutorial Plan: Error Handling in Go**

### **Video Metadata**
- **Title:** Error Handling in Go: Patterns, Wrapping, and Custom Errors
- **Duration Target:** 35-45 minutes
- **Difficulty:** Intermediate
- **Prerequisites:** Go Basics, Interfaces

---

## **Video Structure**

### **1. Introduction (3-4 min)**
- Welcome and what viewers will learn
- Go's philosophy: errors are values
- No exceptions in Go (and why that's good)
- Preview: Building robust error handling
- Show the final example: File processor with comprehensive error handling

---

### **2. The Error Interface (4-5 min)**

**Topics to cover:**
- `error` is just an interface
- The `Error()` method
- Creating errors with `errors.New` and `fmt.Errorf`
- Zero value of error is `nil`

**Code Examples:**
```go
package main

import (
    "errors"
    "fmt"
)

// error is a built-in interface
// type error interface {
//     Error() string
// }

func main() {
    // Creating errors
    err1 := errors.New("something went wrong")
    fmt.Println(err1.Error())  // something went wrong
    fmt.Println(err1)          // same - fmt knows about Error()

    // With formatting
    name := "config.json"
    err2 := fmt.Errorf("failed to open file: %s", name)
    fmt.Println(err2)  // failed to open file: config.json

    // Error is nil when no error occurred
    var err error  // Zero value is nil
    if err == nil {
        fmt.Println("No error")
    }

    // Checking for errors
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

**Key teaching points:**
- `error` is an interface with single `Error() string` method
- Any type with `Error() string` is an error
- `errors.New()` creates simple errors
- `fmt.Errorf()` creates formatted errors
- `nil` means no error

---

### **3. Error Handling Patterns (6-7 min)**

**Topics to cover:**
- The if err != nil pattern
- Early returns
- Error propagation
- Don't ignore errors!

**Code Examples:**
```go
// Pattern 1: Basic error checking
func ReadFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err  // Propagate error
    }
    defer file.Close()

    // Continue with file operations
    return nil
}

// Pattern 2: Handle and continue
func ProcessFiles(paths []string) {
    for _, path := range paths {
        err := processFile(path)
        if err != nil {
            fmt.Printf("Warning: %s - %v\n", path, err)
            continue  // Handle error and continue
        }
        fmt.Printf("Processed: %s\n", path)
    }
}

// Pattern 3: Multiple error checks
func CreateUser(name, email string) (*User, error) {
    // Validate name
    if name == "" {
        return nil, errors.New("name cannot be empty")
    }

    // Validate email
    if !strings.Contains(email, "@") {
        return nil, errors.New("invalid email format")
    }

    // Create in database
    user, err := db.Create(name, email)
    if err != nil {
        return nil, err
    }

    // Send welcome email
    err = sendWelcomeEmail(user.Email)
    if err != nil {
        // Log but don't fail user creation
        log.Printf("Warning: welcome email failed: %v", err)
    }

    return user, nil
}

// Pattern 4: Cleanup on error
func ProcessWithCleanup() error {
    resource, err := acquireResource()
    if err != nil {
        return err
    }
    defer resource.Release()  // Always cleanup

    err = doWork(resource)
    if err != nil {
        return err  // resource.Release() still called
    }

    return nil
}

// Anti-pattern: Ignoring errors
func BadExample() {
    file, _ := os.Open("file.txt")  // BAD: ignoring error!
    // If file.txt doesn't exist, file is nil
    // Next line will panic
    file.Read(make([]byte, 100))
}

// Pattern 5: Inline error declaration
func ProcessData() error {
    if err := validateInput(); err != nil {
        return err
    }

    if data, err := fetchData(); err != nil {
        return err
    } else {
        return processData(data)
    }
}
```

**Key teaching points:**
- Always check returned errors
- Use early returns to reduce nesting
- `_` to ignore is almost always wrong
- Defer ensures cleanup even on error
- Log non-critical errors, return critical ones

---

### **4. Error Wrapping (7-8 min)**

**Topics to cover:**
- Why wrap errors? (context!)
- `fmt.Errorf` with `%w` verb
- `errors.Unwrap`
- `errors.Is` for comparison
- `errors.As` for type assertion

**Code Examples:**
```go
import (
    "errors"
    "fmt"
    "os"
)

// Define sentinel errors
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrInvalidInput = errors.New("invalid input")
)

// Wrapping errors adds context
func ReadConfig(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        // Wrap with context - %w preserves original error
        return nil, fmt.Errorf("reading config file %s: %w", path, err)
    }
    return data, nil
}

func LoadConfig() (*Config, error) {
    data, err := ReadConfig("/etc/app/config.json")
    if err != nil {
        // Add more context at each layer
        return nil, fmt.Errorf("loading configuration: %w", err)
    }

    config, err := parseConfig(data)
    if err != nil {
        return nil, fmt.Errorf("parsing configuration: %w", err)
    }

    return config, nil
}

// errors.Is - checks if error is (or wraps) a target error
func main() {
    _, err := LoadConfig()
    if err != nil {
        // Check for specific underlying error
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("Config file not found - using defaults")
            // Use default config
        } else if errors.Is(err, os.ErrPermission) {
            fmt.Println("Permission denied - check file permissions")
        } else {
            fmt.Printf("Configuration error: %v\n", err)
        }
    }

    // Full error chain:
    // "loading configuration: reading config file /etc/app/config.json: open /etc/app/config.json: no such file or directory"
}

// errors.As - extracts specific error type from chain
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

func ValidateUser(u *User) error {
    if u.Email == "" {
        return &ValidationError{Field: "email", Message: "required"}
    }
    return nil
}

func CreateUser(u *User) error {
    if err := ValidateUser(u); err != nil {
        return fmt.Errorf("user creation failed: %w", err)
    }
    // ...
    return nil
}

func main() {
    err := CreateUser(&User{})
    if err != nil {
        var valErr *ValidationError
        if errors.As(err, &valErr) {
            // Can access ValidationError fields
            fmt.Printf("Fix the %s field: %s\n", valErr.Field, valErr.Message)
        } else {
            fmt.Println("Error:", err)
        }
    }
}

// Unwrap - manually traverse error chain
func printErrorChain(err error) {
    for err != nil {
        fmt.Printf("- %v\n", err)
        err = errors.Unwrap(err)
    }
}
```

**Key teaching points:**
- `%w` wraps error, preserving the chain
- `%v` formats error but breaks the chain
- `errors.Is()` checks entire chain for match
- `errors.As()` extracts typed error from chain
- Add context at each layer for debugging
- Error messages should flow: "outer: middle: inner: root"

---

### **5. Custom Error Types (6-7 min)**

**Topics to cover:**
- When to create custom error types
- Implementing the error interface
- Adding fields for context
- Implementing `Unwrap()` for wrapping

**Code Examples:**
```go
// Simple custom error
type NotFoundError struct {
    Resource string
    ID       interface{}
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s with ID %v not found", e.Resource, e.ID)
}

// Usage
func GetUser(id int) (*User, error) {
    user := db.FindByID(id)
    if user == nil {
        return nil, &NotFoundError{Resource: "User", ID: id}
    }
    return user, nil
}

// Check for specific error type
func main() {
    user, err := GetUser(123)
    if err != nil {
        var notFound *NotFoundError
        if errors.As(err, &notFound) {
            fmt.Printf("Could not find %s #%v\n", notFound.Resource, notFound.ID)
        }
    }
}

// Custom error with wrapped cause
type DatabaseError struct {
    Operation string
    Table     string
    Cause     error  // Wrapped error
}

func (e *DatabaseError) Error() string {
    return fmt.Sprintf("database error during %s on %s: %v",
        e.Operation, e.Table, e.Cause)
}

// Implement Unwrap to support errors.Is/As
func (e *DatabaseError) Unwrap() error {
    return e.Cause
}

// Usage
func SaveUser(u *User) error {
    _, err := db.Exec("INSERT INTO users...")
    if err != nil {
        return &DatabaseError{
            Operation: "insert",
            Table:     "users",
            Cause:     err,
        }
    }
    return nil
}

func main() {
    err := SaveUser(user)
    if err != nil {
        // Can check wrapped error
        if errors.Is(err, sql.ErrNoRows) {
            // Handle specific SQL error
        }

        // Can extract DatabaseError
        var dbErr *DatabaseError
        if errors.As(err, &dbErr) {
            log.Printf("DB operation '%s' on '%s' failed",
                dbErr.Operation, dbErr.Table)
        }
    }
}

// Custom error with multiple causes (Go 1.20+)
type MultiError struct {
    Errors []error
}

func (m *MultiError) Error() string {
    var msgs []string
    for _, err := range m.Errors {
        msgs = append(msgs, err.Error())
    }
    return strings.Join(msgs, "; ")
}

// Go 1.20+: Unwrap returns all errors
func (m *MultiError) Unwrap() []error {
    return m.Errors
}

// HTTP error with status code
type HTTPError struct {
    StatusCode int
    Message    string
    Cause      error
}

func (e *HTTPError) Error() string {
    if e.Cause != nil {
        return fmt.Sprintf("HTTP %d: %s: %v", e.StatusCode, e.Message, e.Cause)
    }
    return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

func (e *HTTPError) Unwrap() error {
    return e.Cause
}

// Usage in handler
func handleRequest() error {
    user, err := GetUser(id)
    if err != nil {
        var notFound *NotFoundError
        if errors.As(err, &notFound) {
            return &HTTPError{
                StatusCode: 404,
                Message:    "User not found",
                Cause:      err,
            }
        }
        return &HTTPError{
            StatusCode: 500,
            Message:    "Internal error",
            Cause:      err,
        }
    }
    return nil
}
```

**Key teaching points:**
- Custom types add structured error data
- Implement `Error() string` to satisfy interface
- Implement `Unwrap() error` to support chain operations
- Use pointer receiver for error methods
- Custom errors enable programmatic error handling

---

### **6. Sentinel Errors (4-5 min)**

**Topics to cover:**
- What are sentinel errors?
- When to use them
- Package-level error variables
- Comparing with errors.Is

**Code Examples:**
```go
// Sentinel errors - predefined error values
// Convention: ErrXxx naming

package user

import "errors"

// Package-level sentinel errors
var (
    ErrNotFound       = errors.New("user not found")
    ErrDuplicateEmail = errors.New("email already exists")
    ErrInvalidAge     = errors.New("age must be positive")
    ErrUnauthorized   = errors.New("unauthorized access")
)

func GetUser(id int) (*User, error) {
    user := db.Find(id)
    if user == nil {
        return nil, ErrNotFound
    }
    return user, nil
}

func CreateUser(u *User) error {
    if exists := db.FindByEmail(u.Email); exists != nil {
        return ErrDuplicateEmail
    }
    if u.Age < 0 {
        return ErrInvalidAge
    }
    return db.Save(u)
}

// Caller code
func main() {
    user, err := user.GetUser(123)
    if err != nil {
        switch {
        case errors.Is(err, user.ErrNotFound):
            fmt.Println("User doesn't exist")
        case errors.Is(err, user.ErrUnauthorized):
            fmt.Println("You don't have permission")
        default:
            fmt.Println("Unexpected error:", err)
        }
    }
}

// Wrapping sentinel errors
func GetUserWithContext(id int) (*User, error) {
    user, err := db.Find(id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            // Wrap with our sentinel error
            return nil, fmt.Errorf("%w: id=%d", ErrNotFound, id)
        }
        return nil, fmt.Errorf("database error: %w", err)
    }
    return user, nil
}

// Standard library sentinel errors
import (
    "io"
    "os"
)

func ProcessFile(path string) error {
    f, err := os.Open(path)
    if errors.Is(err, os.ErrNotExist) {
        // File doesn't exist
    }
    if errors.Is(err, os.ErrPermission) {
        // No permission
    }

    // Reading
    buf := make([]byte, 1024)
    _, err = f.Read(buf)
    if errors.Is(err, io.EOF) {
        // End of file (not really an error)
    }
}
```

**When to use sentinel errors:**
```go
// GOOD: Sentinel errors
// - Error has no dynamic data
// - Callers need to check for specific error
// - Error is part of your API contract

var ErrNotFound = errors.New("not found")

// BETTER: Custom error type
// - Error needs context (which user? which file?)
// - Multiple pieces of information needed

type NotFoundError struct {
    Resource string
    ID       int
}

// AVOID: String comparison
// BAD:
if err.Error() == "not found" {  // Fragile!
    // ...
}

// GOOD:
if errors.Is(err, ErrNotFound) {  // Robust
    // ...
}
```

**Key teaching points:**
- Sentinel errors are package-level variables
- Use `errors.Is()` to compare, not `==` (for wrapped errors)
- Name convention: `ErrXxx`
- Document sentinel errors as part of API
- Choose sentinel vs custom type based on needs

---

### **7. Panic and Recover (5-6 min)**

**Topics to cover:**
- What is panic?
- When to use panic (almost never!)
- Recover for graceful handling
- Panic vs error

**Code Examples:**
```go
// Panic - stops normal execution
func doPanic() {
    panic("something terrible happened")
    fmt.Println("This never runs")
}

// Common causes of panic:
// - nil pointer dereference
// - index out of range
// - type assertion failure
// - calling panic() explicitly

// When to panic:
// 1. Unrecoverable programmer error
func MustCompileRegex(pattern string) *regexp.Regexp {
    r, err := regexp.Compile(pattern)
    if err != nil {
        panic(fmt.Sprintf("invalid regex pattern: %s", pattern))
    }
    return r
}

// 2. Initialization that cannot fail
var config = mustLoadConfig()

func mustLoadConfig() Config {
    cfg, err := loadConfig()
    if err != nil {
        panic(fmt.Sprintf("failed to load config: %v", err))
    }
    return cfg
}

// When NOT to panic:
// - File not found
// - Network error
// - Invalid user input
// - Any error that can occur at runtime
// These should return errors!

// Recover - catch panics
func safeOperation() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()

    riskyOperation()
    return nil
}

func riskyOperation() {
    // Might panic
    panic("oops!")
}

// Recover in HTTP server (simplified)
func recoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic recovered: %v\n%s", err, debug.Stack())
                http.Error(w, "Internal Server Error", 500)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

// Recover only works in deferred function
func badRecover() {
    if r := recover(); r != nil {
        // This DOESN'T work - not in deferred function
        fmt.Println("Recovered:", r)
    }
    panic("oops!")
}

// Panic across goroutines - each goroutine must recover itself
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Main recovered:", r)
        }
    }()

    go func() {
        // This panic CANNOT be recovered by main's defer
        // The program will crash
        panic("panic in goroutine")
    }()

    time.Sleep(time.Second)
}

// Each goroutine needs its own recovery
func safeGoroutine(fn func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Goroutine panic: %v", r)
            }
        }()
        fn()
    }()
}
```

**Panic vs Error decision:**
```go
// Return error for:
// - Expected failure conditions
// - User input validation
// - Resource not found
// - Network/IO errors
// - Anything the caller might handle

func OpenFile(path string) (*File, error) {
    // File might not exist - return error
}

// Use panic for:
// - Programmer mistakes (should be caught in testing)
// - Invariant violations
// - Impossible states (indicates bug)
// - Initialization failures (app can't run)

func (s *Stack) Pop() interface{} {
    if s.Len() == 0 {
        panic("pop from empty stack")  // Programmer error
    }
    // ...
}
```

**Key teaching points:**
- Panic should be rare in Go code
- Return errors for expected failures
- Panic for programmer errors/impossible states
- Recover only works in deferred functions
- Each goroutine must recover its own panics
- Convention: `Must*` functions panic on error

---

### **8. Practical Example: File Processor (8-10 min)**

**Build together:** A robust file processing system

```go
package main

import (
    "bufio"
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "path/filepath"
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

func (e *ProcessingError) Unwrap() []error {
    return e.Errors
}

// ========================================
// Data Structures
// ========================================

type Record struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

// ========================================
// File Operations
// ========================================

func openFile(path string) (*os.File, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, &FileError{
            Path:      path,
            Operation: "opening",
            Cause:     err,
        }
    }
    return file, nil
}

func validatePath(path string) error {
    ext := filepath.Ext(path)
    if ext != ".json" && ext != ".txt" {
        return fmt.Errorf("%w: expected .json or .txt, got %s", ErrInvalidFormat, ext)
    }
    return nil
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
// File Processor
// ========================================

type FileProcessor struct {
    StrictMode bool // Fail on first error vs collect all errors
}

func NewFileProcessor(strict bool) *FileProcessor {
    return &FileProcessor{StrictMode: strict}
}

func (p *FileProcessor) ProcessFile(path string) ([]Record, error) {
    // Validate path
    if err := validatePath(path); err != nil {
        return nil, fmt.Errorf("path validation: %w", err)
    }

    // Open file
    file, err := openFile(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Check if file is empty
    stat, err := file.Stat()
    if err != nil {
        return nil, &FileError{Path: path, Operation: "stat", Cause: err}
    }
    if stat.Size() == 0 {
        return nil, fmt.Errorf("processing %s: %w", path, ErrEmptyFile)
    }

    // Process based on extension
    ext := filepath.Ext(path)
    switch ext {
    case ".json":
        return p.processJSON(file, path)
    case ".txt":
        return p.processTXT(file, path)
    default:
        return nil, fmt.Errorf("%w: %s", ErrInvalidFormat, ext)
    }
}

func (p *FileProcessor) processJSON(file *os.File, path string) ([]Record, error) {
    var records []Record
    decoder := json.NewDecoder(file)

    if err := decoder.Decode(&records); err != nil {
        return nil, &FileError{
            Path:      path,
            Operation: "parsing JSON",
            Cause:     err,
        }
    }

    return p.validateRecords(records, path)
}

func (p *FileProcessor) processTXT(file *os.File, path string) ([]Record, error) {
    var records []Record
    scanner := bufio.NewScanner(file)
    lineNum := 0

    var errs []error

    for scanner.Scan() {
        lineNum++
        line := strings.TrimSpace(scanner.Text())

        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        record, err := parseLine(line, lineNum)
        if err != nil {
            if p.StrictMode {
                return nil, fmt.Errorf("parsing line %d: %w", lineNum, err)
            }
            errs = append(errs, err)
            continue
        }

        records = append(records, *record)
    }

    if err := scanner.Err(); err != nil {
        return nil, &FileError{Path: path, Operation: "reading", Cause: err}
    }

    if len(errs) > 0 {
        return records, &ProcessingError{File: path, Errors: errs}
    }

    return p.validateRecords(records, path)
}

func parseLine(line string, lineNum int) (*Record, error) {
    // Format: id,name,email,age
    parts := strings.Split(line, ",")
    if len(parts) != 4 {
        return nil, &ValidationError{
            Line:    lineNum,
            Field:   "format",
            Message: fmt.Sprintf("expected 4 fields, got %d", len(parts)),
        }
    }

    var age int
    if _, err := fmt.Sscanf(parts[3], "%d", &age); err != nil {
        return nil, &ValidationError{
            Line:    lineNum,
            Field:   "age",
            Message: "must be a number",
        }
    }

    return &Record{
        ID:    strings.TrimSpace(parts[0]),
        Name:  strings.TrimSpace(parts[1]),
        Email: strings.TrimSpace(parts[2]),
        Age:   age,
    }, nil
}

func (p *FileProcessor) validateRecords(records []Record, path string) ([]Record, error) {
    var errs []error
    var valid []Record

    for i, r := range records {
        if err := validateRecord(&r, i+1); err != nil {
            if p.StrictMode {
                return nil, fmt.Errorf("record %d: %w", i+1, err)
            }
            errs = append(errs, err)
            continue
        }
        valid = append(valid, r)
    }

    if len(errs) > 0 {
        return valid, &ProcessingError{File: path, Errors: errs}
    }

    return valid, nil
}

// ========================================
// Main - Demonstration
// ========================================

func main() {
    fmt.Println("=== File Processor Demo ===\n")

    processor := NewFileProcessor(false) // Non-strict mode

    // Test with various files
    testFiles := []string{
        "testdata/valid.json",
        "testdata/records.txt",
        "testdata/missing.txt",
        "testdata/invalid.csv",
        "testdata/malformed.json",
    }

    for _, path := range testFiles {
        fmt.Printf("Processing: %s\n", path)
        fmt.Println(strings.Repeat("-", 40))

        records, err := processor.ProcessFile(path)
        if err != nil {
            handleError(err)
        }

        if len(records) > 0 {
            fmt.Printf("Successfully processed %d records:\n", len(records))
            for _, r := range records {
                fmt.Printf("  - %s: %s (%s)\n", r.ID, r.Name, r.Email)
            }
        }
        fmt.Println()
    }
}

func handleError(err error) {
    // Check for specific error types
    var fileErr *FileError
    var valErr *ValidationError
    var procErr *ProcessingError

    switch {
    case errors.Is(err, os.ErrNotExist):
        fmt.Println("ERROR: File does not exist")

    case errors.Is(err, os.ErrPermission):
        fmt.Println("ERROR: Permission denied")

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
```

**Test data files:**
```json
// testdata/valid.json
[
    {"id": "1", "name": "Alice", "email": "alice@example.com", "age": 30},
    {"id": "2", "name": "Bob", "email": "bob@example.com", "age": 25}
]
```

```txt
# testdata/records.txt
1,Alice,alice@example.com,30
2,Bob,bob@example.com,25
3,Invalid,no-email,999
4,Charlie,charlie@example.com,35
```

**Walk through:**
- Custom error types for different failure modes
- Error wrapping with context
- Sentinel errors for common cases
- Strict vs non-strict processing modes
- Error handler uses Is/As to respond appropriately
- Multiple errors collected and returned together

---

### **9. Best Practices Summary (3-4 min)**

**Cover these guidelines:**

```go
// 1. ALWAYS check errors
result, err := doSomething()
if err != nil {
    return fmt.Errorf("doing something: %w", err)
}

// 2. Add context when wrapping
// BAD:
return err

// GOOD:
return fmt.Errorf("loading user %d: %w", id, err)

// 3. Use errors.Is for sentinel errors
if errors.Is(err, ErrNotFound) {
    // Handle not found
}

// 4. Use errors.As for custom error types
var valErr *ValidationError
if errors.As(err, &valErr) {
    // Access valErr.Field, valErr.Message
}

// 5. Return early for cleaner code
func process() error {
    if err := step1(); err != nil {
        return err
    }
    if err := step2(); err != nil {
        return err
    }
    return step3()
}

// 6. Document errors in function comments
// GetUser retrieves a user by ID.
// Returns ErrNotFound if user doesn't exist.
// Returns ErrUnauthorized if caller lacks permission.
func GetUser(id int) (*User, error)

// 7. Use defer for cleanup
file, err := os.Open(path)
if err != nil {
    return err
}
defer file.Close()

// 8. Don't use panic for expected errors
// Return error instead

// 9. Name sentinel errors with Err prefix
var ErrNotFound = errors.New("not found")

// 10. Implement Unwrap for custom error types
func (e *MyError) Unwrap() error {
    return e.Cause
}
```

---

### **10. Next Steps & Wrap-up (2-3 min)**

**Recap what was covered:**
- Error interface and creating errors
- Error handling patterns
- Wrapping with %w
- errors.Is and errors.As
- Custom error types
- Sentinel errors
- Panic and recover
- File processor example
- Best practices

**Preview next topics:**
- Concurrency (errors in goroutines)
- Testing error conditions
- Logging and errors

**Homework/Practice suggestions:**
1. **Easy:** Create custom errors for a calculator (DivByZero, Overflow)
2. **Medium:** Build an API client with proper error handling
3. **Challenge:** Implement retry logic with different error types
4. **Advanced:** Create an error aggregator for concurrent operations

**Resources:**
- Go Blog: "Error handling in Go"
- Go Blog: "Working with Errors in Go 1.13"
- Your GitHub repo with file processor code

---

## **Production Notes**

### **Screen Setup:**
- Code editor: 70% of screen
- Terminal output: 30% of screen
- Font size: 18-20pt minimum

### **Teaching Techniques:**
- Show error messages flowing through call stack
- Demonstrate errors.Is/As with wrapped errors
- Show what happens when errors are ignored
- Compare panic vs error returns
- Debug error chains step by step

### **Visual Aids:**
- Diagram: Error wrapping chain
- Diagram: errors.Is traversing chain
- Diagram: errors.As type extraction
- Flowchart: Error handling decision tree

### **Engagement:**
- "Should this panic or return error?" quizzes
- "What context would you add?" exercises
- Live debugging of error chains
- Refactor bad error handling code

---

## **Supplementary Materials to Provide**

1. **GitHub Repository:**
   - Complete file processor
   - All code examples from video
   - Test data files
   - Practice exercises with solutions

2. **Cheat Sheet (PDF/Gist):**
   ```
   Create error:        errors.New("message")
   Format error:        fmt.Errorf("context: %w", err)
   Check error type:    errors.Is(err, ErrTarget)
   Extract error:       errors.As(err, &targetPtr)
   Unwrap one level:    errors.Unwrap(err)

   Custom error:
     type MyError struct { ... }
     func (e *MyError) Error() string { ... }
     func (e *MyError) Unwrap() error { ... }
   ```

3. **Practice Exercises:**
   - **Easy:** Validation error type
   - **Medium:** HTTP client with retries
   - **Challenge:** Transaction system with rollback
   - **Advanced:** Distributed error collection

4. **Error Handling Decision Tree (PDF):**
   - When to use each approach
   - Error type selection guide
   - Best practices checklist

---

This tutorial covers Go's error handling philosophy comprehensively while building toward production-ready patterns. The file processor example demonstrates real-world error handling with custom types, wrapping, and proper error checking.
