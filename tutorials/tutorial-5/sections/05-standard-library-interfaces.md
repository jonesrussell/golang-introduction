# Standard Library Interfaces

**Duration:** 6-7 minutes

## Topics to cover:
- [`io.Reader`](https://pkg.go.dev/io#Reader) and [`io.Writer`](https://pkg.go.dev/io#Writer)
- [`fmt.Stringer`](https://pkg.go.dev/fmt#Stringer)
- [`error` interface](https://pkg.go.dev/builtin#error)
- Why small interfaces matter

## Code Examples

```go
import (
    "fmt"
    "io"
    "strings"
)

// io.Reader - one of the most important interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Custom type implementing io.Reader
type RepeatReader struct {
    char  byte
    count int
    read  int
}

func (r *RepeatReader) Read(p []byte) (n int, err error) {
    if r.read >= r.count {
        return 0, io.EOF
    }

    toRead := r.count - r.read
    if toRead > len(p) {
        toRead = len(p)
    }

    for i := 0; i < toRead; i++ {
        p[i] = r.char
    }
    r.read += toRead

    return toRead, nil
}

func main() {
    // Our RepeatReader works with any function expecting io.Reader
    reader := &RepeatReader{char: 'A', count: 10}
    data, _ := io.ReadAll(reader)
    fmt.Println(string(data))  // AAAAAAAAAA

    // strings.Reader implements io.Reader
    sr := strings.NewReader("Hello, World!")
    io.Copy(os.Stdout, sr)
}
```

## fmt.Stringer Interface

```go
// fmt.Stringer - custom string representation
type Stringer interface {
    String() string
}

type User struct {
    ID   int
    Name string
}

func (u User) String() string {
    return fmt.Sprintf("User#%d: %s", u.ID, u.Name)
}

func main() {
    user := User{ID: 1, Name: "Alice"}
    fmt.Println(user)  // User#1: Alice (uses String() method)
}
```

## error Interface

```go
// error interface - extremely simple
type error interface {
    Error() string
}

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

func ValidateUser(name string) error {
    if name == "" {
        return ValidationError{Field: "name", Message: "cannot be empty"}
    }
    return nil
}
```

## Interface Composition

```go
// Composing interfaces
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// This is how standard library defines io.ReadWriteCloser!
```

## Key teaching points:
- Small interfaces are powerful ([`Reader`](https://pkg.go.dev/io#Reader) has 1 method)
- [Composition](https://go.dev/ref/spec#Interface_types) creates larger interfaces
- [`fmt.Stringer`](https://pkg.go.dev/fmt#Stringer) customizes print output
- [`error`](https://pkg.go.dev/builtin#error) is just an interface with one method
- Standard library interfaces enable ecosystem interop
