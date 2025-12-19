# The Empty Interface

**Duration:** 4-5 minutes

## Topics to cover:
- [`interface{}`](https://go.dev/ref/spec#Interface_types) and [`any`](https://go.dev/ref/spec#Predeclared_identifiers) (Go 1.18+)
- Why it accepts any type
- [Type assertions](https://go.dev/ref/spec#Type_assertions)
- When to use (and not use) empty interface

## Code Examples

```go
// Empty interface - has zero methods
// Every type has at least zero methods, so everything satisfies it
func PrintAnything(v interface{}) {
    fmt.Println(v)
}

// Go 1.18+: 'any' is alias for interface{}
func PrintAny(v any) {
    fmt.Println(v)
}

func main() {
    PrintAnything(42)
    PrintAnything("hello")
    PrintAnything([]int{1, 2, 3})
    PrintAnything(struct{ X int }{X: 10})

    // Type assertion - extract concrete type
    var i interface{} = "hello"

    // Basic assertion (panics if wrong type)
    s := i.(string)
    fmt.Println(s)  // hello

    // Safe assertion with comma-ok
    s, ok := i.(string)
    if ok {
        fmt.Println("It's a string:", s)
    }

    n, ok := i.(int)
    if !ok {
        fmt.Println("Not an int")
    }

    // Type switch - handle multiple types
    describe(42)
    describe("hello")
    describe(true)
    describe([]int{1, 2})
}

func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

## When to Use Empty Interface

```go
// GOOD: Generic containers (before Go 1.18 generics)
type Cache struct {
    data map[string]interface{}
}

// GOOD: JSON unmarshaling when structure is unknown
var data interface{}
json.Unmarshal([]byte(`{"key": "value"}`), &data)

// GOOD: Printf-style variadic functions
func Log(format string, args ...interface{}) {
    fmt.Printf(format, args...)
}

// BAD: Avoid when you know the type
func ProcessUser(u interface{}) {  // BAD - just use User type!
    user := u.(User)
    // ...
}

// Since Go 1.18, prefer generics over interface{} where applicable
func First[T any](items []T) T {
    return items[0]
}
```

## Key teaching points:
- [`interface{}`](https://go.dev/ref/spec#Interface_types) / [`any`](https://go.dev/ref/spec#Predeclared_identifiers) accepts any type
- Use [type assertions](https://go.dev/ref/spec#Type_assertions) to get concrete type back
- Comma-ok pattern prevents panics
- [Type switch](https://go.dev/ref/spec#Type_switches) for multiple type handling
- Prefer specific interfaces over `interface{}`
- Go 1.18 [generics](https://go.dev/doc/tutorial/generics) often better than `interface{}`
