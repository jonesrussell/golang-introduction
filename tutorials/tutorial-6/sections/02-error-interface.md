# The Error Interface

**Duration:** 4-5 minutes

## Topics to cover:
- `error` is just an interface
- The `Error()` method
- Creating errors with `errors.New` and `fmt.Errorf`
- Zero value of error is `nil`

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
- `error` is an interface with single `Error() string` method
- Any type with `Error() string` is an error
- `errors.New()` creates simple errors
- `fmt.Errorf()` creates formatted errors
- `nil` means no error
