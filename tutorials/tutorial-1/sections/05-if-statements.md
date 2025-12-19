# Control Flow: If Statements

**Duration:** 4-5 minutes

## Topics to cover:
- Basic [if/else](https://go.dev/ref/spec#If_statements)
- [If with initialization statement](https://go.dev/tour/flowcontrol/6)
- No parentheses needed (Go style)
- [Comparison operators](https://go.dev/ref/spec#Comparison_operators)

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
- [No parentheses around condition](https://go.dev/ref/spec#If_statements) (Go enforces clean style)
- Braces are [mandatory](https://go.dev/ref/spec#Blocks) (prevents bugs)
- [Init statement scope](https://go.dev/tour/flowcontrol/6) is limited to if/else block
- Standard [comparison operators](https://go.dev/ref/spec#Comparison_operators): `==`, `!=`, `<`, `>`, `<=`, `>=`
