# Advanced Preview & Next Steps

**Duration:** 3-4 minutes

## Preview of Upcoming Topics

```go
// Struct embedding (next video)
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person        // Embedded struct
    EmployeeID int
    Department string
}

// JSON encoding/decoding
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

data, _ := json.Marshal(user)

// Interfaces with structs
type Stringer interface {
    String() string
}

// Any type with String() method satisfies Stringer
```

## Recap What Was Covered:
- [Defining structs](https://go.dev/ref/spec#Struct_types)
- [Initialization patterns](https://go.dev/ref/spec#Composite_literals) (literals, constructors)
- [Value vs pointer receivers](https://go.dev/doc/effective_go#pointers_vs_values)
- [Methods](https://go.dev/ref/spec#Method_declarations) and when to use each
- Best practices and common mistakes

## Practice Suggestions:
1. **Easy:** Create a Book struct with methods for GetFullTitle, IsLongBook
2. **Medium:** Build a TodoList struct that manages Todo items
3. **Challenge:** Create a BankAccount system with Account and Transaction structs

## Resources:
- [Effective Go](https://go.dev/doc/effective_go): go.dev/doc/effective_go
- [Go by Example](https://gobyexample.com/structs): gobyexample.com/structs
- [Go Tour - Structs](https://go.dev/tour/moretypes/2): go.dev/tour/moretypes/2

## Key teaching points:
- [Structs](https://go.dev/ref/spec#Struct_types) are fundamental to Go programming
- Master [value vs pointer semantics](https://go.dev/doc/effective_go#pointers_vs_values)
- [Constructor pattern](https://go.dev/doc/effective_go#composite_literals) ensures valid state
- [Composition](https://go.dev/doc/faq#inheritance) enables code reuse
