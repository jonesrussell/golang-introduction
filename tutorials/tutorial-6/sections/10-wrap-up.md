# Next Steps & Wrap-up

**Duration:** 2-3 minutes

## Recap What Was Covered:
- Error interface and creating errors
- Error handling patterns
- Wrapping with %w
- errors.Is and errors.As
- Custom error types
- Sentinel errors
- Panic and recover
- File processor example
- Best practices

## Preview Next Topics:
- Concurrency (errors in goroutines)
- Testing error conditions
- Logging and errors

## Practice Suggestions:
1. **Easy:** Create custom errors for a calculator (DivByZero, Overflow)
2. **Medium:** Build an API client with proper error handling
3. **Challenge:** Implement retry logic with different error types
4. **Advanced:** Create an error aggregator for concurrent operations

## Cheat Sheet

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

## Resources:
- Go Blog: "Error handling in Go"
- Go Blog: "Working with Errors in Go 1.13"

## Key teaching points:
- Go's error handling is explicit and predictable
- Wrap errors to add context
- Use custom types for rich error data
- Document which errors your API can return
