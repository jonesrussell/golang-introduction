# Pass by Value vs Pass by Reference

**Duration:** 7-8 minutes

## Topics to cover:
- [Go is always pass by value](https://go.dev/doc/faq#pass_by_value)
- What "value" means for different types
- Simulating pass by reference with pointers
- When copies happen

## Code Examples

```go runnable
package main

import "fmt"

// Pass by value - function gets a COPY
func doubleValue(n int) {
    n = n * 2
    fmt.Println("Inside function:", n)
}

// Pass by pointer - function can modify original
func doublePointer(n *int) {
    *n = *n * 2
    fmt.Println("Inside function:", *n)
}

func main() {
    // Demonstrate pass by value
    num := 10
    fmt.Println("Before doubleValue:", num)  // 10
    doubleValue(num)                          // Inside function: 20
    fmt.Println("After doubleValue:", num)   // 10 - unchanged!

    // Demonstrate pass by pointer
    num2 := 10
    fmt.Println("Before doublePointer:", num2) // 10
    doublePointer(&num2)                        // Inside function: 20
    fmt.Println("After doublePointer:", num2)  // 20 - changed!
}
```

## Struct Copying Example

```go
type Person struct {
    Name string
    Age  int
}

// Value receiver - works on a copy
func (p Person) Birthday() {
    p.Age++
    fmt.Printf("Inside method: %s is %d\n", p.Name, p.Age)
}

// Pointer receiver - modifies original
func (p *Person) BirthdayPtr() {
    p.Age++
    fmt.Printf("Inside method: %s is %d\n", p.Name, p.Age)
}

func main() {
    alice := Person{Name: "Alice", Age: 30}

    alice.Birthday()
    fmt.Println("After Birthday():", alice.Age)  // 30 - unchanged

    alice.BirthdayPtr()
    fmt.Println("After BirthdayPtr():", alice.Age)  // 31 - changed
}
```

## Key teaching points:
- Go [ALWAYS passes by value](https://go.dev/doc/faq#pass_by_value) (copies the argument)
- Passing a [pointer](https://go.dev/ref/spec#Pointer_types) copies the address, not the data
- Both caller and function can access same memory
- Value semantics prevent accidental mutation
- Choose based on whether mutation is needed
