# Best Practices Summary

**Duration:** 3-4 minutes

## Guidelines

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

## Key teaching points:
- Check every error
- Add context when wrapping
- Use [`Is`](https://pkg.go.dev/errors#Is)/[`As`](https://pkg.go.dev/errors#As) for type-safe checks
- Document errors in API
- Return early to reduce nesting
