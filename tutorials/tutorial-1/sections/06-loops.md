# Control Flow: Loops

**Duration:** 5-6 minutes

## Topics to cover:
- For loop (the only loop in Go!)
- While-style loop
- Infinite loop
- Range over collections
- Break and continue

## Code Examples

```go snippet
// Classic for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-style loop
count := 0
for count < 3 {
    fmt.Println("Count:", count)
    count++
}

// Infinite loop (with break)
counter := 0
for {
    counter++
    if counter > 5 {
        break
    }
    if counter == 3 {
        continue  // Skip to next iteration
    }
    fmt.Println(counter)
}

// Range over string
name := "Go"
for index, char := range name {
    fmt.Printf("Index %d: %c\n", index, char)
}

// Ignore index with _
for _, char := range name {
    fmt.Printf("%c ", char)
}
```

## Key teaching points:
- Go only has `for` - no `while` or `do-while`
- Different forms of `for` cover all loop needs
- `range` is idiomatic for iterating
- Use `_` to ignore values you don't need
- `break` exits loop, `continue` skips to next iteration
