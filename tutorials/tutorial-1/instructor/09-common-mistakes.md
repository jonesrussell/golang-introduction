# Instructor Notes: Common Beginner Mistakes

## Teaching Techniques
- Show each mistake FIRST, then the fix
- Let the error message display
- Explain WHY Go enforces these rules

## Demo Sequence

### 1. Unused Variables
```go
x := 10  // Error: x declared but not used
```
- Run and show error
- "Go won't let you have dead code"

### 2. Variable Shadowing
```go
count := 5
if true {
    count := 10  // NEW variable!
}
fmt.Println(count)  // Still 5!
```
- This is sneaky - show it carefully
- "This caused me bugs when I was learning"

### 3. Wrong Scope with :=
```go
var err error
if data, err := getData(); err != nil {
    // This err is different!
}
// Original err unchanged
```
- Very common gotcha
- Show the fix with `=` instead of `:=`

### 4. Implicit Type Conversion
```go
var x int = 10
var y float64 = x  // Error!
```
- Show error, then fix with `float64(x)`

## Engagement
- "I made this mistake many times when learning"
- "The compiler is trying to help you"
- "Notice how Go forces us to be explicit"

## Key Message
- These aren't restrictions - they're safety features
- The compiler catches bugs before they happen
- Embrace Go's strictness
