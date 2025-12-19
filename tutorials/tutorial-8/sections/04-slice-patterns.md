# Common Slice Patterns

**Duration:** 5-6 minutes

## Topics to cover:
- Filtering
- Removing elements
- Insert at position
- Stack/queue operations using [slices](https://go.dev/ref/spec#Slice_types)

## Filter Patterns

```go
// Filter in-place (no allocation)
func filter(s []int, keep func(int) bool) []int {
    n := 0
    for _, v := range s {
        if keep(v) {
            s[n] = v
            n++
        }
    }
    return s[:n]
}

// Filter with new slice (preserves original)
func filterCopy(s []int, keep func(int) bool) []int {
    result := make([]int, 0, len(s))
    for _, v := range s {
        if keep(v) {
            result = append(result, v)
        }
    }
    return result
}
```

## Remove Element

```go
// Remove element at index
func remove(s []int, i int) []int {
    return append(s[:i], s[i+1:]...)
}

// Remove without preserving order (faster)
func removeUnordered(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}
```

## Insert at Index

```go
// Insert at index
func insert(s []int, i int, v int) []int {
    s = append(s, 0)           // Grow by 1
    copy(s[i+1:], s[i:])       // Shift right
    s[i] = v
    return s
}
```

## Stack and Queue

```go runnable
package main

import "fmt"

// Stack operations
type Stack []int

func (s *Stack) Push(v int) {
    *s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
    if len(*s) == 0 {
        return 0, false
    }
    i := len(*s) - 1
    v := (*s)[i]
    *s = (*s)[:i]
    return v, true
}

// Queue operations
type Queue []int

func (q *Queue) Enqueue(v int) {
    *q = append(*q, v)
}

func (q *Queue) Dequeue() (int, bool) {
    if len(*q) == 0 {
        return 0, false
    }
    v := (*q)[0]
    *q = (*q)[1:]
    return v, true
}

func main() {
    // Test Stack
    var stack Stack
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    
    v, _ := stack.Pop()
    fmt.Println("Stack pop:", v)  // 3

    // Test Queue
    var queue Queue
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
    
    v, _ = queue.Dequeue()
    fmt.Println("Queue dequeue:", v)  // 1
}
```

## Utility Functions

```go
// Reverse slice
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// Deduplicate (sorted slice)
func dedupe(s []int) []int {
    if len(s) < 2 {
        return s
    }
    j := 1
    for i := 1; i < len(s); i++ {
        if s[i] != s[i-1] {
            s[j] = s[i]
            j++
        }
    }
    return s[:j]
}
```
