# Multiple Embedding

**Duration:** 6-7 minutes

## Topics to cover:
- Embedding [multiple structs](https://go.dev/ref/spec#Struct_types)
- Field/method [name conflicts](https://go.dev/ref/spec#Selectors)
- [Resolution order](https://go.dev/ref/spec#Selectors)
- Explicitly accessing embedded types

## Code Examples

```go
type Timestamps struct {
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (t Timestamps) Age() time.Duration {
    return time.Since(t.CreatedAt)
}

type Metadata struct {
    Tags        []string
    Description string
}

func (m Metadata) HasTag(tag string) bool {
    for _, t := range m.Tags {
        if t == tag {
            return true
        }
    }
    return false
}

// Embed multiple structs
type Article struct {
    Timestamps  // Embedded
    Metadata    // Embedded
    Title       string
    Content     string
    Author      string
}

// Usage
article := Article{
    Timestamps: Timestamps{
        CreatedAt: time.Now().Add(-24 * time.Hour),
        UpdatedAt: time.Now(),
    },
    Metadata: Metadata{
        Tags:        []string{"golang", "tutorial"},
        Description: "Learn about struct embedding",
    },
    Title:   "Go Structs",
    Content: "...",
    Author:  "Russell",
}

// Can access all promoted fields
fmt.Println(article.CreatedAt)   // From Timestamps
fmt.Println(article.Tags)        // From Metadata
fmt.Println(article.Title)       // From Article

// Can call all promoted methods
fmt.Println(article.Age())              // From Timestamps
fmt.Println(article.HasTag("golang"))   // From Metadata
```

## Handling Conflicts

```go
// Handling conflicts - same field/method name in multiple embedded structs
type A struct {
    Name string
}

func (a A) Display() {
    fmt.Println("From A:", a.Name)
}

type B struct {
    Name string
}

func (b B) Display() {
    fmt.Println("From B:", b.Name)
}

type C struct {
    A  // Both have Name field and Display method
    B
}

c := C{
    A: A{Name: "A's name"},
    B: B{Name: "B's name"},
}

// Ambiguous - won't compile:
// fmt.Println(c.Name)     // ERROR: ambiguous selector c.Name
// c.Display()             // ERROR: ambiguous selector c.Display

// Must be explicit when there's a conflict:
fmt.Println(c.A.Name)  // OK - explicit
fmt.Println(c.B.Name)  // OK - explicit
c.A.Display()          // OK - explicit
c.B.Display()          // OK - explicit

// No conflict if outer struct has the same field
type D struct {
    A
    B
    Name string  // Shadows embedded Name fields
}

d := D{
    A:    A{Name: "A's name"},
    B:    B{Name: "B's name"},
    Name: "D's name",
}

fmt.Println(d.Name)    // D's name (outer struct wins)
fmt.Println(d.A.Name)  // A's name (still accessible)
fmt.Println(d.B.Name)  // B's name (still accessible)
```

## Key teaching points:
- Can embed [multiple structs](https://go.dev/ref/spec#Struct_types)
- All fields/methods from all embedded types are [promoted](https://go.dev/ref/spec#Selectors)
- [Name conflicts](https://go.dev/ref/spec#Selectors) cause compile errors (not runtime!)
- Must explicitly specify which embedded type when [ambiguous](https://go.dev/ref/spec#Selectors)
- Outer struct fields [shadow](https://go.dev/ref/spec#Selectors) embedded fields with same name
- This is safer than inheritance (explicit is better than implicit)
