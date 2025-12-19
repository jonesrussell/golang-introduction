# Best Practices & Code Smells

**Duration:** 5-6 minutes

## Good vs Bad Patterns

```go
// ✅ GOOD: Constructor with validation
func NewProduct(name string, price float64) (*Product, error) {
    if price < 0 {
        return nil, fmt.Errorf("price cannot be negative")
    }
    return &Product{Name: name, Price: price}, nil
}

// ❌ BAD: No validation, invalid state possible
func NewProduct(name string, price float64) *Product {
    return &Product{Name: name, Price: price}  // Could have negative price!
}

// ✅ GOOD: Consistent receiver types (all pointers)
type Counter struct {
    count int
}

func (c *Counter) Increment() { c.count++ }
func (c *Counter) Decrement() { c.count-- }
func (c *Counter) Value() int { return c.count }  // Even for read-only!

// ❌ BAD: Mixed receiver types
func (c *Counter) Increment() { c.count++ }
func (c Counter) Value() int { return c.count }  // Inconsistent!

// ✅ GOOD: Named field initialization
person := Person{
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
}

// ❌ BAD: Positional initialization (brittle)
person := Person{"John", "Doe", 30}

// ✅ GOOD: Pointer for structs that will be modified
user := &User{Name: "Alice"}
user.Activate()

// ✅ GOOD: Value for small, immutable structs
point := Point{X: 10, Y: 20}
distance := point.DistanceFromOrigin()

// ✅ GOOD: Export only what's necessary
type user struct {  // Lowercase = private
    id       int
    password string  // Keep private!
}

func (u user) GetID() int { return u.id }  // Controlled access

// ❌ BAD: Everything public (no encapsulation)
type User struct {
    ID       int
    Password string  // Should be private!
}

// ✅ GOOD: Zero value is useful
type Config struct {
    MaxRetries int  // 0 is sensible default
    Timeout    int  // 0 could mean no timeout
}

// Consider providing defaults via constructor
func NewConfig() *Config {
    return &Config{
        MaxRetries: 3,
        Timeout:    30,
    }
}
```

## Common Mistakes

```go
// Mistake 1: Forgetting pointer receiver when modifying
func (u User) Deactivate() {  // Value receiver!
    u.IsActive = false  // Modifies copy, not original!
}

// Mistake 2: Nil pointer dereference
var user *User
fmt.Println(user.Username)  // PANIC! user is nil

// Safe approach:
if user != nil {
    fmt.Println(user.Username)
}

// Mistake 3: Comparing structs with non-comparable fields
type User struct {
    Name    string
    Friends []string  // Slice is not comparable!
}

// This won't compile:
// if user1 == user2 { ... }

// Mistake 4: Large struct by value (expensive copying)
type HugeStruct struct {
    Data [1000000]int
}

func process(h HugeStruct) {  // Copies entire array!
    // ...
}

// Better:
func process(h *HugeStruct) {  // Just copies pointer
    // ...
}
```

## Key teaching points:
- Always validate in [constructors](https://go.dev/doc/effective_go#composite_literals)
- Be consistent with [receiver types](https://go.dev/doc/effective_go#pointers_vs_values)
- Use [named field initialization](https://go.dev/doc/effective_go#composite_literals)
- [Export only what's necessary](https://go.dev/ref/spec#Exported_identifiers)
- Design for [useful zero values](https://go.dev/ref/spec#The_zero_value)
