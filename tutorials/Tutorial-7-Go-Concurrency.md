## **Video Tutorial Plan: Go Concurrency**

### **Video Metadata**
- **Title:** Go Concurrency Explained: Goroutines and Channels
- **Duration Target:** 45-55 minutes
- **Difficulty:** Intermediate
- **Prerequisites:** Go Basics, Functions, Error Handling

---

## **Video Structure**

### **1. Introduction (3-4 min)**
- Welcome and what viewers will learn
- Why Go excels at concurrency
- Concurrency vs parallelism
- Go's philosophy: "Don't communicate by sharing memory; share memory by communicating"
- Preview: Building a concurrent web scraper

---

### **2. Goroutines - The Basics (6-7 min)**

**Topics to cover:**
- What is a goroutine?
- Creating goroutines with `go` keyword
- Main goroutine behavior
- Goroutine lifecycle

**Code Examples:**
```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    // Regular function call - blocks until complete
    sayHello()

    // Goroutine - runs concurrently
    go sayHello()

    // Problem: main might exit before goroutine runs!
    // This sleep is just for demonstration
    time.Sleep(100 * time.Millisecond)

    fmt.Println("Main function done")
}

// Multiple goroutines
func count(name string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s: %d\n", name, i)
        time.Sleep(50 * time.Millisecond)
    }
}

func main() {
    go count("goroutine-1")
    go count("goroutine-2")
    count("main")  // Runs in main goroutine

    // Output is interleaved - concurrent execution!
}

// Anonymous goroutines
func main() {
    go func() {
        fmt.Println("Anonymous goroutine!")
    }()

    // Capture variables (be careful!)
    for i := 0; i < 3; i++ {
        go func(n int) {
            fmt.Println("Value:", n)
        }(i)  // Pass i as argument
    }

    time.Sleep(100 * time.Millisecond)
}

// Common gotcha: loop variable capture
func badExample() {
    for i := 0; i < 3; i++ {
        go func() {
            fmt.Println(i)  // BAD: likely prints "3" three times
        }()
    }
}

func goodExample() {
    for i := 0; i < 3; i++ {
        i := i  // Shadow variable
        go func() {
            fmt.Println(i)  // OK: each goroutine has its own copy
        }()
    }
}
```

**Key teaching points:**
- `go` keyword starts a goroutine
- Goroutines are lightweight (~2KB stack)
- Main exiting kills all goroutines
- Don't use `time.Sleep` for synchronization
- Capture loop variables correctly

---

### **3. Channels - Goroutine Communication (8-10 min)**

**Topics to cover:**
- What are channels?
- Creating channels
- Sending and receiving
- Channel direction
- Buffered vs unbuffered

**Code Examples:**
```go
// Creating and using channels
func main() {
    // Create an unbuffered channel
    ch := make(chan string)

    // Send in goroutine
    go func() {
        ch <- "Hello from goroutine!"  // Send to channel
    }()

    // Receive in main
    msg := <-ch  // Receive from channel (blocks until data arrives)
    fmt.Println(msg)
}

// Channels synchronize goroutines
func worker(done chan bool) {
    fmt.Println("Working...")
    time.Sleep(time.Second)
    fmt.Println("Done!")

    done <- true  // Signal completion
}

func main() {
    done := make(chan bool)
    go worker(done)

    <-done  // Wait for worker to signal
    fmt.Println("Worker finished")
}

// Channel direction in function signatures
func send(ch chan<- string) {  // Send-only channel
    ch <- "message"
}

func receive(ch <-chan string) {  // Receive-only channel
    msg := <-ch
    fmt.Println(msg)
}

// Buffered channels
func main() {
    // Unbuffered: send blocks until receive
    unbuffered := make(chan int)

    // Buffered: send blocks only when buffer is full
    buffered := make(chan int, 3)

    buffered <- 1  // Doesn't block
    buffered <- 2  // Doesn't block
    buffered <- 3  // Doesn't block
    // buffered <- 4  // Would block! Buffer full

    fmt.Println(<-buffered)  // 1
    fmt.Println(<-buffered)  // 2
    fmt.Println(<-buffered)  // 3
}

// Closing channels
func producer(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)  // Signal no more values
}

func main() {
    ch := make(chan int)
    go producer(ch)

    // Range over channel until closed
    for val := range ch {
        fmt.Println(val)
    }

    // Check if channel is closed
    ch2 := make(chan int)
    close(ch2)

    val, ok := <-ch2
    fmt.Printf("Value: %d, Open: %v\n", val, ok)  // Value: 0, Open: false
}
```

