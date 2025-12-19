# Instructor Notes: Basic Types

## Teaching Techniques
- Compare to other languages briefly
- Show fmt.Printf for type inspection
- Demo type conversion explicitly

## Demo Flow
1. Show numeric types with examples
2. Demo strings - regular and backtick multiline
3. Show booleans
4. IMPORTANT: Demo type conversion error first, then fix

## Key Emphasis
- No implicit conversion - this is intentional!
- Show the error: `var y float64 = x` where x is int
- Then show the fix: `var y float64 = float64(x)`

## Printf Patterns to Show
```go
fmt.Printf("Type: %T, Value: %v\n", x, x)
fmt.Printf("String: %s\n", str)
fmt.Printf("Number: %d\n", num)
fmt.Printf("Float: %.2f\n", price)
```

## Engagement
- Ask: "Why do you think Go requires explicit conversion?"
- Answer: Prevents subtle bugs from implicit casting
- "Notice how Go forces us to be explicit"

## Real-World Context
- int for counts, indexes
- float64 for money (or use int cents)
- string for text
- bool for flags
