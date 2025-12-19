# Hello World & Package Basics

**Duration:** 3-4 minutes

## Topics to cover:
- Creating `main.go`
- [`package main`](https://go.dev/doc/code#Organization) declaration
- [`import "fmt"`](https://pkg.go.dev/fmt)
- [`func main()`](https://go.dev/ref/spec#Program_initialization_and_execution) as entry point
- [`fmt.Println()`](https://pkg.go.dev/fmt#Println) for output
- Running with [`go run`](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)

## Code Example

```go runnable
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

## Key teaching points:
- Every Go file starts with a [package declaration](https://go.dev/ref/spec#Package_clause)
- [`main` package](https://go.dev/doc/code#Command) is special - it's executable
- [Import](https://go.dev/ref/spec#Import_declarations) standard library packages
- [`main()` function](https://go.dev/ref/spec#Program_initialization_and_execution) is where execution begins
