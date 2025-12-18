## **Video Tutorial Plan: Go Slices and Maps**

### **Video Metadata**
- **Title:** Go Slices and Maps: Internals and Best Practices
- **Duration Target:** 35-45 minutes
- **Difficulty:** Intermediate
- **Prerequisites:** Go Basics, Pointers

---

## **Video Structure**

### **1. Introduction (2-3 min)**
- Welcome and what viewers will learn
- Why understanding internals matters
- Arrays vs slices vs maps
- Preview: Building an in-memory database

---

### **2. Arrays - The Foundation (4-5 min)**

**Topics to cover:**
- Fixed size, value type
- Declaration and initialization
- When to use arrays (rarely!)

**Code Examples:**
```go
// Array declaration - size is part of the type
var arr1 [5]int                    // Zero values
arr2 := [5]int{1, 2, 3, 4, 5}      // With values
arr3 := [...]int{1, 2, 3}          // Size inferred (3)
arr4 := [5]int{0: 1, 4: 5}         // Sparse: [1, 0, 0, 0, 5]

// Arrays are values - copied on assignment
a := [3]int{1, 2, 3}
b := a           // Full copy!
b[0] = 100
fmt.Println(a)   // [1 2 3] - unchanged
fmt.Println(b)   // [100 2 3]

// Arrays passed by value to functions (copied!)
func modifyArray(arr [3]int) {
    arr[0] = 999  // Modifies copy
}

// Different sizes = different types!
var x [3]int
var y [4]int
// x = y  // Compile error: different types!

// When arrays make sense:
// - Known fixed size at compile time
// - Part of struct (embedded)
// - Performance-critical with small, fixed data
type Point struct {
    Coords [3]float64  // x, y, z
}
```

**Key teaching points:**
- Arrays have fixed size (part of type)
- Passed by value (copied)
- Different sizes = different types
- Use slices instead in most cases

---

### **3. Slices - The Workhorse (10-12 min)**

**Topics to cover:**
- Slice header (pointer, length, capacity)
- Creating slices
- Slicing operations
- Append and growth
- Copy

**Code Examples:**
```go
// Slice is a descriptor: (pointer, length, capacity)
// type slice struct {
//     array unsafe.Pointer
//     len   int
//     cap   int
// }

// Creating slices
s1 := []int{1, 2, 3}                // Literal
s2 := make([]int, 5)                // Length 5, capacity 5
s3 := make([]int, 3, 10)            // Length 3, capacity 10
var s4 []int                        // nil slice (len=0, cap=0)

// Slice from array
arr := [5]int{1, 2, 3, 4, 5}
s5 := arr[1:4]    // [2, 3, 4] - shares backing array!

// Length vs Capacity
s := make([]int, 3, 5)
fmt.Println(len(s))  // 3 - number of elements
fmt.Println(cap(s))  // 5 - underlying array size

// Accessing elements
s := []int{10, 20, 30}
fmt.Println(s[0])  // 10
s[1] = 25          // Modify
// s[5] = 50       // Panic: index out of range

// Slicing syntax [low:high:max]
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
s1 := s[2:5]      // [2, 3, 4] len=3, cap=8
s2 := s[2:5:5]    // [2, 3, 4] len=3, cap=3 (limited capacity)
s3 := s[:5]       // [0, 1, 2, 3, 4]
s4 := s[5:]       // [5, 6, 7, 8, 9]
s5 := s[:]        // Full slice (copy of header, same array)

// IMPORTANT: Slices share backing array!
original := []int{1, 2, 3, 4, 5}
slice := original[1:4]
slice[0] = 999
fmt.Println(original)  // [1, 999, 3, 4, 5] - modified!

// Append - may create new backing array
s := []int{1, 2, 3}
fmt.Printf("Before: len=%d, cap=%d, ptr=%p\n", len(s), cap(s), s)

s = append(s, 4)  // Might reuse or create new array
fmt.Printf("After:  len=%d, cap=%d, ptr=%p\n", len(s), cap(s), s)

s = append(s, 5, 6, 7, 8, 9, 10)  // Definitely new array
fmt.Printf("Growth: len=%d, cap=%d, ptr=%p\n", len(s), cap(s), s)

// Append to nil slice works!
var s []int
s = append(s, 1, 2, 3)  // Creates backing array

// Append another slice
s1 := []int{1, 2, 3}
s2 := []int{4, 5, 6}
s3 := append(s1, s2...)  // ... unpacks slice

// Copy - explicit copy of elements
src := []int{1, 2, 3, 4, 5}
dst := make([]int, 3)
n := copy(dst, src)  // Copies min(len(dst), len(src))
fmt.Println(n, dst)  // 3, [1 2 3]

// Safe full copy
original := []int{1, 2, 3, 4, 5}
copied := make([]int, len(original))
copy(copied, original)
// Now modifications to copied don't affect original
```

