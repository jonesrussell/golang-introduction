# Practical Example: In-Memory Database

**Duration:** 8-10 minutes

## Build Together

A simple in-memory database demonstrating slices and maps.

```go runnable
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

func main() {
    fmt.Println("=== In-Memory Database Demo ===\n")

    db := NewDatabase()

    // Insert records
    r1, _ := db.Insert("Alice", "alice@example.com", []string{"admin", "developer"})
    _, _ = db.Insert("Bob", "bob@example.com", []string{"developer"})
    _, _ = db.Insert("Charlie", "charlie@example.com", []string{"designer", "developer"})
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

    // All records
    fmt.Println("\nAll records:")
    for _, r := range db.All() {
        fmt.Printf("  [%d] %s <%s> %v\n", r.ID, r.Name, r.Email, r.Tags)
    }
}
```

## Walk Through:
- Primary storage with map
- Secondary indexes for fast lookups
- Thread-safe with RWMutex
- Query function for flexible filtering
- Proper cleanup when deleting
