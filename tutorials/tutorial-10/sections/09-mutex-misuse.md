# Mutex Misuse

**Duration:** 4-5 minutes

## The Anti-Pattern

```go
// BAD: Copying mutex
type Counter struct {
    sync.Mutex
    count int
}

func (c Counter) Increment() {  // Value receiver copies mutex!
    c.Lock()
    c.count++
    c.Unlock()
}

// BAD: Holding lock too long
func (s *Service) ProcessAll() {
    s.mu.Lock()
    defer s.mu.Unlock()

    for _, item := range s.items {
        s.processItem(item)  // Slow operation under lock!
        s.callExternalAPI()  // Network call under lock!
    }
}

// BAD: Nested locks (deadlock risk)
func (s *Service) Update() {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.helper.DoSomething()  // If helper locks, potential deadlock
}
```

## The Fix

```go
// GOOD: Pointer receiver with mutex
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Increment() {  // Pointer receiver
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

// GOOD: Minimize lock scope
func (s *Service) ProcessAll() {
    s.mu.Lock()
    items := make([]Item, len(s.items))
    copy(items, s.items)  // Copy under lock
    s.mu.Unlock()

    // Process outside lock
    for _, item := range items {
        s.processItem(item)
    }
}

// GOOD: Use RWMutex for read-heavy workloads
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()  // Multiple readers allowed
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()  // Exclusive for writes
    defer c.mu.Unlock()
    c.data[key] = value
}
```

## Key teaching points:
- Never copy [mutexes](https://pkg.go.dev/sync#Mutex) (use pointer receivers)
- Keep [lock scope](https://pkg.go.dev/sync#Mutex) minimal
- Use [`RWMutex`](https://pkg.go.dev/sync#RWMutex) for read-heavy workloads
- Avoid nested locks to prevent deadlocks
- Always use [`defer`](https://go.dev/ref/spec#Defer_statements) with Lock/Unlock