**Key teaching points:**
- Channels are typed conduits
- Unbuffered channels synchronize
- Buffered channels can hold N values
- Always close channels from sender side
- `range` iterates until channel closes
- Comma-ok checks if channel is open

---

### **4. Select Statement (6-7 min)**

**Topics to cover:**
- Multiple channel operations
- Non-blocking with default
- Timeouts
- First-response pattern

**Code Examples:**
```go
// Select waits on multiple channels
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "from ch1"
    }()

    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "from ch2"
    }()

    // Select waits on both, takes first available
    for i := 0; i < 2; i++ {
        select {
        case msg := <-ch1:
            fmt.Println("Received", msg)
        case msg := <-ch2:
            fmt.Println("Received", msg)
        }
    }
}

// Non-blocking operations with default
func main() {
    ch := make(chan int, 1)

    // Non-blocking receive
    select {
    case val := <-ch:
        fmt.Println("Received:", val)
    default:
        fmt.Println("No value ready")
    }

    // Non-blocking send
    ch <- 42
    select {
    case ch <- 100:
        fmt.Println("Sent 100")
    default:
        fmt.Println("Channel full")
    }
}

// Timeout pattern
func main() {
    ch := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch <- "result"
    }()

    select {
    case result := <-ch:
        fmt.Println("Got:", result)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout!")
    }
}

// Quit channel pattern
func worker(quit <-chan struct{}) {
    for {
        select {
        case <-quit:
            fmt.Println("Worker stopping")
            return
        default:
            fmt.Println("Working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    quit := make(chan struct{})

    go worker(quit)

    time.Sleep(2 * time.Second)
    close(quit)  // Signal worker to stop

    time.Sleep(100 * time.Millisecond)
}

// First response wins
func fetchFromServers() string {
    ch := make(chan string, 3)

    go func() { ch <- queryServer("server1") }()
    go func() { ch <- queryServer("server2") }()
    go func() { ch <- queryServer("server3") }()

    return <-ch  // Return first response
}
```

**Key teaching points:**
- Select handles multiple channel operations
- First ready case wins (random if multiple ready)
- Default makes operations non-blocking
- time.After for timeouts
- Empty struct channel for signals (zero memory)

---

### **5. Common Concurrency Patterns (8-10 min)**

**Topics to cover:**
- Worker pool
- Fan-out/fan-in
- Pipeline
- Semaphore

**Code Examples:**
```go
// Worker Pool Pattern
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second)  // Simulate work
        results <- job * 2
    }
}

func main() {
    numJobs := 10
    numWorkers := 3

    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Start workers
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for r := 1; r <= numJobs; r++ {
        fmt.Println("Result:", <-results)
    }
}

// Fan-out / Fan-in Pattern
func fanOut(input <-chan int, n int) []<-chan int {
    outputs := make([]<-chan int, n)
    for i := 0; i < n; i++ {
        outputs[i] = process(input)
    }
    return outputs
}

func fanIn(inputs ...<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup

    for _, ch := range inputs {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for val := range c {
                output <- val
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(output)
    }()

    return output
}

// Pipeline Pattern
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func double(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * 2
        }
        close(out)
    }()
    return out
}

func main() {
    // Pipeline: generate -> square -> double -> print
    nums := generate(1, 2, 3, 4, 5)
    squared := square(nums)
    doubled := double(squared)

    for result := range doubled {
        fmt.Println(result)  // 2, 8, 18, 32, 50
    }
}

// Semaphore Pattern (limit concurrency)
func main() {
    maxConcurrent := 3
    sem := make(chan struct{}, maxConcurrent)

    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            sem <- struct{}{}        // Acquire
            defer func() { <-sem }() // Release

            fmt.Printf("Task %d running\n", id)
            time.Sleep(time.Second)
        }(i)
    }

    wg.Wait()
}
```

