# Variables & Declaration

**Duration:** 6-7 minutes

## Topics to cover:
- Variable declaration with `var`
- Type inference
- Short declaration (`:=`)
- Zero values
- Constants with `const`
- Multiple variable declaration

## Code Examples

```go snippet
// Explicit type declaration
var name string = "Russell"
var age int = 30

// Type inference
var city = "Toronto"

// Short declaration (most common)
country := "Canada"

// Zero values
var count int        // 0
var isActive bool    // false
var message string   // ""

// Constants
const MaxRetries = 3
const Pi = 3.14159

// Multiple declaration
var (
    firstName string = "John"
    lastName  string = "Doe"
    score     int    = 95
)
```

## Key teaching points:
- `:=` can only be used inside functions
- Go is statically typed but has type inference
- Unused variables are compilation errors (good for code quality!)
- Zero values prevent uninitialized variable bugs
- Constants must be compile-time values
