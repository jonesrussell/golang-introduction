# Best Practices Summary

**Duration:** 3-4 minutes

## Guidelines

```go
// 1. Keep goroutines focused and short-lived
go func() {
    result := doOneTask()
    resultCh <- result
}()

// 2. Use context for cancellation
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // Work
        }
    }
}

// 3. Prefer channels for communication
// DON'T share memory with mutex when channels work better

// 4. Close channels from sender, not receiver
go func() {
    defer close(ch)
    for _, item := range items {
        ch <- item
    }
}()

// 5. Use buffered channels to prevent blocking
ch := make(chan int, 100)

// 6. Always handle done/quit signals
select {
case result := <-resultCh:
    return result
case <-ctx.Done():
    return nil, ctx.Err()
}

// 7. Use WaitGroup for multiple goroutines
var wg sync.WaitGroup
for _, item := range items {
    wg.Add(1)
    go func(i Item) {
        defer wg.Done()
        process(i)
    }(item)
}
wg.Wait()

// 8. Limit concurrency
sem := make(chan struct{}, maxConcurrent)
sem <- struct{}{}  // Acquire
<-sem              // Release

// 9. Run race detector in tests
// go test -race ./...

// 10. Keep mutex scope small
func (c *Cache) Get(key string) string {
    c.mu.Lock()
    val := c.data[key]
    c.mu.Unlock()  // Release before returning
    return val
}
```

## Key teaching points:
- Design for cancellation from the start
- Prefer channels over shared memory when possible
- Always clean up goroutines
- Use the race detector
- Keep critical sections small
