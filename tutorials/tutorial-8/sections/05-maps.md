# Maps - Key-Value Storage

**Duration:** 8-10 minutes

## Topics to cover:
- [Map creation](https://go.dev/ref/spec#Map_types) and operations
- [Key requirements](https://go.dev/ref/spec#Map_types)
- [Iteration order](https://go.dev/ref/spec#For_statements)
- nil maps vs empty maps
- Concurrent access

## Code Examples

```go runnable
package main

import "fmt"

func main() {
    // Creating maps
    m1 := make(map[string]int)           // Empty map
    m2 := map[string]int{}               // Empty map (literal)
    m3 := map[string]int{                // With values
        "alice": 95,
        "bob":   87,
    }
    var m4 map[string]int                // nil map

    fmt.Println("m1:", m1)
    fmt.Println("m2:", m2)
    fmt.Println("m3:", m3)
    fmt.Printf("m4 is nil: %v\n", m4 == nil)

    // Basic operations
    m := make(map[string]int)

    m["alice"] = 95            // Set
    score := m["alice"]        // Get (returns zero value if missing)
    fmt.Println("alice score:", score)

    delete(m, "alice")         // Delete

    // Check if key exists
    m["bob"] = 87
    score, ok := m["bob"]
    if ok {
        fmt.Printf("bob's score: %d\n", score)
    }

    // Idiom: check and use
    if score, ok := m["bob"]; ok {
        fmt.Printf("Bob's score: %d\n", score)
    }

    // Length
    fmt.Println("map size:", len(m))
}
```

## nil map vs empty map

```go
var nilMap map[string]int     // nil
emptyMap := map[string]int{}  // empty but initialized

// Reading from nil map returns zero value
_ = nilMap["key"]  // OK, returns 0

// Writing to nil map panics!
// nilMap["key"] = 1  // PANIC!

// Always initialize before writing
if nilMap == nil {
    nilMap = make(map[string]int)
}
nilMap["key"] = 1
```

## Iteration

```go runnable
package main

import (
    "fmt"
    "sort"
)

func main() {
    // Iteration - ORDER IS RANDOM!
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    fmt.Println("Random order iteration:")
    for key, value := range m {
        fmt.Printf("  %s: %d\n", key, value)
    }
    // Output order varies each run!

    // Sorted iteration
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    fmt.Println("Sorted iteration:")
    for _, k := range keys {
        fmt.Printf("  %s: %d\n", k, m[k])
    }
}
```

## Valid Key Types

```go
// Valid key types (must be comparable)
// OK: int, string, float, bool, pointer, struct (if all fields comparable)
// NOT OK: slice, map, function

type Point struct{ X, Y int }
pointMap := make(map[Point]string)
pointMap[Point{1, 2}] = "origin"
```

## Maps with Struct Values

```go
type User struct {
    Name  string
    Email string
}

users := make(map[int]User)
users[1] = User{Name: "Alice", Email: "alice@example.com"}

// Can't modify struct field directly!
// users[1].Name = "Alicia"  // Compile error!

// Must replace entire value
u := users[1]
u.Name = "Alicia"
users[1] = u

// Or use pointer values
usersPtr := make(map[int]*User)
usersPtr[1] = &User{Name: "Alice"}
usersPtr[1].Name = "Alicia"  // OK!
```

## Key teaching points:
- Always initialize before writing
- Use comma-ok for existence check
- [Iteration order](https://go.dev/ref/spec#For_statements) is random
- Only [comparable types](https://go.dev/ref/spec#Comparison_operators) as keys
- Not safe for concurrent use