---

### **6. sync Package Essentials (6-7 min)**

**Topics to cover:**
- sync.WaitGroup
- sync.Mutex
- sync.RWMutex
- sync.Once

**Code Examples:**
```go
import "sync"

// WaitGroup - wait for goroutines to complete
func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }

    wg.Wait()  // Block until all workers done
    fmt.Println("All workers completed")
}

// Mutex - protect shared state
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := &Counter{}
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    wg.Wait()
    fmt.Println("Final count:", counter.Value())  // Always 1000
}

// RWMutex - multiple readers, single writer
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()  // Read lock - allows concurrent reads
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()  // Write lock - exclusive access
    defer c.mu.Unlock()
    c.data[key] = value
}

// sync.Once - run initialization exactly once
var (
    instance *Database
    once     sync.Once
)

func GetDatabase() *Database {
    once.Do(func() {
        fmt.Println("Initializing database...")
        instance = &Database{}
    })
    return instance
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            db := GetDatabase()  // Only initializes once
            _ = db
        }()
    }
    wg.Wait()
}
```

**Key teaching points:**
- WaitGroup for waiting on multiple goroutines
- Mutex for protecting shared state
- RWMutex when reads >> writes
- Once for one-time initialization
- Always use defer with Lock/Unlock

---

### **7. Context Package (5-6 min)**

**Topics to cover:**
- Cancellation propagation
- Timeouts and deadlines
- Passing request-scoped values
- Context best practices

**Code Examples:**
```go
import "context"

// Cancellation
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

// Timeout
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

// Deadline
func main() {
    deadline := time.Now().Add(3 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()

    // ...
}

// Values (use sparingly!)
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

// HTTP handler with context
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    result := make(chan string, 1)
    go func() {
        // Simulate slow operation
        time.Sleep(2 * time.Second)
        result <- "done"
    }()

    select {
    case res := <-result:
        fmt.Fprintln(w, res)
    case <-ctx.Done():
        // Client disconnected or timeout
        http.Error(w, "Request cancelled", http.StatusRequestTimeout)
    }
}
```

**Key teaching points:**
- Context propagates cancellation
- Always pass context as first parameter
- Use WithCancel, WithTimeout, WithDeadline
- Avoid context.WithValue for most data
- Check ctx.Done() in long operations

---

### **8. Practical Example: Concurrent Web Scraper (10-12 min)**

**Build together:** A concurrent web scraper with rate limiting

