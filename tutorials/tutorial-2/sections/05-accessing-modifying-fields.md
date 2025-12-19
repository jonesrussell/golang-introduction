# Accessing and Modifying Struct Fields

**Duration:** 4-5 minutes

## Topics to cover:
- [Dot notation](https://go.dev/ref/spec#Selectors)
- [Automatic pointer dereferencing](https://go.dev/ref/spec#Selectors)
- Value vs pointer receivers (preview)
- [Comparing structs](https://go.dev/ref/spec#Comparison_operators)

## Code Examples

```go
type Book struct {
    Title  string
    Author string
    Pages  int
}

// Create a book
book := Book{
    Title:  "The Go Programming Language",
    Author: "Donovan & Kernighan",
    Pages:  380,
}

// Accessing fields
fmt.Println(book.Title)   // Read
book.Pages = 400          // Modify

// Working with pointers - automatic dereferencing
bookPtr := &Book{
    Title:  "Learning Go",
    Author: "Jon Bodner",
    Pages:  350,
}

// Go automatically dereferences, these are equivalent:
fmt.Println(bookPtr.Title)
fmt.Println((*bookPtr).Title)

// Comparing structs (all fields must be comparable)
book1 := Book{Title: "Go", Author: "Pike", Pages: 200}
book2 := Book{Title: "Go", Author: "Pike", Pages: 200}
book3 := Book{Title: "Rust", Author: "Klabnik", Pages: 500}

fmt.Println(book1 == book2)  // true (all fields match)
fmt.Println(book1 == book3)  // false

// Structs with slices/maps cannot be compared with ==
type Library struct {
    Name  string
    Books []Book  // Contains slice - not comparable!
}

// This would be a compilation error:
// lib1 := Library{Name: "City Library"}
// lib2 := Library{Name: "City Library"}
// fmt.Println(lib1 == lib2)  // ERROR!
```

## Key teaching points:
- [Dot notation](https://go.dev/ref/spec#Selectors) works the same for values and pointers
- Go's [automatic dereferencing](https://go.dev/ref/spec#Selectors) reduces pointer boilerplate
- [Struct comparison](https://go.dev/ref/spec#Comparison_operators) requires all fields to be comparable
- [Slices, maps, and functions](https://go.dev/ref/spec#Comparison_operators) make structs non-comparable
- Use [reflect.DeepEqual()](https://pkg.go.dev/reflect#DeepEqual) for complex comparisons
