# What Are Pointers?

**Duration:** 5-6 minutes

## Topics to cover:
- Memory addresses explained
- Pointer declaration and syntax
- The `&` operator (address-of)
- The `*` operator (dereference)
- Zero value of pointers (`nil`)

## Code Examples

```go runnable
package main

import "fmt"

func main() {
    // A variable stores a value
    x := 42
    fmt.Println("Value of x:", x)           // 42
    fmt.Println("Address of x:", &x)        // 0xc0000b4008 (memory address)

    // A pointer stores a memory address
    var p *int = &x                         // p points to x
    fmt.Println("Value of p:", p)           // 0xc0000b4008
    fmt.Println("Value at p:", *p)          // 42 (dereference)

    // Modify value through pointer
    *p = 100
    fmt.Println("x is now:", x)             // 100

    // Short declaration
    y := 50
    ptr := &y
    fmt.Println(*ptr)                       // 50

    // Zero value of pointer is nil
    var nilPtr *int
    fmt.Println("nil pointer:", nilPtr)     // <nil>

    // DANGER: Dereferencing nil pointer causes panic
    // fmt.Println(*nilPtr)  // PANIC: invalid memory address

    // Safe nil check
    if nilPtr != nil {
        fmt.Println(*nilPtr)
    }
}
```

## Key teaching points:
- `&` gives the address of a variable
- `*` in type declaration creates a pointer type
- `*` before pointer variable dereferences it
- Nil is the zero value - always check before dereferencing
- Pointers enable indirect modification
