# Hello World & Package Basics

**Duration:** 3-4 minutes

## Topics to cover:
- Creating `main.go`
- `package main` declaration
- `import "fmt"`
- `func main()` as entry point
- `fmt.Println()` for output
- Running with `go run main.go`

## Code Example

```go runnable
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

## Key teaching points:
- Every Go file starts with a package declaration
- `main` package is special - it's executable
- Import standard library packages
- `main()` function is where execution begins