**Slice growth pattern:**
```go
// Growth algorithm (approximately):
// - cap < 256: double
// - cap >= 256: grow by ~25%

func demonstrateGrowth() {
    var s []int
    prevCap := 0

    for i := 0; i < 20; i++ {
        s = append(s, i)
        if cap(s) != prevCap {
            fmt.Printf("len=%2d, cap=%2d\n", len(s), cap(s))
            prevCap = cap(s)
        }
    }
}
// Output shows: 1, 2, 4, 8, 16, 32...
```

**Key teaching points:**
- Slice is header pointing to array
- Slicing creates new header, same array
- Append may reallocate (assign result!)
- Pre-allocate with make for known sizes
- Copy for independent slice

---

### **4. Common Slice Patterns (5-6 min)**

**Topics to cover:**
- Filtering
- Removing elements
- Insert at position
- Stack/queue operations

**Code Examples:**
```go
// Filter in-place (no allocation)
func filter(s []int, keep func(int) bool) []int {
    n := 0
    for _, v := range s {
        if keep(v) {
            s[n] = v
            n++
        }
    }
    return s[:n]
}

// Filter with new slice (preserves original)
func filterCopy(s []int, keep func(int) bool) []int {
    result := make([]int, 0, len(s))
    for _, v := range s {
        if keep(v) {
            result = append(result, v)
        }
    }
    return result
}

// Remove element at index
func remove(s []int, i int) []int {
    return append(s[:i], s[i+1:]...)
}

// Remove without preserving order (faster)
func removeUnordered(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

// Insert at index
func insert(s []int, i int, v int) []int {
    s = append(s, 0)           // Grow by 1
    copy(s[i+1:], s[i:])       // Shift right
    s[i] = v
    return s
}

// Stack operations
type Stack []int

func (s *Stack) Push(v int) {
    *s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
    if len(*s) == 0 {
        return 0, false
    }
    i := len(*s) - 1
    v := (*s)[i]
    *s = (*s)[:i]
    return v, true
}

// Queue operations
type Queue []int

func (q *Queue) Enqueue(v int) {
    *q = append(*q, v)
}

func (q *Queue) Dequeue() (int, bool) {
    if len(*q) == 0 {
        return 0, false
    }
    v := (*q)[0]
    *q = (*q)[1:]
    return v, true
}

// Reverse slice
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// Deduplicate (sorted slice)
func dedupe(s []int) []int {
    if len(s) < 2 {
        return s
    }
    j := 1
    for i := 1; i < len(s); i++ {
        if s[i] != s[i-1] {
            s[j] = s[i]
            j++
        }
    }
    return s[:j]
}
```

---

### **5. Maps - Key-Value Storage (8-10 min)**

**Topics to cover:**
- Map creation and operations
- Key requirements
- Iteration order
- nil maps vs empty maps
- Concurrent access

