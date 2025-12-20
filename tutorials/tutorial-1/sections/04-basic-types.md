# Basic Types

**Duration:** 5-6 minutes

## Topics to cover:
- Numeric types: [`int`](https://go.dev/ref/spec#Numeric_types), [`int64`](https://go.dev/ref/spec#Numeric_types), [`float64`](https://go.dev/ref/spec#Numeric_types)
- [String type](https://go.dev/ref/spec#String_types)
- [Boolean type](https://go.dev/ref/spec#Boolean_types)
- [Type conversion](https://go.dev/ref/spec#Conversions) (explicit only)
- String concatenation

## Numeric types

Go provides several numeric types including `int`, `int64`, and `float64`. The `int` type is platform-dependent, while `int64` is always 64 bits.

```go snippet
// Numeric types
var count int = 42
var price float64 = 19.99
var distance int64 = 1000000

fmt.Println("Count:", count)
fmt.Println("Price:", price)
fmt.Println("Distance:", distance)
```

## [String type](https://go.dev/ref/spec#String_types)

Strings in Go are sequences of bytes and support both regular strings and raw string literals using backticks.

```go snippet
// Strings
message := "Learning Go"
multiLine := `This is a
multi-line string
using backticks`

fmt.Println(message)
fmt.Println(multiLine)
```

## [Boolean type](https://go.dev/ref/spec#Boolean_types)

Booleans represent truth values: `true` or `false`.

```go snippet
// Booleans
isComplete := true
hasError := false

fmt.Println("Complete:", isComplete)
fmt.Println("Has error:", hasError)
```

## [Type conversion](https://go.dev/ref/spec#Conversions) (explicit only)

Go requires explicit type conversion - there's no automatic type coercion to prevent bugs.

```go snippet
// Type conversion (explicit)
var x int = 10
var y float64 = float64(x)  // Must convert explicitly
// var z float64 = x  // This would be an error!

fmt.Printf("x: %d, y: %.2f\n", x, y)
```

## String concatenation

Strings can be concatenated using the `+` operator, and you can use `fmt.Printf` for formatted output.

```go snippet
// String operations
firstName := "Jane"
lastName := "Smith"
fullName := firstName + " " + lastName
fmt.Printf("Name: %s, Length: %d\n", fullName, len(fullName))
```

## Key teaching points:
- [No implicit type conversion](https://go.dev/ref/spec#Conversions) (prevents bugs)
- [`int` vs `int64`](https://go.dev/ref/spec#Numeric_types) - platform-dependent vs explicit size
- [String concatenation](https://go.dev/ref/spec#String_concatenation) with `+`
- [`fmt.Printf`](https://pkg.go.dev/fmt#Printf) for formatted output
- [Raw string literals](https://go.dev/ref/spec#String_literals) (backticks) for multi-line strings
