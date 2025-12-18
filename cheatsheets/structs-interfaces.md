# Structs & Interfaces Cheat Sheet

## Struct Definition

```go
type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
}

// With tags
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}
```

## Struct Initialization

```go
// Zero value
var u1 User

// Named fields (preferred)
u2 := User{
    ID:   1,
    Name: "Alice",
}

// Pointer
u3 := &User{
    ID:   2,
    Name: "Bob",
}

// Using new()
u4 := new(User)
u4.ID = 3

// Constructor pattern
func NewUser(name string) *User {
    return &User{
        Name:      name,
        CreatedAt: time.Now(),
    }
}
```

## Methods

```go
// Value receiver (doesn't modify)
func (u User) FullName() string {
    return u.FirstName + " " + u.LastName
}

// Pointer receiver (can modify)
func (u *User) UpdateEmail(email string) {
    u.Email = email
}
```

## When to Use Pointer Receiver

- Method modifies the receiver
- Struct is large (avoid copying)
- Consistency with other methods
- Receiver can be nil (needs to check)

## Embedding

```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person           // Embedded (promoted fields)
    EmployeeID string
    Department string
}

// Usage
e := Employee{
    Person:     Person{Name: "Alice", Age: 30},
    EmployeeID: "E001",
}

fmt.Println(e.Name)  // Promoted from Person
fmt.Println(e.Age)   // Promoted from Person
```

## Interfaces

```go
// Interface definition
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Interface composition
type ReadWriter interface {
    Reader
    Writer
}

// Empty interface (any type)
var anything interface{}
anything = 42
anything = "hello"

// Go 1.18+
var anything any
```

## Type Assertions

```go
var i interface{} = "hello"

// Basic assertion (panics if wrong)
s := i.(string)

// Safe assertion
s, ok := i.(string)
if ok {
    fmt.Println(s)
}

// Type switch
switch v := i.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown type")
}
```

## Interface Satisfaction

```go
// Implicit - no "implements" keyword
type Stringer interface {
    String() string
}

type User struct {
    Name string
}

// User now satisfies Stringer
func (u User) String() string {
    return u.Name
}

// Compile-time check
var _ Stringer = User{}
var _ Stringer = (*User)(nil)
```

## Common Interfaces

```go
// fmt.Stringer
type Stringer interface {
    String() string
}

// error
type error interface {
    Error() string
}

// io.Reader
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer
type Writer interface {
    Write(p []byte) (n int, err error)
}

// io.Closer
type Closer interface {
    Close() error
}
```

## Best Practices

1. **Accept interfaces, return structs**
   ```go
   func NewService(repo Repository) *Service
   ```

2. **Keep interfaces small**
   ```go
   type Saver interface {
       Save(v interface{}) error
   }
   ```

3. **Define interfaces at point of use**
   ```go
   // In consumer package
   type userFinder interface {
       FindByID(id int) (*User, error)
   }
   ```

4. **Use pointer receivers consistently**
   ```go
   func (s *Service) Method1() {}
   func (s *Service) Method2() {}
   func (s *Service) Method3() {}
   ```