**Code Examples:**
```go
// Creating maps
m1 := make(map[string]int)           // Empty map
m2 := map[string]int{}               // Empty map (literal)
m3 := map[string]int{                // With values
    "alice": 95,
    "bob":   87,
}
var m4 map[string]int                // nil map

// Basic operations
m := make(map[string]int)

m["alice"] = 95            // Set
score := m["alice"]        // Get (returns zero value if missing)
delete(m, "alice")         // Delete

// Check if key exists
score, ok := m["bob"]
if !ok {
    fmt.Println("bob not found")
}

// Idiom: check and use
if score, ok := m["alice"]; ok {
    fmt.Printf("Alice's score: %d\n", score)
}

// Length
fmt.Println(len(m))  // Number of key-value pairs

// nil map vs empty map
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

// Iteration - ORDER IS RANDOM!
m := map[string]int{"a": 1, "b": 2, "c": 3}
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}
// Output order varies each run!

// Sorted iteration
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    fmt.Printf("%s: %d\n", k, m[k])
}

// Valid key types (must be comparable)
// OK: int, string, float, bool, pointer, struct (if all fields comparable)
// NOT OK: slice, map, function

type Point struct{ X, Y int }
pointMap := make(map[Point]string)
pointMap[Point{1, 2}] = "origin"

// Maps with struct values
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

// Maps are NOT safe for concurrent access
// Use sync.Map or protect with mutex
```

**Key teaching points:**
- Always initialize before writing
- Use comma-ok for existence check
- Iteration order is random
- Only comparable types as keys
- Not safe for concurrent use

---

### **6. Advanced Map Patterns (5-6 min)**

**Topics to cover:**
- Sets using maps
- Counting/grouping
- Cache patterns
- Map of slices/maps

**Code Examples:**
```go
// Set implementation
type Set map[string]struct{}

func NewSet() Set {
    return make(Set)
}

func (s Set) Add(item string) {
    s[item] = struct{}{}
}

func (s Set) Remove(item string) {
    delete(s, item)
}

func (s Set) Contains(item string) bool {
    _, ok := s[item]
    return ok
}

func (s Set) Size() int {
    return len(s)
}

// Usage
set := NewSet()
set.Add("apple")
set.Add("banana")
fmt.Println(set.Contains("apple"))  // true

// Counting / Frequency
words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
count := make(map[string]int)
for _, word := range words {
    count[word]++  // Zero value (0) works here!
}
// count = {"apple": 3, "banana": 2, "cherry": 1}

// Grouping
type Person struct {
    Name    string
    Country string
}

people := []Person{
    {"Alice", "USA"},
    {"Bob", "UK"},
    {"Charlie", "USA"},
    {"Diana", "UK"},
}

byCountry := make(map[string][]Person)
for _, p := range people {
    byCountry[p.Country] = append(byCountry[p.Country], p)
}
// byCountry["USA"] = [{Alice USA}, {Charlie USA}]

// Simple cache
type Cache struct {
    mu   sync.RWMutex
    data map[string]interface{}
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]interface{}),
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

// Nested maps
graph := make(map[string]map[string]int)

// Must initialize inner map!
if graph["A"] == nil {
    graph["A"] = make(map[string]int)
}
graph["A"]["B"] = 5

// Helper function
func setEdge(g map[string]map[string]int, from, to string, weight int) {
    if g[from] == nil {
        g[from] = make(map[string]int)
    }
    g[from][to] = weight
}

// Default value pattern
func getWithDefault(m map[string]int, key string, defaultVal int) int {
    if val, ok := m[key]; ok {
        return val
    }
    return defaultVal
}
```

---

### **7. Performance Considerations (4-5 min)**

**Topics to cover:**
- Pre-allocation
- Memory layout
- Clearing collections
- When to use arrays

