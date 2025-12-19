# Slices - The Workhorse

**Duration:** 10-12 minutes

## Topics to cover:
- [Slice header](https://go.dev/ref/spec#Slice_types) (pointer, length, capacity)
- Creating slices
- [Slicing operations](https://go.dev/ref/spec#Slice_expressions)
- [`append`](https://pkg.go.dev/builtin#append) and growth
- [`copy`](https://pkg.go.dev/builtin#copy)

## Code Examples

```go runnable
package main

import "fmt"

func main() {
    // Slice is a descriptor: (pointer, length, capacity)
    // type slice struct {
    //     array unsafe.Pointer
    //     len   int
    //     cap   int
    // }

    // Creating slices
    s1 := []int{1, 2, 3}                // Literal
    s2 := make([]int, 5)                // Length 5, capacity 5
    s3 := make([]int, 3, 10)            // Length 3, capacity 10
    var s4 []int                        // nil slice (len=0, cap=0)

    fmt.Println("s1:", s1)
    fmt.Printf("s2: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))
    fmt.Printf("s3: %v (len=%d, cap=%d)\n", s3, len(s3), cap(s3))
    fmt.Printf("s4: %v (nil=%v)\n", s4, s4 == nil)

    // Slice from array
    arr := [5]int{1, 2, 3, 4, 5}
    s5 := arr[1:4]    // [2, 3, 4] - shares backing array!
    fmt.Println("s5 from array:", s5)

    // IMPORTANT: Slices share backing array!
    original := []int{1, 2, 3, 4, 5}
    slice := original[1:4]
    slice[0] = 999
    fmt.Println("original:", original)  // [1, 999, 3, 4, 5] - modified!
}
```

## Slicing Syntax

```go
// Slicing syntax [low:high:max]
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
s1 := s[2:5]      // [2, 3, 4] len=3, cap=8
s2 := s[2:5:5]    // [2, 3, 4] len=3, cap=3 (limited capacity)
s3 := s[:5]       // [0, 1, 2, 3, 4]
s4 := s[5:]       // [5, 6, 7, 8, 9]
s5 := s[:]        // Full slice (copy of header, same array)
```

## Append

```go runnable
package main

import "fmt"

func main() {
    // Append - may create new backing array
    s := []int{1, 2, 3}
    fmt.Printf("Before: len=%d, cap=%d\n", len(s), cap(s))

    s = append(s, 4)  // Might reuse or create new array
    fmt.Printf("After:  len=%d, cap=%d\n", len(s), cap(s))

    s = append(s, 5, 6, 7, 8, 9, 10)  // Definitely new array
    fmt.Printf("Growth: len=%d, cap=%d\n", len(s), cap(s))

    // Append to nil slice works!
    var nilSlice []int
    nilSlice = append(nilSlice, 1, 2, 3)  // Creates backing array
    fmt.Println("nil slice after append:", nilSlice)

    // Append another slice
    s1 := []int{1, 2, 3}
    s2 := []int{4, 5, 6}
    s3 := append(s1, s2...)  // ... unpacks slice
    fmt.Println("combined:", s3)
}
```

## Copy

```go
// Copy - explicit copy of elements
src := []int{1, 2, 3, 4, 5}
dst := make([]int, 3)
n := copy(dst, src)  // Copies min(len(dst), len(src))
fmt.Println(n, dst)  // 3, [1 2 3]

// Safe full copy
original := []int{1, 2, 3, 4, 5}
copied := make([]int, len(original))
copy(copied, original)
// Now modifications to copied don't affect original
```

## Slice Growth Pattern

```go runnable
package main

import "fmt"

func main() {
    // Growth algorithm (approximately):
    // - cap < 256: double
    // - cap >= 256: grow by ~25%

    var s []int
    prevCap := 0

    for i := 0; i < 20; i++ {
        s = append(s, i)
        if cap(s) != prevCap {
            fmt.Printf("len=%2d, cap=%2d\n", len(s), cap(s))
            prevCap = cap(s)
        }
    }
}
```

## Key teaching points:
- [Slice](https://go.dev/ref/spec#Slice_types) is header pointing to array
- [Slicing](https://go.dev/ref/spec#Slice_expressions) creates new header, same array
- [`append`](https://pkg.go.dev/builtin#append) may reallocate (assign result!)
- Pre-allocate with [`make`](https://pkg.go.dev/builtin#make) for known sizes
- [`copy`](https://pkg.go.dev/builtin#copy) for independent slice
