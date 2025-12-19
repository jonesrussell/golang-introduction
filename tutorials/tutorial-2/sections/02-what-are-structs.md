# What Are Structs?

**Duration:** 3-4 minutes

## Topics to cover:
- Structs as custom types
- Grouping related data together
- Type definition syntax
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
- Structs create custom types
- Fields (properties) are declared with name and type
- Exported fields start with capital letters (public)
- Lowercase fields are package-private
- Structs are value types (copied by default)
