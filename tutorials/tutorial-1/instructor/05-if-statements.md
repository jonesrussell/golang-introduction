# Instructor Notes: Control Flow - If Statements

## Teaching Techniques
- Compare Go style to C/Java (no parentheses)
- Show the init statement feature - unique to Go
- Demo scope limitations

## Demo Flow
1. Basic if/else
2. if/else if/else chain
3. if with init statement
4. Show scope: variable not available outside block

## Key Emphasis
- No parentheses around condition
- Braces are MANDATORY (show error without them)
- Init statement keeps scope clean

## Code to Type Live
```go
// This won't compile - show the error
if age >= 18
    fmt.Println("Adult")

// Correct version
if age >= 18 {
    fmt.Println("Adult")
}
```

## Engagement
- Ask: "What does this print?" before running
- Challenge: Write an if statement with init
- "Pause and try this yourself"

## Comparison Operators Reference
- `==` equals
- `!=` not equals
- `<` `>` `<=` `>=`
- `&&` and
- `||` or
- `!` not
