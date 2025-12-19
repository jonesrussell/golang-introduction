# When to Use Structs vs Other Types

**Duration:** 3-4 minutes

## Decision Guide

```go
// Use struct when: Grouping related data
type Address struct {
    Street, City, State, Zip string
}

// Use map when: Dynamic keys, not fixed structure
userScores := map[string]int{
    "alice": 95,
    "bob":   87,
}

// Use struct when: Fixed fields with different types
type Event struct {
    Name      string
    Timestamp time.Time
    Attendees []string
    Metadata  map[string]interface{}
}

// Use interface when: Behavior matters more than data
type Writer interface {
    Write([]byte) (int, error)
}

// Use struct when: You need methods with state
type Database struct {
    conn *sql.DB
}

func (db *Database) Query(sql string) (*Rows, error) {
    // Uses db.conn
}
```

## Guidelines:
- **Struct:** Fixed schema, type safety, [methods](https://go.dev/ref/spec#Method_declarations)
- **Map:** Dynamic keys, uniform value types
- **Interface:** Define contracts, [polymorphism](https://go.dev/doc/faq#polymorphism)
- **Slice:** Ordered collection of same type

## Key teaching points:
- Choose based on whether structure is fixed or dynamic
- [Structs](https://go.dev/ref/spec#Struct_types) provide compile-time type safety
- [Maps](https://go.dev/ref/spec#Map_types) provide runtime flexibility
- [Interfaces](https://go.dev/ref/spec#Interface_types) define behavior contracts
