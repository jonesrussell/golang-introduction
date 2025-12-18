# Go Concurrency Cheat Sheet

## Goroutines

```go
// Start a goroutine
go func() {
    fmt.Println("Hello from goroutine")
}()

// Pass arguments
for i := 0; i < 3; i++ {
    go func(n int) {
        fmt.Println(n)
    }(i)  // Pass i as argument
}
```

## Channels

```go
// Create channels
ch := make(chan int)        // Unbuffered
ch := make(chan int, 10)    // Buffered (capacity 10)

// Send and receive
ch <- 42        // Send
val := <-ch     // Receive
val, ok := <-ch // Receive with close check

// Close channel (sender only)
close(ch)

// Range over channel
for val := range ch {
    fmt.Println(val)
}
```

## Channel Direction

```go
func send(ch chan<- int) {    // Send-only
    ch <- 42
}

func receive(ch <-chan int) { // Receive-only
    val := <-ch
}
```

## Select

```go
select {
case val := <-ch1:
    fmt.Println("from ch1:", val)
case val := <-ch2:
    fmt.Println("from ch2:", val)
case ch3 <- 42:
    fmt.Println("sent to ch3")
default:
    fmt.Println("no channel ready")
}

// Timeout
select {
case result := <-ch:
    fmt.Println(result)
case <-time.After(5 * time.Second):
    fmt.Println("timeout")
}
```

## sync Package

### WaitGroup
```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println(n)
    }(i)
}

wg.Wait()
```

### Mutex
```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

### RWMutex
```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()         // Read lock
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *Cache) Set(key, val string) {
    c.mu.Lock()          // Write lock
    defer c.mu.Unlock()
    c.data[key] = val
}
```

### Once
```go
var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

## Context

```go
// With cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// With deadline
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour))
defer cancel()

// Check cancellation
select {
case <-ctx.Done():
    return ctx.Err()
default:
    // Continue work
}
```

## Common Patterns

### Worker Pool
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for r := 1; r <= 9; r++ {
        <-results
    }
}
```

### Semaphore
```go
sem := make(chan struct{}, 3)  // Max 3 concurrent

for _, item := range items {
    sem <- struct{}{}        // Acquire
    go func(i Item) {
        defer func() { <-sem }()  // Release
        process(i)
    }(item)
}
```

### Done Channel
```go
done := make(chan struct{})

go func() {
    for {
        select {
        case <-done:
            return
        default:
            // Work
        }
    }
}()

// Signal stop
close(done)
```

## Race Detection

```bash
go run -race main.go
go test -race ./...
```

## Best Practices

1. **Don't communicate by sharing memory; share memory by communicating**

2. **Always close channels from sender**

3. **Use context for cancellation**

4. **Prefer `sync.WaitGroup` over `time.Sleep`**

5. **Keep mutex scope small**

6. **Each goroutine must have exit path**

7. **Use buffered channels to prevent blocking**

8. **Run race detector in CI**
