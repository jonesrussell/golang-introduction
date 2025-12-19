# Common Interface Mistakes

**Duration:** 4-5 minutes

## Topics to cover:
- Interface pollution
- Premature abstraction
- Type assertion abuse

## Code Examples

```go
// MISTAKE 1: Interface pollution (too many interfaces)

// BAD: Every type has its own interface
type UserInterface interface {
    GetName() string
}

type User struct {
    Name string
}

func (u User) GetName() string {
    return u.Name
}

// Just use the concrete type!
func ProcessUser(u User) {  // Not UserInterface
    fmt.Println(u.Name)
}

// MISTAKE 2: Premature abstraction

// BAD: Creating interface before you need it
type Logger interface {
    Log(msg string)
}

// You only have one implementation!
type FileLogger struct{}

func (f FileLogger) Log(msg string) { /* ... */ }

// Don't create interface until you need multiple implementations
// or need to mock for testing

// MISTAKE 3: Returning interface instead of concrete type

// BAD: Returns interface
func NewService() ServiceInterface {
    return &service{}
}

// GOOD: Returns concrete type
func NewService() *Service {
    return &Service{}
}

// MISTAKE 4: Accepting concrete when interface would work

// BAD: Only works with this specific type
func ProcessFile(f *os.File) error {
    // reads from file
}

// GOOD: Works with any reader
func ProcessReader(r io.Reader) error {
    // reads from any source
}

// MISTAKE 5: Interface{} abuse

// BAD: Using interface{} when type is known
func BadProcess(data interface{}) {
    user := data.(User)  // Why not just accept User?
    // ...
}

// GOOD: Use actual type
func GoodProcess(user User) {
    // ...
}

// MISTAKE 6: Large interfaces

// BAD: Forces implementations to have many methods
type Repository interface {
    Create(v interface{}) error
    Read(id int) (interface{}, error)
    Update(v interface{}) error
    Delete(id int) error
    List() ([]interface{}, error)
    Count() (int, error)
    Search(query string) ([]interface{}, error)
    // ... more methods
}

// GOOD: Small, focused interfaces
type Creator interface {
    Create(v interface{}) error
}

type Reader interface {
    Read(id int) (interface{}, error)
}

// Compose when needed
type ReadWriter interface {
    Reader
    Creator
}
```

## Key teaching points:
- Don't create interfaces for single implementations
- Wait until you need the abstraction
- Return concrete types, accept interfaces
- Keep interfaces small
- Avoid `interface{}` when you know the type
- Don't wrap every type in an interface
