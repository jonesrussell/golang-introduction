# Next Steps & Wrap-up

**Duration:** 2-3 minutes

## Recap What Was Covered:
- [Error interface](https://pkg.go.dev/builtin#error) and creating errors
- Error handling patterns
- Wrapping with [`%w`](https://pkg.go.dev/fmt#Errorf)
- [`errors.Is`](https://pkg.go.dev/errors#Is) and [`errors.As`](https://pkg.go.dev/errors#As)
- Custom error types
- Sentinel errors
- [Panic and recover](https://go.dev/ref/spec#Handling_panics)
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
- [Go Blog: "Error handling in Go"](https://go.dev/blog/error-handling-and-go): go.dev/blog/error-handling-and-go
- [Go Blog: "Working with Errors in Go 1.13"](https://go.dev/blog/go1.13-errors): go.dev/blog/go1.13-errors
- [Effective Go - Errors](https://go.dev/doc/effective_go#errors): go.dev/doc/effective_go#errors

## Key teaching points:
- Go's error handling is explicit and predictable
- Wrap errors to add context
- Use custom types for rich error data
- Document which errors your API can return
