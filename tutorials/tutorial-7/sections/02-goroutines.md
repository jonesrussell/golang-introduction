# Goroutines - The Basics

**Duration:** 6-7 minutes

## Topics to cover:
- What is a [goroutine](https://go.dev/ref/spec#Go_statements)?
- Creating goroutines with [`go` keyword](https://go.dev/ref/spec#Go_statements)
- Main goroutine behavior
- Goroutine lifecycle

## Code Examples

```go runnable
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func count(name string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s: %d\n", name, i)
        time.Sleep(50 * time.Millisecond)
    }
}

func main() {
    // Regular function call - blocks until complete
    sayHello()

    // Goroutine - runs concurrently
    go sayHello()

    // Multiple goroutines
    go count("goroutine-1")
    go count("goroutine-2")
    count("main")  // Runs in main goroutine

    // Output is interleaved - concurrent execution!
}
```

## Anonymous Goroutines

```go
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

## Key teaching points:
- [`go` keyword](https://go.dev/ref/spec#Go_statements) starts a goroutine
- Goroutines are lightweight (~2KB stack)
- Main exiting kills all goroutines
- Don't use `time.Sleep` for synchronization
- Capture loop variables correctly
