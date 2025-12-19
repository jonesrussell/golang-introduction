# Struct Initialization

**Duration:** 7-8 minutes

## Topics to cover:
- [Zero value](https://go.dev/ref/spec#The_zero_value) initialization
- [Struct literals](https://go.dev/ref/spec#Composite_literals)
- Named field initialization
- Positional initialization (avoid)
- [Pointer to struct](https://go.dev/ref/spec#Address_operators) with `&`
- [Constructor functions](https://go.dev/doc/effective_go#composite_literals) (idiomatic pattern)

## Code Examples

```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// 1. Zero value - all fields get their type's zero value
var p1 Person
fmt.Printf("%+v\n", p1)  // {FirstName: LastName: Age:0}

// 2. Struct literal with named fields (RECOMMENDED)
p2 := Person{
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
}

// 3. Partial initialization (remaining fields get zero values)
p3 := Person{
    FirstName: "Jane",
    LastName:  "Smith",
    // Age will be 0
}

// 4. Positional initialization (AVOID - brittle to field reordering)
p4 := Person{"Bob", "Wilson", 25}  // Works but not recommended

// 5. Pointer to struct (common for large structs or when mutating)
p5 := &Person{
    FirstName: "Alice",
    LastName:  "Johnson",
    Age:       28,
}

// 6. Using new() - returns pointer with zero values
p6 := new(Person)
p6.FirstName = "Charlie"

// 7. Constructor function (IDIOMATIC PATTERN)
func NewPerson(firstName, lastName string, age int) *Person {
    return &Person{
        FirstName: firstName,
        LastName:  lastName,
        Age:       age,
    }
}

// Usage
p7 := NewPerson("David", "Brown", 35)
```

## Key teaching points:
- Named field initialization is most readable and maintainable
- Trailing comma is required on multi-line initialization
- [Zero values](https://go.dev/ref/spec#The_zero_value) make structs safe to use even when empty
- [Pointer initialization](https://go.dev/ref/spec#Address_operators) with `&` avoids copying large structs
- [Constructor functions](https://go.dev/doc/effective_go#composite_literals) allow validation and default values
- Convention: `New` prefix for constructors (NewUser, NewProduct, etc.)
