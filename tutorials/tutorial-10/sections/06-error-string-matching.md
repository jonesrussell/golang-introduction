# Error String Matching

**Duration:** 4-5 minutes

## The Anti-Pattern

```go
// BAD: String matching for error handling
func HandleError(err error) {
    if err.Error() == "user not found" {
        // Handle not found
    }
    if strings.Contains(err.Error(), "timeout") {
        // Handle timeout
    }
    if strings.HasPrefix(err.Error(), "validation") {
        // Handle validation
    }
}
```

## Problems:
- Fragile (error message changes break code)
- No compile-time safety
- Doesn't work with wrapped errors

## The Fix

```go
// GOOD: Sentinel errors
var (
    ErrNotFound   = errors.New("user not found")
    ErrTimeout    = errors.New("operation timed out")
    ErrValidation = errors.New("validation failed")
)

func HandleError(err error) {
    if errors.Is(err, ErrNotFound) {
        // Handle not found
    }
    if errors.Is(err, ErrTimeout) {
        // Handle timeout
    }
}
```

## Custom Error Types

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func HandleError(err error) {
    var valErr *ValidationError
    if errors.As(err, &valErr) {
        fmt.Printf("Invalid field: %s\n", valErr.Field)
    }
}
```

## Key teaching points:
- Never match error strings
- Use [sentinel errors](https://pkg.go.dev/errors#New) for specific error conditions
- Use [`errors.Is()`](https://pkg.go.dev/errors#Is) to check for sentinel errors
- Use [`errors.As()`](https://pkg.go.dev/errors#As) for custom error types
- Works with [wrapped errors](https://pkg.go.dev/fmt#Errorf)
