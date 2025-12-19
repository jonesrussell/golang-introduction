# Interface Satisfaction Rules

**Duration:** 5-6 minutes

## Topics to cover:
- [Method sets](https://go.dev/ref/spec#Method_sets) and receivers
- [Pointer vs value receiver](https://go.dev/ref/spec#Method_sets) implications
- All methods must match exactly

## Code Examples

```go
type Writer interface {
    Write(data []byte) error
}

// Value receiver - both Type and *Type satisfy interface
type Buffer struct {
    data []byte
}

func (b Buffer) Write(data []byte) error {
    // Note: can't actually modify b.data here (value receiver)
    return nil
}

var _ Writer = Buffer{}   // OK
var _ Writer = &Buffer{}  // OK

// Pointer receiver - only *Type satisfies interface
type File struct {
    path string
}

func (f *File) Write(data []byte) error {
    // Can modify f here (pointer receiver)
    return nil
}

// var _ Writer = File{}   // ERROR! File doesn't have Write method
var _ Writer = &File{}     // OK

// Rule: If method has pointer receiver, must use pointer to satisfy interface
// Rule: If method has value receiver, both value and pointer work

// Method signature must match EXACTLY
type Processor interface {
    Process(input string) (string, error)
}

type MyProcessor struct{}

// This satisfies Processor:
func (p MyProcessor) Process(input string) (string, error) {
    return input, nil
}

// This does NOT satisfy Processor (different signature):
// func (p MyProcessor) Process(input string) string { ... }
```

## Common Mistake

```go
type Stringer interface {
    String() string
}

type User struct {
    Name string
}

// Pointer receiver
func (u *User) String() string {
    return u.Name
}

func PrintString(s Stringer) {
    fmt.Println(s.String())
}

func main() {
    user := User{Name: "Alice"}

    // PrintString(user)   // ERROR: User doesn't implement Stringer
    PrintString(&user)     // OK: *User implements Stringer
}
```

## Key teaching points:
- [Value receiver](https://go.dev/ref/spec#Method_sets): Type and *Type both satisfy
- [Pointer receiver](https://go.dev/ref/spec#Method_sets): only *Type satisfies
- All methods must have exact matching signatures
- Use compile-time check: `var _ Interface = Type{}`
