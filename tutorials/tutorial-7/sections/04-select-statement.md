# Select Statement

**Duration:** 6-7 minutes

## Topics to cover:
- Multiple channel operations
- Non-blocking with default
- Timeouts
- First-response pattern

## Code Examples

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
```

## Non-blocking with Default

```go
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
```

## Timeout Pattern

```go
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
```

## Quit Channel Pattern

```go
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
```

## Key teaching points:
- Select handles multiple channel operations
- First ready case wins (random if multiple ready)
- Default makes operations non-blocking
- time.After for timeouts
- Empty struct channel for signals (zero memory)
