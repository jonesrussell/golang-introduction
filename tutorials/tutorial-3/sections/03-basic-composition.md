# Basic Struct Composition

**Duration:** 5-6 minutes

## Topics to cover:
- Composition via [explicit fields](https://go.dev/ref/spec#Struct_types)
- Has-a relationships
- [Accessing nested fields](https://go.dev/ref/spec#Selectors)
- When explicit composition is appropriate

## Code Examples

```go
// Basic composition - one struct contains another

type Address struct {
    Street  string
    City    string
    State   string
    ZipCode string
}

type Person struct {
    FirstName string
    LastName  string
    Address   Address  // Explicit field - composition
}

// Usage - explicit field access
person := Person{
    FirstName: "John",
    LastName:  "Doe",
    Address: Address{
        Street:  "123 Main St",
        City:    "Springfield",
        State:   "IL",
        ZipCode: "62701",
    },
}

// Accessing nested fields
fmt.Println(person.Address.Street)  // Must go through Address field
fmt.Println(person.Address.City)

// Methods on composed struct
func (a Address) FullAddress() string {
    return fmt.Sprintf("%s, %s, %s %s", 
        a.Street, a.City, a.State, a.ZipCode)
}

// Must access through field name
fmt.Println(person.Address.FullAddress())

// Another example - explicit composition
type Engine struct {
    Horsepower int
    Type       string
}

type Car struct {
    Brand  string
    Model  string
    Engine Engine  // Car HAS-AN Engine
}

car := Car{
    Brand: "Toyota",
    Model: "Camry",
    Engine: Engine{
        Horsepower: 203,
        Type:       "V6",
    },
}

fmt.Printf("%s %s has %d HP\n", 
    car.Brand, car.Model, car.Engine.Horsepower)
```

## Key teaching points:
- Regular composition uses [named fields](https://go.dev/ref/spec#Struct_types)
- Represents clear "has-a" relationships
- Must explicitly reference the [field name](https://go.dev/ref/spec#Selectors)
- Good when the relationship is explicit (Car has Engine)
- Fields and methods are accessed through the [field name](https://go.dev/ref/spec#Selectors)
