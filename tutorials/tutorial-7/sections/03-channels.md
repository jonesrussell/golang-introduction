# Channels - Goroutine Communication

**Duration:** 8-10 minutes

## Topics to cover:
- What are channels?
- Creating channels
- Sending and receiving
- Channel direction
- Buffered vs unbuffered

## Code Examples

```go runnable
package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("Working...")
    time.Sleep(time.Second)
    fmt.Println("Done!")
    done <- true  // Signal completion
}

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

    // Channels synchronize goroutines
    done := make(chan bool)
    go worker(done)
    <-done  // Wait for worker to signal
    fmt.Println("Worker finished")
}
```

## Channel Direction

```go
func send(ch chan<- string) {  // Send-only channel
    ch <- "message"
}

func receive(ch <-chan string) {  // Receive-only channel
    msg := <-ch
    fmt.Println(msg)
}
```

## Buffered Channels

```go
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
```

## Closing Channels

```go
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

## Key teaching points:
- Channels are typed conduits
- Unbuffered channels synchronize
- Buffered channels can hold N values
- Always close channels from sender side
- `range` iterates until channel closes
- Comma-ok checks if channel is open
