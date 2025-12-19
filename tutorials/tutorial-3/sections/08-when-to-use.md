# When to Use Embedding vs Composition

**Duration:** 5-6 minutes

## Decision Guide

```go
// Use EMBEDDING when:
// 1. You want promoted fields/methods (convenience)
// 2. You're implementing mixin behavior
// 3. The embedded type is a "base" or "common" functionality
// 4. You want the outer type to satisfy interfaces of embedded type

type User struct {
    Auditable  // ✅ Good - mixin behavior
    ID         int
    Name       string
}

// Use EXPLICIT COMPOSITION when:
// 1. The relationship is clearly "has-a"
// 2. You want explicit access (clarity over convenience)
// 3. You might have multiple of the same type
// 4. The field has semantic meaning

type Car struct {
    Engine Engine  // ✅ Good - Car HAS-AN Engine
    Brand  string
}

type House struct {
    PrimaryBathroom Bathroom  // ✅ Good - semantic meaning
    Bedrooms        []Bedroom // Multiple bedrooms
}
```

## Anti-Patterns to Avoid

```go
// ❌ BAD: Embedding just to avoid typing
type Person struct {
    Address  // BAD - this should be explicit field
    Name string
}
// Problem: person.Street is confusing - is it Person's street?
// Better: person.Address.Street - clear relationship

// ❌ BAD: Embedding unrelated types
type Config struct {
    sync.Mutex  // BAD - Config is not a specialized Mutex
    Settings map[string]string
}
// Problem: Exposes Lock/Unlock at wrong level
// Better: Have a mutex as private field

// ✅ GOOD: Embedding mutex in correct context
type SafeCounter struct {
    sync.Mutex  // OK - SafeCounter IS synchronized
    count int
}

func (sc *SafeCounter) Increment() {
    sc.Lock()
    defer sc.Unlock()
    sc.count++
}

// ❌ BAD: Embedding for code reuse when composition is clearer
type AdminUser struct {
    User  // Confusing - is AdminUser a specialized User?
    AdminLevel int
}
// Better: explicit composition or separate types

// ✅ GOOD: Clear "is-a-specialized-version" relationship
type Buffer struct {
    bytes.Buffer  // OK - adding functionality to Buffer
}

func (b *Buffer) WriteJSON(v interface{}) error {
    data, err := json.Marshal(v)
    if err != nil {
        return err
    }
    b.Write(data)
    return nil
}
```

## Decision Flowchart

```
Does the outer type need ALL behaviors of the inner type?
    |
    No ─── Use explicit composition
    |
    Yes
    |
    Is it a "base" or "mixin" pattern?
        |
        Yes ─── Use embedding
        |
        No ─── Consider explicit composition
```

## Key teaching points:
- Embedding is powerful but can be overused
- Prefer explicit composition for clarity
- Use embedding for mixins and base patterns
- Don't embed just to save typing
- Think about the conceptual relationship
- "Has-a" → composition; "Is-a-kind-of" → maybe embedding
