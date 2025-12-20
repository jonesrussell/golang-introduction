# Variables & Declaration

**Duration:** 6-7 minutes

## Topics to cover:
- Variable declaration with [`var`](https://go.dev/ref/spec#Variable_declarations)
- [Type inference](https://go.dev/tour/basics/14)
- [Short declaration](https://go.dev/ref/spec#Short_variable_declarations) (`:=`)
- [Zero values](https://go.dev/ref/spec#The_zero_value)
- [Constants](https://go.dev/ref/spec#Constant_declarations) with `const`
- Multiple variable declaration

## Variable declaration with [`var`](https://go.dev/ref/spec#Variable_declarations)

The `var` keyword declares variables with explicit types. You can optionally provide an initial value.

```go snippet
// Explicit type declaration
var name string = "Russell"
var age int = 30

fmt.Println("Name:", name)
fmt.Println("Age:", age)
```

## [Type inference](https://go.dev/tour/basics/14)

Go can infer the variable type from the initial value, so you can omit the type declaration.

```go snippet
// Type inference - Go determines the type from the value
var city = "Toronto"

fmt.Println("City:", city)
```

## [Short declaration](https://go.dev/ref/spec#Short_variable_declarations) (`:=`)

The short variable declaration (`:=`) is the most common way to declare and initialize variables. It automatically infers the type and can only be used inside functions.

```go snippet
// Short declaration (most common)
country := "Canada"

fmt.Println("Country:", country)
```

## [Zero values](https://go.dev/ref/spec#The_zero_value)

When you declare a variable without an initial value, Go assigns it a "zero value" based on its type. This prevents uninitialized variable bugs.

```go snippet
// Zero values
var count int        // 0
var isActive bool    // false
var message string   // ""

fmt.Println("Count:", count)
fmt.Println("Is Active:", isActive)
fmt.Println("Message:", message)
```

## [Constants](https://go.dev/ref/spec#Constant_declarations) with `const`

Constants are declared with the `const` keyword and must be compile-time values. They cannot be changed after declaration.

```go snippet
// Constants
const MaxRetries = 3
const Pi = 3.14159

fmt.Println("Max Retries:", MaxRetries)
fmt.Println("Pi:", Pi)
```

## Multiple variable declaration

You can declare multiple variables together using parentheses, which is useful for grouping related declarations.

```go snippet
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
