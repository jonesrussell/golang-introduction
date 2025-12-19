# Panic and Recover

**Duration:** 5-6 minutes

## Topics to cover:
- What is panic?
- When to use panic (almost never!)
- Recover for graceful handling
- Panic vs error

## Code Examples

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
```

## Recover

```go
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
```

## Panic Across Goroutines

```go
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

## Panic vs Error Decision

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

## Key teaching points:
- Panic should be rare in Go code
- Return errors for expected failures
- Panic for programmer errors/impossible states
- Recover only works in deferred functions
- Each goroutine must recover its own panics
- Convention: `Must*` functions panic on error
