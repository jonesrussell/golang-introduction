# The Error Interface

**Duration:** 4-5 minutes

## Topics to cover:
- [`error`](https://pkg.go.dev/builtin#error) is just an interface
- The [`Error()`](https://pkg.go.dev/builtin#error) method
- Creating errors with [`errors.New`](https://pkg.go.dev/errors#New) and [`fmt.Errorf`](https://pkg.go.dev/fmt#Errorf)
- [Zero value](https://go.dev/ref/spec#The_zero_value) of error is `nil`

## Code Examples

```go runnable
package main

import (
    "errors"
    "fmt"
)

// error is a built-in interface
// type error interface {
//     Error() string
// }

func main() {
    // Creating errors
    err1 := errors.New("something went wrong")
    fmt.Println(err1.Error())  // something went wrong
    fmt.Println(err1)          // same - fmt knows about Error()

    // With formatting
    name := "config.json"
    err2 := fmt.Errorf("failed to open file: %s", name)
    fmt.Println(err2)  // failed to open file: config.json

    // Error is nil when no error occurred
    var err error  // Zero value is nil
    if err == nil {
        fmt.Println("No error")
    }

    // Checking for errors
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

## Key teaching points:
- [`error`](https://pkg.go.dev/builtin#error) is an interface with single [`Error() string`](https://pkg.go.dev/builtin#error) method
- Any type with [`Error() string`](https://pkg.go.dev/builtin#error) is an error
- [`errors.New()`](https://pkg.go.dev/errors#New) creates simple errors
- [`fmt.Errorf()`](https://pkg.go.dev/fmt#Errorf) creates formatted errors
- [`nil`](https://go.dev/ref/spec#The_zero_value) means no error
