# Basic Types

**Duration:** 5-6 minutes

## Topics to cover:
- Numeric types: `int`, `int64`, `float64`
- String type
- Boolean type
- Type conversion (explicit only)
- String concatenation

## Code Examples

```go snippet
// Numeric types
var count int = 42
var price float64 = 19.99
var distance int64 = 1000000

// Strings
message := "Learning Go"
multiLine := `This is a
multi-line string
using backticks`

// Booleans
isComplete := true
hasError := false

// Type conversion (explicit)
var x int = 10
var y float64 = float64(x)  // Must convert explicitly
// var z float64 = x  // This would be an error!

// String operations
firstName := "Jane"
lastName := "Smith"
fullName := firstName + " " + lastName
fmt.Printf("Name: %s, Length: %d\n", fullName, len(fullName))
```

## Key teaching points:
- No implicit type conversion (prevents bugs)
- `int` vs `int64` - platform-dependent vs explicit size
- String concatenation with `+`
- `fmt.Printf` for formatted output
- Backticks for raw/multi-line strings
