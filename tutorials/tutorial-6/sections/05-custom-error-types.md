# Custom Error Types

**Duration:** 6-7 minutes

## Topics to cover:
- When to create custom error types
- Implementing the error interface
- Adding fields for context
- Implementing `Unwrap()` for wrapping

## Code Examples

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
```

## Custom Error with Wrapped Cause

```go
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
```

## HTTP Error Example

```go
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

## Key teaching points:
- Custom types add structured error data
- Implement `Error() string` to satisfy interface
- Implement `Unwrap() error` to support chain operations
- Use pointer receiver for error methods
- Custom errors enable programmatic error handling
