# What Are Structs?

**Duration:** 3-4 minutes

## Topics to cover:
- [Structs](https://go.dev/ref/spec#Struct_types) as custom types
- Grouping related data together
- [Type definition](https://go.dev/ref/spec#Type_declarations) syntax
- When to use structs vs primitives

## Code Example

```go
// Without structs - messy and error-prone
firstName := "Jane"
lastName := "Smith"
email := "jane@example.com"
age := 28

// With structs - organized and type-safe
type User struct {
    FirstName string
    LastName  string
    Email     string
    Age       int
}
```

## Key teaching points:
- Structs create [custom types](https://go.dev/ref/spec#Type_declarations)
- Fields (properties) are declared with name and type
- [Exported fields](https://go.dev/ref/spec#Exported_identifiers) start with capital letters (public)
- Lowercase fields are [package-private](https://go.dev/ref/spec#Exported_identifiers)
- Structs are [value types](https://go.dev/ref/spec#Types) (copied by default)