**Code Examples:**
```go
// Pre-allocate slices when size is known
func badAppend() []int {
    var s []int
    for i := 0; i < 10000; i++ {
        s = append(s, i)  // Many reallocations
    }
    return s
}

func goodAppend() []int {
    s := make([]int, 0, 10000)  // Pre-allocate
    for i := 0; i < 10000; i++ {
        s = append(s, i)  // No reallocations
    }
    return s
}

// Pre-allocate maps
m := make(map[string]int, 10000)  // Hint: ~10000 entries

// Clearing a slice (reuse backing array)
s := []int{1, 2, 3, 4, 5}
s = s[:0]  // Length 0, capacity preserved

// Clearing a map (create new)
m := map[string]int{"a": 1, "b": 2}
// Option 1: Create new map
m = make(map[string]int)
// Option 2: Delete all (Go 1.21+ has clear())
for k := range m {
    delete(m, k)
}

// Slice of structs vs slice of pointers
// Structs: better cache locality, fewer allocations
type DataValue struct {
    ID   int
    Name string
}
valSlice := make([]DataValue, 1000)  // Contiguous memory

// Pointers: easier modification, but scattered memory
ptrSlice := make([]*DataValue, 1000)  // Pointers scattered

// Avoid slice header copies in hot loops
func process(data []int) {
    // Each iteration copies slice header (small but measurable)
    for i := range data {
        data[i] *= 2
    }
}

// String building
// BAD: String concatenation in loop
func badConcat(items []string) string {
    result := ""
    for _, item := range items {
        result += item + ","  // Creates new string each time!
    }
    return result
}

// GOOD: Use strings.Builder
func goodConcat(items []string) string {
    var sb strings.Builder
    for i, item := range items {
        if i > 0 {
            sb.WriteString(",")
        }
        sb.WriteString(item)
    }
    return sb.String()
}
```

---

### **8. Practical Example: In-Memory Database (8-10 min)**

**Build together:** A simple in-memory database

