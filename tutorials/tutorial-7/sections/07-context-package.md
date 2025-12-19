# Context Package

**Duration:** 5-6 minutes

## Topics to cover:
- [Cancellation propagation](https://pkg.go.dev/context)
- [Timeouts and deadlines](https://pkg.go.dev/context#WithTimeout)
- [Passing request-scoped values](https://pkg.go.dev/context#WithValue)
- Context best practices

## Cancellation

```go
import "context"

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d working...\n", id)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }

    time.Sleep(2 * time.Second)
    cancel()  // Signal all workers to stop

    time.Sleep(100 * time.Millisecond)
}
```

## Timeout

```go
func slowOperation(ctx context.Context) error {
    select {
    case <-time.After(5 * time.Second):
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    if err := slowOperation(ctx); err != nil {
        fmt.Println("Operation failed:", err)  // context deadline exceeded
    }
}
```

## Context Values

```go
type contextKey string

func main() {
    ctx := context.WithValue(context.Background(), contextKey("requestID"), "abc123")
    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    if reqID, ok := ctx.Value(contextKey("requestID")).(string); ok {
        fmt.Println("Request ID:", reqID)
    }
}
```

## HTTP Handler with Context

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    result := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        result <- "done"
    }()

    select {
    case res := <-result:
        fmt.Fprintln(w, res)
    case <-ctx.Done():
        http.Error(w, "Request cancelled", http.StatusRequestTimeout)
    }
}
```

## Key teaching points:
- [Context](https://pkg.go.dev/context) propagates cancellation
- Always pass context as first parameter
- Use [`WithCancel`](https://pkg.go.dev/context#WithCancel), [`WithTimeout`](https://pkg.go.dev/context#WithTimeout), [`WithDeadline`](https://pkg.go.dev/context#WithDeadline)
- Avoid [`WithValue`](https://pkg.go.dev/context#WithValue) for most data
- Check [`ctx.Done()`](https://pkg.go.dev/context#Context) in long operations
