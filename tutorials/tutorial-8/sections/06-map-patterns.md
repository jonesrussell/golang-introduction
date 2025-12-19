# Advanced Map Patterns

**Duration:** 5-6 minutes

## Topics to cover:
- Sets using maps
- Counting/grouping
- Cache patterns
- Map of slices/maps

## Set Implementation

```go runnable
package main

import "fmt"

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

func main() {
    set := NewSet()
    set.Add("apple")
    set.Add("banana")
    set.Add("apple")  // Duplicate, ignored
    
    fmt.Println("Contains apple:", set.Contains("apple"))  // true
    fmt.Println("Contains cherry:", set.Contains("cherry"))  // false
    fmt.Println("Size:", set.Size())  // 2
}
```

## Counting / Frequency

```go runnable
package main

import "fmt"

func main() {
    words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
    count := make(map[string]int)
    
    for _, word := range words {
        count[word]++  // Zero value (0) works here!
    }
    
    fmt.Println("Word counts:")
    for word, c := range count {
        fmt.Printf("  %s: %d\n", word, c)
    }
}
```

## Grouping

```go runnable
package main

import "fmt"

type Person struct {
    Name    string
    Country string
}

func main() {
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

    fmt.Println("People by country:")
    for country, persons := range byCountry {
        fmt.Printf("  %s:\n", country)
        for _, p := range persons {
            fmt.Printf("    - %s\n", p.Name)
        }
    }
}
```

## Simple Cache

```go
import "sync"

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
```

## Nested Maps

```go
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
```

## Default Value Pattern

```go
func getWithDefault(m map[string]int, key string, defaultVal int) int {
    if val, ok := m[key]; ok {
        return val
    }
    return defaultVal
}
```
