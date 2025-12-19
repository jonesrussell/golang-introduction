# Context.Value Abuse

**Duration:** 6-7 minutes

## The Anti-Pattern

```go
// BAD: Using context for everything
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    ctx = context.WithValue(ctx, "userID", 123)
    ctx = context.WithValue(ctx, "requestID", "abc-123")
    ctx = context.WithValue(ctx, "permissions", []string{"read", "write"})
    ctx = context.WithValue(ctx, "config", &Config{})
    ctx = context.WithValue(ctx, "logger", logger)
    ctx = context.WithValue(ctx, "db", database)

    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    // Type assertions everywhere, no compile-time safety
    userID := ctx.Value("userID").(int)
    config := ctx.Value("config").(*Config)
    logger := ctx.Value("logger").(Logger)
}
```

## Problems:
- No type safety (runtime panics)
- Hidden dependencies
- Hard to test
- Unclear API contracts

## The Fix

```go
// GOOD: Explicit parameters for dependencies
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    // Context only for: cancellation, deadlines, request-scoped values
    ctx = context.WithValue(ctx, requestIDKey, "abc-123")

    userID := getUserID(r)
    processRequest(ctx, userID, s.config, s.logger, s.db)
}

// Explicit dependencies
func processRequest(
    ctx context.Context,
    userID int,
    config *Config,
    logger Logger,
    db Database,
) error {
    // Clear what this function needs
}

// Use typed keys for context values
type contextKey string
const requestIDKey contextKey = "requestID"
```

## When to Use Context.Value:
- Request ID / Trace ID
- Cancellation signals
- Deadlines
- Request-scoped auth info

## DON'T Use For:
- Dependencies (database, logger)
- Configuration
- Business logic data
