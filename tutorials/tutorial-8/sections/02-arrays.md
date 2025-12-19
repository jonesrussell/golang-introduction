# Arrays - The Foundation

**Duration:** 4-5 minutes

## Topics to cover:
- Fixed size, [value type](https://go.dev/ref/spec#Array_types)
- Declaration and initialization
- When to use arrays (rarely!)

## Code Examples

```go runnable
package main

import "fmt"

func main() {
    // Array declaration - size is part of the type
    var arr1 [5]int                    // Zero values
    arr2 := [5]int{1, 2, 3, 4, 5}      // With values
    arr3 := [...]int{1, 2, 3}          // Size inferred (3)
    arr4 := [5]int{0: 1, 4: 5}         // Sparse: [1, 0, 0, 0, 5]

    fmt.Println("arr1:", arr1)
    fmt.Println("arr2:", arr2)
    fmt.Println("arr3:", arr3)
    fmt.Println("arr4:", arr4)

    // Arrays are values - copied on assignment
    a := [3]int{1, 2, 3}
    b := a           // Full copy!
    b[0] = 100
    fmt.Println("a:", a)   // [1 2 3] - unchanged
    fmt.Println("b:", b)   // [100 2 3]
}
```

## Arrays as Struct Fields

```go
// When arrays make sense:
// - Known fixed size at compile time
// - Part of struct (embedded)
// - Performance-critical with small, fixed data
type Point struct {
    Coords [3]float64  // x, y, z
}
```

## Key teaching points:
- [Arrays](https://go.dev/ref/spec#Array_types) have fixed size (part of type)
- Passed by value (copied)
- Different sizes = different types
- Use [slices](https://go.dev/ref/spec#Slice_types) instead in most cases
