# Control Flow: Loops

**Duration:** 5-6 minutes

## Topics to cover:
- [For loop](https://go.dev/ref/spec#For_statements) (the only loop in Go!)
- [While-style loop](https://go.dev/tour/flowcontrol/3)
- [Infinite loop](https://go.dev/tour/flowcontrol/4)
- [Range over collections](https://go.dev/ref/spec#For_range)
- [Break and continue](https://go.dev/ref/spec#Break_statements)

## Classic [for loop](https://go.dev/ref/spec#For_statements)

The classic for loop has three parts: initialization, condition, and post statement.

```go snippet
// Classic for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## [While-style loop](https://go.dev/tour/flowcontrol/3)

In Go, you use `for` with just a condition to create a while-style loop.

```go snippet
// While-style loop
count := 0
for count < 3 {
    fmt.Println("Count:", count)
    count++
}
```

## [Infinite loop](https://go.dev/tour/flowcontrol/4)

An infinite loop uses `for` without any condition, and you can use `break` to exit.

```go snippet
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
```

## [Range over collections](https://go.dev/ref/spec#For_range)

The `range` keyword allows you to iterate over strings, arrays, slices, and maps.

```go snippet
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
- Go [only has `for`](https://go.dev/ref/spec#For_statements) - no `while` or `do-while`
- Different forms of [`for`](https://go.dev/tour/flowcontrol/1) cover all loop needs
- [`range`](https://go.dev/ref/spec#For_range) is idiomatic for iterating
- Use [blank identifier `_`](https://go.dev/ref/spec#Blank_identifier) to ignore values you don't need
- [`break`](https://go.dev/ref/spec#Break_statements) exits loop, [`continue`](https://go.dev/ref/spec#Continue_statements) skips to next iteration
