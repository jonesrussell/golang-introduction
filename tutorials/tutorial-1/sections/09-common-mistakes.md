# Common Beginner Mistakes

**Duration:** 3-4 minutes

## Cover these pitfalls:

```go snippet
// 1. Unused variables (compilation error)
func badExample() {
    x := 10  // Declared but never used - won't compile!
}

// 2. Shadowing variables
count := 5
if true {
    count := 10  // This creates a NEW variable!
    fmt.Println(count)  // Prints 10
}
fmt.Println(count)  // Still prints 5

// 3. Wrong scope with :=
var err error
if data, err := getData(); err != nil {  // Creates NEW err!
    return err
}
// Original err is still nil here

// Better:
var data string
data, err = getData()  // Use existing err
if err != nil {
    return err
}

// 4. Implicit type conversion
var x int = 10
var y float64 = x  // ERROR! Must use float64(x)
```

## Key teaching points:
- Always use declared variables ([compiler enforces this](https://go.dev/doc/faq#unused_variables_and_imports))
- Be careful with [`:=` creating new variables](https://go.dev/doc/faq#shadowing) in inner scopes
- [Shadowing](https://go.dev/doc/faq#shadowing) can lead to subtle bugs
- Go requires [explicit type conversion](https://go.dev/ref/spec#Conversions)
