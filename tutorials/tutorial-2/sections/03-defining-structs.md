# Defining Structs

**Duration:** 5-6 minutes

## Topics to cover:
- Basic struct definition
- Field naming conventions
- Field tags (preview for JSON)
- Nested structs
- Anonymous fields

## Code Examples

```go
// Basic struct definition
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// Struct with multiple field types
type Product struct {
    ID          int
    Name        string
    Price       float64
    InStock     bool
    Tags        []string
}

// Struct with field tags (for JSON, validation, etc.)
type User struct {
    ID        int    `json:"id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    CreatedAt string `json:"created_at"`
}

// Nested structs
type Address struct {
    Street  string
    City    string
    ZipCode string
}

type Employee struct {
    Name    string
    Email   string
    Address Address  // Nested struct
}

// Anonymous struct (inline, one-off use)
config := struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080,
}
```

## Key teaching points:
- Type definition creates a new type in your package
- Fields are accessed with dot notation
- Field tags are metadata (used by encoding/json, validation libraries)
- Nested structs model "has-a" relationships
- Anonymous structs are useful for one-off data structures
- Convention: one struct per logical concept
