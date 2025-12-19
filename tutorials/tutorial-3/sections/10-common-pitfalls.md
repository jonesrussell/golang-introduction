# Common Pitfalls & Best Practices

**Duration:** 5-6 minutes

## Common Mistakes

```go
// ❌ PITFALL 1: Embedding to avoid typing (wrong motivation)
type User struct {
    Address  // BAD - unclear semantic relationship
    Name string
}

// ✅ BETTER: Explicit composition for clarity
type User struct {
    Name    string
    Address Address  // Clear: User has an Address
}

// ❌ PITFALL 2: Over-embedding creates confusion
type Admin struct {
    User
    Manager
    DatabaseConnection  // Too many embedded types!
    Logger
}
// Hard to understand what Admin IS

// ✅ BETTER: Selective embedding + composition
type Admin struct {
    User                 // Admin IS a specialized User
    ManagedTeam   []User // Composition for clarity
    db            *sql.DB  // Private, not promoted
    logger        Logger   // Private, not promoted
}

// ❌ PITFALL 3: Circular embedding
type A struct {
    B  // A embeds B
}

type B struct {
    A  // B embeds A - CIRCULAR!
}
// Won't compile - infinite size

// ❌ PITFALL 4: Embedding with name conflicts (unintentional)
type Author struct {
    Name string
}

type Book struct {
    Author  // Embeds Author
    Name string  // Conflicts! Which Name?
}

book := Book{
    Author: Author{Name: "Tolkien"},
    Name:   "The Hobbit",
}

// book.Name is the Book's Name (outer wins)
// book.Author.Name is the Author's Name

// ✅ BETTER: Use different field names or explicit composition
type Book struct {
    Title  string
    Author Author  // Explicit - no confusion
}

// ❌ PITFALL 5: Exposing internals through embedding
type Cache struct {
    sync.Mutex  // BAD - exposes Lock/Unlock publicly
    data map[string]interface{}
}

// Now users can do: cache.Lock() - should be internal!

// ✅ BETTER: Private field for sync
type Cache struct {
    mu   sync.Mutex  // Private
    data map[string]interface{}
}

func (c *Cache) Get(key string) interface{} {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.data[key]
}

// ✅ EXCEPTION: sync.Mutex embedding is OK for types that ARE synchronization primitives
type SafeCounter struct {
    sync.Mutex  // OK - SafeCounter IS synchronized
    count int
}
```

## Best Practices Summary

```go
// ✅ DO: Use embedding for mixins
type Auditable struct { /* ... */ }
type Product struct {
    Auditable  // Adds audit functionality
    // ...
}

// ✅ DO: Use embedding for base functionality
type BaseEntity struct { /* ... */ }
type User struct {
    BaseEntity
    // ...
}

// ✅ DO: Use composition for "has-a" relationships
type Car struct {
    Engine Engine  // Car HAS an Engine
}

// ✅ DO: Keep embedding shallow (1-2 levels max)
type Admin struct {
    User  // OK - one level
}

// ✅ DO: Document why you're embedding
type Product struct {
    Auditable  // Embedding for audit mixin pattern
    // ...
}

// ❌ DON'T: Embed just to save typing
// ❌ DON'T: Embed more than 3-4 types in one struct
// ❌ DON'T: Embed when composition is clearer
// ❌ DON'T: Create deep embedding hierarchies (3+ levels)
```

## Key teaching points:
- Be intentional about why you're [embedding](https://go.dev/ref/spec#Struct_types)
- [Composition](https://go.dev/ref/spec#Struct_types) is often clearer than embedding
- Watch for [name conflicts](https://go.dev/ref/spec#Selectors)
- Don't expose internal synchronization
- Keep embedding shallow
