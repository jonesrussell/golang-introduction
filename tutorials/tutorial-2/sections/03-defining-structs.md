# Defining Structs

**Duration:** 5-6 minutes

## Topics to cover:
- Basic [struct definition](https://go.dev/ref/spec#Struct_types)
- [Field naming conventions](https://go.dev/doc/effective_go#names)
- [Field tags](https://pkg.go.dev/encoding/json#Marshal) (preview for JSON)
- [Nested structs](https://go.dev/ref/spec#Struct_types)
- [Anonymous fields](https://go.dev/ref/spec#Struct_types)

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
- [Type definition](https://go.dev/ref/spec#Type_declarations) creates a new type in your package
- Fields are accessed with [dot notation](https://go.dev/ref/spec#Selectors)
- [Field tags](https://pkg.go.dev/encoding/json#Marshal) are metadata (used by encoding/json, validation libraries)
- Nested structs model "has-a" relationships
- [Anonymous structs](https://go.dev/ref/spec#Struct_types) are useful for one-off data structures
- Convention: one struct per logical concept
