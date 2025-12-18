# Go Basics Cheat Sheet

## Variables

```go
// Declaration with type
var name string = "Alice"
var age int = 30

// Type inference
var city = "Toronto"

// Short declaration (inside functions)
country := "Canada"

// Multiple declaration
var (
    firstName string = "John"
    lastName  string = "Doe"
    score     int    = 95
)

// Constants
const MaxRetries = 3
const Pi = 3.14159
```

## Zero Values

| Type | Zero Value |
|------|------------|
| `int`, `float64` | `0` |
| `string` | `""` |
| `bool` | `false` |
| `pointer`, `slice`, `map`, `chan`, `func` | `nil` |

## Basic Types

```go
// Numeric
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
complex64, complex128

// String
string

// Boolean
bool
```

## Type Conversion

```go
var x int = 10
var y float64 = float64(x)  // Explicit conversion required
var z int = int(y)
```

## Control Flow

### If Statement
```go
if age >= 18 {
    fmt.Println("Adult")
} else if age >= 13 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Child")
}

// With initialization
if score := calculateScore(); score > 90 {
    fmt.Println("Excellent!")
}
```

### For Loop
```go
// Classic for
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-style
count := 0
for count < 3 {
    count++
}

// Infinite loop
for {
    // break to exit
}

// Range
for index, value := range slice {
    fmt.Println(index, value)
}

// Ignore index
for _, value := range slice {
    fmt.Println(value)
}
```

### Switch
```go
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Midweek")
}

// Without expression
switch {
case hour < 12:
    fmt.Println("Morning")
case hour < 17:
    fmt.Println("Afternoon")
default:
    fmt.Println("Evening")
}
```

## Functions

```go
// Basic function
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func rectangle(width, height int) (area, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return  // Naked return
}

// Variadic function
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

## Printf Format Verbs

| Verb | Description |
|------|-------------|
| `%v` | Default format |
| `%+v` | Include field names (structs) |
| `%#v` | Go syntax representation |
| `%T` | Type |
| `%d` | Integer |
| `%f` | Float |
| `%s` | String |
| `%t` | Boolean |
| `%p` | Pointer |
| `%%` | Literal % |

## Quick Commands

```bash
go run main.go      # Run
go build            # Compile
go fmt ./...        # Format code
go vet ./...        # Check for errors
go test ./...       # Run tests
go mod init <name>  # Initialize module
go mod tidy         # Clean dependencies
```