```go
package main

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "sync"
    "time"
)

// Result holds scraping results
type Result struct {
    URL        string
    StatusCode int
    Size       int64
    Duration   time.Duration
    Error      error
}

// Scraper handles concurrent URL fetching
type Scraper struct {
    client      *http.Client
    maxWorkers  int
    rateLimit   time.Duration
    results     chan Result
}

func NewScraper(maxWorkers int, timeout, rateLimit time.Duration) *Scraper {
    return &Scraper{
        client: &http.Client{
            Timeout: timeout,
        },
        maxWorkers: maxWorkers,
        rateLimit:  rateLimit,
        results:    make(chan Result),
    }
}

func (s *Scraper) Scrape(ctx context.Context, urls []string) <-chan Result {
    // Create URL channel
    urlChan := make(chan string)

    // Create worker pool
    var wg sync.WaitGroup
    for i := 0; i < s.maxWorkers; i++ {
        wg.Add(1)
        go s.worker(ctx, i, urlChan, &wg)
    }

    // Send URLs with rate limiting
    go func() {
        ticker := time.NewTicker(s.rateLimit)
        defer ticker.Stop()

        for _, url := range urls {
            select {
            case <-ctx.Done():
                break
            case <-ticker.C:
                urlChan <- url
            }
        }
        close(urlChan)
    }()

    // Close results when all workers done
    go func() {
        wg.Wait()
        close(s.results)
    }()

    return s.results
}

func (s *Scraper) worker(ctx context.Context, id int, urls <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()

    for url := range urls {
        select {
        case <-ctx.Done():
            return
        default:
            result := s.fetch(ctx, url)
            s.results <- result
        }
    }
}

func (s *Scraper) fetch(ctx context.Context, url string) Result {
    start := time.Now()

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return Result{URL: url, Error: err, Duration: time.Since(start)}
    }

    resp, err := s.client.Do(req)
    if err != nil {
        return Result{URL: url, Error: err, Duration: time.Since(start)}
    }
    defer resp.Body.Close()

    size, _ := io.Copy(io.Discard, resp.Body)

    return Result{
        URL:        url,
        StatusCode: resp.StatusCode,
        Size:       size,
        Duration:   time.Since(start),
    }
}

// Progress tracks scraping progress
type Progress struct {
    mu        sync.Mutex
    total     int
    completed int
    success   int
    failed    int
}

func (p *Progress) Update(success bool) {
    p.mu.Lock()
    defer p.mu.Unlock()

    p.completed++
    if success {
        p.success++
    } else {
        p.failed++
    }
}

func (p *Progress) Print() {
    p.mu.Lock()
    defer p.mu.Unlock()

    fmt.Printf("\rProgress: %d/%d (Success: %d, Failed: %d)",
        p.completed, p.total, p.success, p.failed)
}

func main() {
    fmt.Println("=== Concurrent Web Scraper ===\n")

    urls := []string{
        "https://httpbin.org/get",
        "https://httpbin.org/delay/1",
        "https://httpbin.org/status/404",
        "https://httpbin.org/status/500",
        "https://httpbin.org/headers",
        "https://invalid.url.example",
        "https://httpbin.org/ip",
        "https://httpbin.org/user-agent",
    }

    // Create scraper
    scraper := NewScraper(
        3,                      // 3 concurrent workers
        10*time.Second,         // 10s timeout per request
        200*time.Millisecond,   // 200ms between requests
    )

    // Create cancellable context
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Track progress
    progress := &Progress{total: len(urls)}

    // Start scraping
    fmt.Printf("Scraping %d URLs with %d workers...\n\n", len(urls), 3)
    results := scraper.Scrape(ctx, urls)

    // Process results as they arrive
    var allResults []Result
    for result := range results {
        allResults = append(allResults, result)
        progress.Update(result.Error == nil && result.StatusCode == 200)
        progress.Print()
    }

    // Print summary
    fmt.Println("\n\n=== Results ===")
    for _, r := range allResults {
        if r.Error != nil {
            fmt.Printf("FAIL  %s: %v\n", r.URL, r.Error)
        } else {
            fmt.Printf("%-5d %s (%d bytes, %v)\n",
                r.StatusCode, r.URL, r.Size, r.Duration.Round(time.Millisecond))
        }
    }

    // Statistics
    fmt.Printf("\n=== Statistics ===\n")
    fmt.Printf("Total: %d, Success: %d, Failed: %d\n",
        progress.total, progress.success, progress.failed)
}
```

**Walk through:**
- Worker pool for concurrent fetching
- Rate limiting with ticker
- Context for cancellation
- WaitGroup for synchronization
- Mutex for progress tracking
- Channels for result collection
- HTTP client with timeouts

---

### **9. Common Pitfalls (5-6 min)**

**Topics to cover:**
- Race conditions
- Deadlocks
- Goroutine leaks
- Closing channels incorrectly