```go
package main

import (
    "errors"
    "fmt"
    "sort"
    "strings"
    "sync"
)

// Errors
var (
    ErrNotFound     = errors.New("record not found")
    ErrDuplicateID  = errors.New("duplicate ID")
    ErrInvalidQuery = errors.New("invalid query")
)

// Record represents a database record
type Record struct {
    ID     int
    Name   string
    Email  string
    Tags   []string
    Active bool
}

// Database is an in-memory database
type Database struct {
    mu       sync.RWMutex
    records  map[int]*Record      // Primary storage
    byEmail  map[string]*Record   // Email index
    byTag    map[string][]*Record // Tag index
    nextID   int
}

// NewDatabase creates a new database
func NewDatabase() *Database {
    return &Database{
        records: make(map[int]*Record),
        byEmail: make(map[string]*Record),
        byTag:   make(map[string][]*Record),
        nextID:  1,
    }
}

// Insert adds a new record
func (db *Database) Insert(name, email string, tags []string) (*Record, error) {
    db.mu.Lock()
    defer db.mu.Unlock()

    // Check for duplicate email
    if _, exists := db.byEmail[email]; exists {
        return nil, fmt.Errorf("email %s: %w", email, ErrDuplicateID)
    }

    // Create record
    record := &Record{
        ID:     db.nextID,
        Name:   name,
        Email:  email,
        Tags:   make([]string, len(tags)),
        Active: true,
    }
    copy(record.Tags, tags)
    db.nextID++

    // Store in primary map
    db.records[record.ID] = record

    // Update indexes
    db.byEmail[email] = record
    for _, tag := range tags {
        db.byTag[tag] = append(db.byTag[tag], record)
    }

    return record, nil
}

// Get retrieves a record by ID
func (db *Database) Get(id int) (*Record, error) {
    db.mu.RLock()
    defer db.mu.RUnlock()

    record, ok := db.records[id]
    if !ok {
        return nil, ErrNotFound
    }
    return record, nil
}

// GetByEmail retrieves a record by email
func (db *Database) GetByEmail(email string) (*Record, error) {
    db.mu.RLock()
    defer db.mu.RUnlock()

    record, ok := db.byEmail[email]
    if !ok {
        return nil, ErrNotFound
    }
    return record, nil
}

// FindByTag finds all records with a given tag
func (db *Database) FindByTag(tag string) []*Record {
    db.mu.RLock()
    defer db.mu.RUnlock()

    records := db.byTag[tag]
    // Return copy to prevent modification
    result := make([]*Record, len(records))
    copy(result, records)
    return result
}

// Update modifies a record
func (db *Database) Update(id int, name, email string) error {
    db.mu.Lock()
    defer db.mu.Unlock()

    record, ok := db.records[id]
    if !ok {
        return ErrNotFound
    }

    // Update email index if changed
    if email != record.Email {
        if _, exists := db.byEmail[email]; exists {
            return ErrDuplicateID
        }
        delete(db.byEmail, record.Email)
        db.byEmail[email] = record
    }

    record.Name = name
    record.Email = email
    return nil
}

// Delete removes a record
func (db *Database) Delete(id int) error {
    db.mu.Lock()
    defer db.mu.Unlock()

    record, ok := db.records[id]
    if !ok {
        return ErrNotFound
    }

    // Remove from indexes
    delete(db.byEmail, record.Email)
    for _, tag := range record.Tags {
        db.removeFromTagIndex(tag, record)
    }

    // Remove from primary storage
    delete(db.records, id)
    return nil
}

func (db *Database) removeFromTagIndex(tag string, record *Record) {
    records := db.byTag[tag]
    for i, r := range records {
        if r.ID == record.ID {
            // Remove by swapping with last and truncating
            records[i] = records[len(records)-1]
            db.byTag[tag] = records[:len(records)-1]
            break
        }
    }
    // Clean up empty slices
    if len(db.byTag[tag]) == 0 {
        delete(db.byTag, tag)
    }
}

// Query filters records based on criteria
type QueryFunc func(*Record) bool

func (db *Database) Query(filter QueryFunc) []*Record {
    db.mu.RLock()
    defer db.mu.RUnlock()

    result := make([]*Record, 0)
    for _, record := range db.records {
        if filter(record) {
            result = append(result, record)
        }
    }
    return result
}

// Count returns total number of records
func (db *Database) Count() int {
    db.mu.RLock()
    defer db.mu.RUnlock()
    return len(db.records)
}

// All returns all records sorted by ID
func (db *Database) All() []*Record {
    db.mu.RLock()
    defer db.mu.RUnlock()

    ids := make([]int, 0, len(db.records))
    for id := range db.records {
        ids = append(ids, id)
    }
    sort.Ints(ids)

    result := make([]*Record, len(ids))
    for i, id := range ids {
        result[i] = db.records[id]
    }
    return result
}

// Stats returns database statistics
func (db *Database) Stats() map[string]int {
    db.mu.RLock()
    defer db.mu.RUnlock()

    stats := map[string]int{
        "total_records": len(db.records),
        "unique_emails": len(db.byEmail),
        "unique_tags":   len(db.byTag),
    }

    active := 0
    for _, r := range db.records {
        if r.Active {
            active++
        }
    }
    stats["active_records"] = active

    return stats
}

func main() {
    fmt.Println("=== In-Memory Database Demo ===\n")

    db := NewDatabase()

    // Insert records
    r1, _ := db.Insert("Alice", "alice@example.com", []string{"admin", "developer"})
    r2, _ := db.Insert("Bob", "bob@example.com", []string{"developer"})
    r3, _ := db.Insert("Charlie", "charlie@example.com", []string{"designer", "developer"})
    _, _ = db.Insert("Diana", "diana@example.com", []string{"manager"})

    fmt.Printf("Inserted %d records\n\n", db.Count())

    // Get by ID
    record, err := db.Get(r1.ID)
    if err == nil {
        fmt.Printf("Get by ID %d: %s (%s)\n", record.ID, record.Name, record.Email)
    }

    // Get by email
    record, err = db.GetByEmail("bob@example.com")
    if err == nil {
        fmt.Printf("Get by email: %s (ID: %d)\n", record.Name, record.ID)
    }

    // Find by tag
    fmt.Println("\nDevelopers:")
    developers := db.FindByTag("developer")
    for _, r := range developers {
        fmt.Printf("  - %s\n", r.Name)
    }

    // Query with filter
    fmt.Println("\nNames starting with 'C':")
    cNames := db.Query(func(r *Record) bool {
        return strings.HasPrefix(r.Name, "C")
    })
    for _, r := range cNames {
        fmt.Printf("  - %s\n", r.Name)
    }

    // Update
    db.Update(r2.ID, "Robert", "robert@example.com")
    updated, _ := db.Get(r2.ID)
    fmt.Printf("\nUpdated: %s -> %s\n", "Bob", updated.Name)

    // Delete
    db.Delete(r3.ID)
    fmt.Printf("Deleted record %d, remaining: %d\n", r3.ID, db.Count())

    // All records
    fmt.Println("\nAll records:")
    for _, r := range db.All() {
        fmt.Printf("  [%d] %s <%s> %v\n", r.ID, r.Name, r.Email, r.Tags)
    }

    // Stats
    fmt.Println("\nDatabase stats:")
    for k, v := range db.Stats() {
        fmt.Printf("  %s: %d\n", k, v)
    }
}
```

