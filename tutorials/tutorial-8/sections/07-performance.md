# Performance Considerations

**Duration:** 4-5 minutes

## Topics to cover:
- Pre-allocation with [`make`](https://pkg.go.dev/builtin#make)
- Memory layout
- Clearing collections
- When to use [arrays](https://go.dev/ref/spec#Array_types)

## Pre-allocation

```go runnable
package main

import (
    "fmt"
    "time"
)

func badAppend() []int {
    var s []int
    for i := 0; i < 10000; i++ {
        s = append(s, i)  // Many reallocations
    }
    return s
}

func goodAppend() []int {
    s := make([]int, 0, 10000)  // Pre-allocate
    for i := 0; i < 10000; i++ {
        s = append(s, i)  // No reallocations
    }
    return s
}

func main() {
    // Benchmark bad approach
    start := time.Now()
    for i := 0; i < 1000; i++ {
        _ = badAppend()
    }
    fmt.Printf("Without pre-allocation: %v\n", time.Since(start))

    // Benchmark good approach
    start = time.Now()
    for i := 0; i < 1000; i++ {
        _ = goodAppend()
    }
    fmt.Printf("With pre-allocation: %v\n", time.Since(start))
}
```

## Pre-allocate Maps

```go
// Hint: ~10000 entries
m := make(map[string]int, 10000)
```

## Clearing Collections

```go
// Clearing a slice (reuse backing array)
s := []int{1, 2, 3, 4, 5}
s = s[:0]  // Length 0, capacity preserved

// Clearing a map (create new)
m := map[string]int{"a": 1, "b": 2}
// Option 1: Create new map
m = make(map[string]int)
// Option 2: Delete all (Go 1.21+ has clear())
for k := range m {
    delete(m, k)
}
```

## Value vs Pointer Slices

```go
// Slice of structs vs slice of pointers
// Structs: better cache locality, fewer allocations
type DataValue struct {
    ID   int
    Name string
}
valSlice := make([]DataValue, 1000)  // Contiguous memory

// Pointers: easier modification, but scattered memory
ptrSlice := make([]*DataValue, 1000)  // Pointers scattered
```

## String Building

```go runnable
package main

import (
    "fmt"
    "strings"
    "time"
)

// BAD: String concatenation in loop
func badConcat(items []string) string {
    result := ""
    for _, item := range items {
        result += item + ","  // Creates new string each time!
    }
    return result
}

// GOOD: Use strings.Builder
func goodConcat(items []string) string {
    var sb strings.Builder
    for i, item := range items {
        if i > 0 {
            sb.WriteString(",")
        }
        sb.WriteString(item)
    }
    return sb.String()
}

func main() {
    items := make([]string, 1000)
    for i := range items {
        items[i] = "item"
    }

    start := time.Now()
    for i := 0; i < 100; i++ {
        _ = badConcat(items)
    }
    fmt.Printf("String concat: %v\n", time.Since(start))

    start = time.Now()
    for i := 0; i < 100; i++ {
        _ = goodConcat(items)
    }
    fmt.Printf("strings.Builder: %v\n", time.Since(start))
}
```
