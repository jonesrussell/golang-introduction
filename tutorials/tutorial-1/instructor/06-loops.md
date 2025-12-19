# Instructor Notes: Control Flow - Loops

## Teaching Techniques
- Emphasize: Go only has `for` - no while/do-while
- Show how `for` covers all cases
- Demo range keyword importance

## Demo Flow
1. Classic three-part for loop
2. While-style (condition only)
3. Infinite loop with break
4. Range over string
5. Show continue keyword

## Key Emphasis
- "Go only has for - but it's powerful enough"
- Range is idiomatic - use it!
- Blank identifier `_` to ignore values

## Live Coding Sequence
```go
// Start simple
for i := 0; i < 3; i++ {
    fmt.Println(i)
}

// Then while-style
count := 0
for count < 3 {
    count++
}

// Then infinite with break
for {
    // dangerous without break!
    break
}
```

## Engagement
- Ask: "What happens if we forget break in infinite loop?"
- Challenge: "Write a loop that prints 1 to 10"
- "Pause here and try this yourself"

## Common Mistakes to Highlight
- Off-by-one errors
- Forgetting break in infinite loops
- Modifying loop variable incorrectly
