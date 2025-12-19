# Goroutine Leaks

**Duration:** 5-6 minutes

## The Anti-Pattern

```go
// BAD: Goroutine that can't exit
func startWorker() {
    go func() {
        for {
            // Process forever
            item := <-workQueue  // Blocks forever if queue closes
            process(item)
        }
    }()
}

// BAD: Unbounded channel producer
func producer() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; ; i++ {
            ch <- i  // Blocks forever if no consumer
        }
    }()
    return ch
}

// BAD: Fire and forget with unbuffered channel
func process() {
    ch := make(chan result)
    go func() {
        r := doWork()
        ch <- r  // Blocks forever if main doesn't read
    }()

    // Timeout - goroutine leaks!
    select {
    case r := <-ch:
        return r
    case <-time.After(timeout):
        return nil  // Goroutine still blocked on send!
    }
}
```

## The Fix

```go
// GOOD: Goroutine with cancellation
func startWorker(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return  // Clean exit
            case item := <-workQueue:
                process(item)
            }
        }
    }()
}

// GOOD: Buffered channel for fire-and-forget
func process(ctx context.Context) *result {
    ch := make(chan *result, 1)  // Buffered!
    go func() {
        r := doWork()
        ch <- r  // Won't block even if nobody reads
    }()

    select {
    case r := <-ch:
        return r
    case <-ctx.Done():
        return nil  // Goroutine can still complete
    }
}

// GOOD: WaitGroup for cleanup
func processAll(items []Item) {
    var wg sync.WaitGroup
    for _, item := range items {
        wg.Add(1)
        go func(i Item) {
            defer wg.Done()
            process(i)
        }(item)
    }
    wg.Wait()  // Ensure all goroutines complete
}
```
