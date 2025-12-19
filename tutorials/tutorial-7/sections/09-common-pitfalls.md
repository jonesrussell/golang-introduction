# Common Pitfalls

**Duration:** 5-6 minutes

## Topics to cover:
- Race conditions
- Deadlocks
- Goroutine leaks
- Closing channels incorrectly

## Race Condition

```go
// BAD: Race condition
type Counter struct {
    value int  // No mutex!
}

func (c *Counter) Increment() {
    c.value++  // DATA RACE!
}

// FIX: Use mutex or atomic
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

// Or use atomic
import "sync/atomic"

type AtomicCounter struct {
    value int64
}

func (c *AtomicCounter) Increment() {
    atomic.AddInt64(&c.value, 1)
}
```

## Deadlock

```go
// BAD: Deadlock - circular wait
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        <-ch1       // Wait for ch1
        ch2 <- 1    // Send to ch2
    }()

    <-ch2       // Wait for ch2
    ch1 <- 1    // Send to ch1

    // DEADLOCK: Both waiting for each other
}
```

## Goroutine Leak

```go
// BAD: Goroutine leak
func leak() {
    ch := make(chan int)  // Unbuffered

    go func() {
        ch <- 42  // Blocks forever if nobody receives!
    }()

    // Function returns without receiving
    // Goroutine is leaked!
}

// FIX: Use buffered channel or ensure receiver
func noLeak() {
    ch := make(chan int, 1)  // Buffered

    go func() {
        ch <- 42  // Doesn't block
    }()

    // Or use select with context
}
```

## Wrong Channel Close

```go
// BAD: Closing channel from wrong side
func badClose() {
    ch := make(chan int)

    go func() {
        for val := range ch {
            fmt.Println(val)
        }
        close(ch)  // BAD: Receiver closing channel
    }()

    ch <- 1
    ch <- 2
    // close(ch)  // Sender should close
}

// BAD: Sending to closed channel (panic!)
func sendToClosed() {
    ch := make(chan int)
    close(ch)
    ch <- 1  // PANIC!
}
```

## WaitGroup Mistakes

```go
// BAD: WaitGroup copy (value vs pointer)
func badWaitGroup() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func(wg sync.WaitGroup) {  // BAD: wg is copied!
        defer wg.Done()  // Done on copy, not original
    }(wg)

    wg.Wait()  // Waits forever
}

// FIX: Pass pointer or use closure
func goodWaitGroup() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {  // Closure captures wg
        defer wg.Done()
    }()

    wg.Wait()
}
```

## Detecting Race Conditions

```bash
# Run with race detector
go run -race main.go
go test -race ./...
```

## Key teaching points:
- Use `go run -race` to detect races
- Always have a receiver for unbuffered channels
- Only sender closes channels
- Never send to closed channel
- Use context for cancellation
- Don't copy sync types
