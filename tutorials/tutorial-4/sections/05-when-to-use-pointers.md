# When to Use Pointers

**Duration:** 6-7 minutes

## Topics to cover:
- Performance considerations
- Mutation requirements
- API design choices
- [Nil as meaningful value](https://go.dev/ref/spec#The_zero_value)

## Code Examples

```go
// USE POINTERS WHEN:

// 1. You need to modify the original value
func (c *Counter) Increment() {
    c.count++  // Modifies original
}

// 2. Struct is large (avoid copying)
type LargeStruct struct {
    Data [1000000]byte
}

func processLarge(ls *LargeStruct) {  // Good: copies 8 bytes (pointer)
    // ...
}

func processLargeBad(ls LargeStruct) {  // Bad: copies 1MB!
    // ...
}

// 3. Nil is a meaningful value (optional/not set)
type Config struct {
    Timeout  *int  // nil means "use default"
    MaxRetry *int  // nil means "use default"
}

func NewConfig() *Config {
    return &Config{}  // All optional fields nil
}

func (c *Config) GetTimeout() int {
    if c.Timeout == nil {
        return 30  // Default
    }
    return *c.Timeout
}

// 4. Consistency - if any method needs pointer, all should use pointer
type Database struct {
    conn *sql.DB
}

func (db *Database) Query(sql string) {}   // Pointer
func (db *Database) Execute(sql string) {} // Pointer (consistency)
func (db *Database) Close() {}             // Pointer (consistency)

// DON'T USE POINTERS WHEN:

// 1. Small, immutable values
type Point struct {
    X, Y float64
}

func Distance(p1, p2 Point) float64 {  // Values fine for small structs
    dx := p2.X - p1.X
    dy := p2.Y - p1.Y
    return math.Sqrt(dx*dx + dy*dy)
}

// 2. You want immutability guarantees
func (p Point) Move(dx, dy float64) Point {
    return Point{X: p.X + dx, Y: p.Y + dy}  // Returns new point
}

// 3. Maps, slices, channels - already reference types
func processSlice(data []int) {  // Slice header copied, but data shared
    data[0] = 999  // Modifies original!
}
```

## Decision Flowchart

```
Should I use a pointer?

Does the function need to modify the value?
    Yes → Use pointer
    No ↓

Is the struct large (>= 3-4 fields or contains large arrays)?
    Yes → Use pointer
    No → Use value (copy is fine)
```

## Key teaching points:
- Pointers for mutation
- Pointers for large structs (performance)
- Pointers when [nil is meaningful](https://go.dev/ref/spec#The_zero_value)
- Values for small, immutable data
- Be consistent within a type
- [Slices, maps, channels](https://go.dev/ref/spec#Slice_types) are already "reference-like"
