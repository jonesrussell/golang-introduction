# Common Beginner Mistakes

**Duration:** 3-4 minutes

## Unused variables (compilation error)

Go requires all declared variables to be used. This prevents accidental dead code and improves code quality.

```go snippet
// Unused variables cause compilation errors
x := 10  // Declared but never used - won't compile!
// Use it or remove it
fmt.Println(x)
```

## Shadowing variables

When you use `:=` in an inner scope with a variable name that already exists, it creates a new variable instead of reassigning. This is called shadowing and can lead to bugs.

```go snippet
// Shadowing variables
count := 5
if true {
    count := 10  // This creates a NEW variable!
    fmt.Println(count)  // Prints 10
}
fmt.Println(count)  // Still prints 5 (original variable unchanged)
```

## Wrong scope with `:=`

The short declaration operator `:=` always creates new variables, even if variables with the same name exist in outer scopes.

```go snippet
// Wrong: Creates new err variable
var err error
if data, err := getData(); err != nil {  // Creates NEW err!
    return err
}
// Original err is still nil here

// Better: Use assignment instead
var data string
var err error
data, err = getData()  // Use existing err variable
if err != nil {
    return err
}
```

## Implicit type conversion

Go requires explicit type conversion - there's no automatic type coercion, which prevents subtle bugs.

```go snippet
// Implicit type conversion is not allowed
var x int = 10
// var y float64 = x  // ERROR! Must use explicit conversion
var y float64 = float64(x)  // Correct way

fmt.Printf("x: %d, y: %.2f\n", x, y)
```

## Key teaching points:
- Always use declared variables ([compiler enforces this](https://go.dev/doc/faq#unused_variables_and_imports))
- Be careful with [`:=` creating new variables](https://go.dev/doc/faq#shadowing) in inner scopes
- [Shadowing](https://go.dev/doc/faq#shadowing) can lead to subtle bugs
- Go requires [explicit type conversion](https://go.dev/ref/spec#Conversions)
