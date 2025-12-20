# Variables & Declaration

**Duration:** 6-7 minutes

## Topics to cover:
- Variable declaration with [`var`](https://go.dev/ref/spec#Variable_declarations)
- [Type inference](https://go.dev/tour/basics/14)
- [Short declaration](https://go.dev/ref/spec#Short_variable_declarations) (`:=`)
- [Zero values](https://go.dev/ref/spec#The_zero_value)
- [Constants](https://go.dev/ref/spec#Constant_declarations) with `const`
- Multiple variable declaration

## Code Examples

```go snippet
// Explicit type declaration
var name string = "Russell"
var age int = 30

fmt.Println("Name:", name)
fmt.Println("Age:", age)

// Type inference
var city = "Toronto"

fmt.Println("City:", city)

// Short declaration (most common)
country := "Canada"

fmt.Println("Country:", country)

// Zero values
var count int        // 0
var isActive bool    // false
var message string   // ""

fmt.Println("Count:", count)
fmt.Println("Is Active:", isActive)
fmt.Println("Message:", message)

// Constants
const MaxRetries = 3
const Pi = 3.14159

fmt.Println("Max Retries:", MaxRetries)
fmt.Println("Pi:", Pi)

// Multiple declaration
var (
    firstName string = "John"
    lastName  string = "Doe"
    score     int    = 95
)

fmt.Println("First Name:", firstName)
fmt.Println("Last Name:", lastName)
fmt.Println("Score:", score)
```

## Key teaching points:
- [`:=` can only be used inside functions](https://go.dev/ref/spec#Short_variable_declarations)
- Go is [statically typed](https://go.dev/doc/faq#Is_Go_an_object-oriented_language) but has [type inference](https://go.dev/tour/basics/14)
- [Unused variables are compilation errors](https://go.dev/doc/faq#unused_variables_and_imports) (good for code quality!)
- [Zero values](https://go.dev/ref/spec#The_zero_value) prevent uninitialized variable bugs
- [Constants must be compile-time values](https://go.dev/ref/spec#Constants)
