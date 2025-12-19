# sync Package Essentials

**Duration:** 6-7 minutes

## Topics to cover:
- [`sync.WaitGroup`](https://pkg.go.dev/sync#WaitGroup)
- [`sync.Mutex`](https://pkg.go.dev/sync#Mutex)
- [`sync.RWMutex`](https://pkg.go.dev/sync#RWMutex)
- [`sync.Once`](https://pkg.go.dev/sync#Once)

## WaitGroup

```go
import "sync"

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }

    wg.Wait()  // Block until all workers done
    fmt.Println("All workers completed")
}
```

## Mutex

```go
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := &Counter{}
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    wg.Wait()
    fmt.Println("Final count:", counter.Value())  // Always 1000
}
```

## RWMutex

```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()  // Read lock - allows concurrent reads
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()  // Write lock - exclusive access
    defer c.mu.Unlock()
    c.data[key] = value
}
```

## sync.Once

```go
var (
    instance *Database
    once     sync.Once
)

func GetDatabase() *Database {
    once.Do(func() {
        fmt.Println("Initializing database...")
        instance = &Database{}
    })
    return instance
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            db := GetDatabase()  // Only initializes once
            _ = db
        }()
    }
    wg.Wait()
}
```

## Key teaching points:
- [`WaitGroup`](https://pkg.go.dev/sync#WaitGroup) for waiting on multiple goroutines
- [`Mutex`](https://pkg.go.dev/sync#Mutex) for protecting shared state
- [`RWMutex`](https://pkg.go.dev/sync#RWMutex) when reads >> writes
- [`Once`](https://pkg.go.dev/sync#Once) for one-time initialization
- Always use defer with Lock/Unlock
