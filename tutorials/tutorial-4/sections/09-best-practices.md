# Best Practices Summary

**Duration:** 3-4 minutes

## Guidelines

```go
// POINTER RECEIVER GUIDELINES

// Use pointer receiver when:
// 1. Method modifies receiver
func (u *User) UpdateEmail(email string) {
    u.Email = email
}

// 2. Receiver is large struct
func (ls *LargeStruct) Process() {}

// 3. Consistency - if any method uses pointer, all should
type Client struct { /* ... */ }
func (c *Client) Connect() error { /* ... */ }
func (c *Client) Disconnect() error { /* ... */ }  // Pointer for consistency
func (c *Client) IsConnected() bool { /* ... */ }  // Pointer for consistency

// FUNCTION PARAMETER GUIDELINES

// Use pointer parameter when:
// 1. Function needs to modify argument
func LoadConfig(cfg *Config) error {
    // Populate cfg fields
}

// 2. Nil is a valid input (optional parameter)
func Process(opts *Options) {
    if opts == nil {
        opts = DefaultOptions()
    }
}

// RETURN VALUE GUIDELINES

// Return pointer when:
// 1. Nil is meaningful (not found, error case)
func FindUser(id int) *User {
    // Returns nil if not found
}

// 2. Creating large objects
func NewLargeObject() *LargeObject {
    return &LargeObject{/* ... */}
}

// Return value when:
// 1. Small, simple types
func NewPoint(x, y float64) Point {
    return Point{X: x, Y: y}
}

// 2. Immutable data
func (p Point) Translate(dx, dy float64) Point {
    return Point{X: p.X + dx, Y: p.Y + dy}
}

// INITIALIZATION PATTERNS

// Preferred: &Type{} with values
user := &User{
    ID:   1,
    Name: "Alice",
}

// OK: new() when zero values are fine
buffer := new(bytes.Buffer)

// Constructor pattern
func NewUser(id int, name string) *User {
    return &User{
        ID:        id,
        Name:      name,
        CreatedAt: time.Now(),
    }
}
```

## Key teaching points:
- Be consistent with [receiver types](https://go.dev/doc/effective_go#pointers_vs_values)
- Document when [nil is valid](https://go.dev/ref/spec#The_zero_value)
- Prefer [`&Type{}`](https://go.dev/ref/spec#Address_operators) over [`new(Type)`](https://go.dev/ref/spec#Allocation)
- Use constructors for initialization with defaults