**Code Examples:**
```go
// PITFALL 1: Race condition
type Counter struct {
    value int  // No mutex!
}

func (c *Counter) Increment() {
    c.value++  // DATA RACE!
}

// Fix: Use mutex or atomic
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

// Or use atomic
import "sync/atomic"

type AtomicCounter struct {
    value int64
}

func (c *AtomicCounter) Increment() {
    atomic.AddInt64(&c.value, 1)
}

// PITFALL 2: Deadlock - circular wait
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

// PITFALL 3: Goroutine leak
func leak() {
    ch := make(chan int)  // Unbuffered

    go func() {
        ch <- 42  // Blocks forever if nobody receives!
    }()

    // Function returns without receiving
    // Goroutine is leaked!
}

// Fix: Use buffered channel or ensure receiver
func noLeak() {
    ch := make(chan int, 1)  // Buffered

    go func() {
        ch <- 42  // Doesn't block
    }()

    // Or use select with context
}

// PITFALL 4: Closing channel from wrong side
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

// PITFALL 5: Sending to closed channel (panic!)
func sendToClosed() {
    ch := make(chan int)
    close(ch)
    ch <- 1  // PANIC!
}

// PITFALL 6: Forgetting WaitGroup
func forgotWait() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(time.Second)
    }()

    // Forgot wg.Wait()!
    // Main exits before goroutine completes
}

// PITFALL 7: WaitGroup copy (value vs pointer)
func badWaitGroup() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func(wg sync.WaitGroup) {  // BAD: wg is copied!
        defer wg.Done()  // Done on copy, not original
    }(wg)

    wg.Wait()  // Waits forever
}

// Fix: Pass pointer
func goodWaitGroup() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {  // Closure captures wg
        defer wg.Done()
    }()

    wg.Wait()
}
```

**Detecting race conditions:**
```bash
# Run with race detector
go run -race main.go
go test -race ./...
```

**Key teaching points:**
- Use `go run -race` to detect races
- Always have a receiver for unbuffered channels
- Only sender closes channels
- Never send to closed channel
- Use context for cancellation
- Don't copy sync types

---

### **10. Best Practices Summary (3-4 min)**

**Cover these guidelines:**

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

---

### **11. Next Steps & Wrap-up (2-3 min)**

**Recap what was covered:**
- Goroutines and the `go` keyword
- Channels for communication
- Select for multiple channels
- Concurrency patterns
- sync package (WaitGroup, Mutex, Once)
- Context for cancellation
- Web scraper example
- Common pitfalls

**Preview next topics:**
- Slices and maps internals
- Testing concurrent code
- Production monitoring

**Homework/Practice suggestions:**
1. **Easy:** Concurrent file downloader
2. **Medium:** Chat server with multiple clients
3. **Challenge:** Rate-limited API client
4. **Advanced:** Distributed task queue

**Resources:**
- Go Blog: "Go Concurrency Patterns"
- Go Blog: "Pipelines and cancellation"
- Your GitHub repo with scraper code

---

## **Production Notes**

### **Screen Setup:**
- Code editor: 65% of screen
- Terminal output: 35% (shows concurrent output)
- Font size: 18-20pt minimum

### **Teaching Techniques:**
- Show interleaved output from goroutines
- Visualize channel send/receive blocking
- Demonstrate race detector
- Show deadlock messages
- Use timing to demonstrate concurrency

### **Visual Aids:**
- Diagram: Goroutines vs threads
- Diagram: Channel send/receive blocking
- Diagram: Worker pool pattern
- Animation: Pipeline data flow
- Diagram: Select statement flow

### **Engagement:**
- "What happens if we remove the WaitGroup?"
- "Will this deadlock?" puzzles
- Live debugging with race detector
- "Optimize this concurrent code" exercises

---

## **Supplementary Materials to Provide**

1. **GitHub Repository:**
   - Complete web scraper
   - All pattern examples
   - Practice exercises with solutions
   - Race condition examples

2. **Cheat Sheet (PDF/Gist):**
   ```
   Goroutine:           go func()
   Channel:             ch := make(chan T)
   Buffered:            ch := make(chan T, n)
   Send:                ch <- value
   Receive:             value := <-ch
   Close:               close(ch)
   Range:               for v := range ch {}
   Select:              select { case ... }

   WaitGroup:           wg.Add(1), wg.Done(), wg.Wait()
   Mutex:               mu.Lock(), mu.Unlock()
   Context:             ctx, cancel := context.WithCancel(ctx)
   ```

3. **Practice Exercises:**
   - **Easy:** Parallel sum of array
   - **Medium:** Producer-consumer queue
   - **Challenge:** Connection pool
   - **Advanced:** Distributed lock

4. **Concurrency Patterns PDF:**
   - All patterns with diagrams
   - When to use each pattern
   - Anti-patterns to avoid

---

This tutorial covers Go concurrency comprehensively, from basics to production patterns. The web scraper example demonstrates real-world usage with proper error handling, cancellation, and rate limiting.
