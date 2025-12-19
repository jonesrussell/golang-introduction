# Control Flow: If Statements

**Duration:** 4-5 minutes

## Topics to cover:
- Basic if/else
- If with initialization statement
- No parentheses needed (Go style)
- Comparison operators

## Code Examples

```go snippet
// Basic if/else
age := 20
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}

// If with initialization
if score := calculateScore(); score > 90 {
    fmt.Println("Excellent!")
} else if score > 70 {
    fmt.Println("Good job!")
} else {
    fmt.Println("Keep practicing!")
}
// score is only available inside if/else block

// Comparison operators
x, y := 10, 20
if x < y {
    fmt.Println("x is less than y")
}
if x != y {
    fmt.Println("x and y are different")
}
```

## Key teaching points:
- No parentheses around condition (Go enforces clean style)
- Braces are mandatory (prevents bugs)
- Init statement scope is limited to if/else block
- Standard comparison operators: `==`, `!=`, `<`, `>`, `<=`, `>=`