**Walk through:**
- Primary storage with map
- Secondary indexes for fast lookups
- Thread-safe with RWMutex
- Query function for flexible filtering
- Proper cleanup when deleting
- Stats aggregation

---

### **9. Common Pitfalls (4-5 min)**

```go
// PITFALL 1: Modifying slice during iteration
s := []int{1, 2, 3, 4, 5}
for i, v := range s {
    if v%2 == 0 {
        s = append(s[:i], s[i+1:]...)  // BAD: modifies slice!
    }
}
// Fix: iterate backwards or use filter function

// PITFALL 2: Append doesn't always create new array
s1 := []int{1, 2, 3, 4, 5}
s2 := s1[:3]
s2 = append(s2, 100)
fmt.Println(s1)  // [1 2 3 100 5] - s1 modified!

// Fix: Use full slice expression
s2 := s1[:3:3]  // Limit capacity
s2 = append(s2, 100)  // Forces new array

// PITFALL 3: nil slice vs empty slice
var nilSlice []int
emptySlice := []int{}
fmt.Println(nilSlice == nil)   // true
fmt.Println(emptySlice == nil) // false
// JSON: nil -> null, empty -> []

// PITFALL 4: Map iteration during modification
m := map[string]int{"a": 1, "b": 2, "c": 3}
for k := range m {
    if k == "a" {
        delete(m, "b")  // Undefined behavior!
    }
}
// Fix: Collect keys first, then modify

// PITFALL 5: Range loop variable capture
var funcs []func()
for _, v := range []int{1, 2, 3} {
    funcs = append(funcs, func() {
        fmt.Println(v)  // All print 3!
    })
}
// Fix: Shadow variable or pass as argument

// PITFALL 6: Uninitialized nested map
m := make(map[string]map[string]int)
m["outer"]["inner"] = 1  // PANIC: nil map!
// Fix: Initialize inner map first
```

---

### **10. Next Steps & Wrap-up (2-3 min)**

**Recap what was covered:**
- Arrays (fixed, value type)
- Slices (dynamic, reference semantics)
- Slice internals (header, backing array)
- Common slice patterns
- Maps (key-value, hash table)
- Map patterns (sets, grouping)
- Performance tips
- In-memory database example

**Homework/Practice suggestions:**
1. **Easy:** Implement a sliding window function
2. **Medium:** Build a word frequency counter
3. **Challenge:** LRU cache with map and linked list
4. **Advanced:** Time-series database with efficient queries

---

## **Supplementary Materials**

**Cheat Sheet:**
```
Slice:
  make([]T, len)           Create with length
  make([]T, len, cap)      Create with capacity
  append(s, v)             Append (assign result!)
  copy(dst, src)           Copy elements
  s[low:high]              Slice expression
  s[low:high:max]          Full slice expression

Map:
  make(map[K]V)            Create map
  make(map[K]V, hint)      Create with size hint
  m[k] = v                 Set
  v := m[k]                Get (zero if missing)
  v, ok := m[k]            Get with existence check
  delete(m, k)             Delete
  len(m)                   Size
```

---

This tutorial covers Go's essential collection types in depth, with focus on internals and practical patterns. The in-memory database demonstrates real-world usage of both slices and maps together.
