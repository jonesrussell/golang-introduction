## **Video Tutorial Plan: Avoiding Common Go Anti-Patterns**

### **Video Metadata**
- **Title:** Avoiding Common Go Anti-Patterns
- **Duration Target:** 35-45 minutes
- **Difficulty:** Advanced
- **Prerequisites:** Go Basics through Concurrency

---

## **Video Structure**

### **1. Introduction (2-3 min)**
- What are anti-patterns?
- Why they matter for maintainability
- Go-specific pitfalls
- Preview: Before and after refactoring

---

### **2. Context.Value Abuse (6-7 min)**

**The Anti-Pattern:**
```go
// BAD: Using context for everything
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    ctx = context.WithValue(ctx, "userID", 123)
    ctx = context.WithValue(ctx, "requestID", "abc-123")
    ctx = context.WithValue(ctx, "permissions", []string{"read", "write"})
    ctx = context.WithValue(ctx, "config", &Config{})
    ctx = context.WithValue(ctx, "logger", logger)
    ctx = context.WithValue(ctx, "db", database)

    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    // Type assertions everywhere, no compile-time safety
    userID := ctx.Value("userID").(int)
    config := ctx.Value("config").(*Config)
    logger := ctx.Value("logger").(Logger)
}
```

**Problems:**
- No type safety (runtime panics)
- Hidden dependencies
- Hard to test
- Unclear API contracts

**The Fix:**
```go
// GOOD: Explicit parameters for dependencies
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    // Context only for: cancellation, deadlines, request-scoped values
    ctx = context.WithValue(ctx, requestIDKey, "abc-123")

    userID := getUserID(r)
    processRequest(ctx, userID, s.config, s.logger, s.db)
}

// Explicit dependencies
func processRequest(
    ctx context.Context,
    userID int,
    config *Config,
    logger Logger,
    db Database,
) error {
    // Clear what this function needs
}

// Use typed keys for context values
type contextKey string
const requestIDKey contextKey = "requestID"

// Use context.Value for:
// - Request ID / Trace ID
// - Cancellation signals
// - Deadlines
// - Request-scoped auth info

// DON'T use for:
// - Dependencies (database, logger)
// - Configuration
// - Business logic data
```

---

### **3. Global State (6-7 min)**

**The Anti-Pattern:**
```go
// BAD: Global variables everywhere
var (
    db     *sql.DB
    config *Config
    logger *Logger
    cache  *Cache
)

func init() {
    db = connectDB()
    config = loadConfig()
    logger = setupLogger()
}

func GetUser(id int) (*User, error) {
    // Uses global db - hidden dependency
    return db.Query("SELECT * FROM users WHERE id = ?", id)
}

func main() {
    user, _ := GetUser(1)  // Which database? Which config?
}
```

**Problems:**
- Hidden dependencies
- Impossible to test
- Race conditions
- Initialization order issues

**The Fix:**
```go
// GOOD: Explicit dependency injection
type UserService struct {
    db     *sql.DB
    logger Logger
    cache  Cache
}

func NewUserService(db *sql.DB, logger Logger, cache Cache) *UserService {
    return &UserService{db: db, logger: logger, cache: cache}
}

func (s *UserService) GetUser(id int) (*User, error) {
    // Clear where data comes from
    return s.db.Query("SELECT * FROM users WHERE id = ?", id)
}

func main() {
    db := connectDB(config.DatabaseURL)
    logger := setupLogger(config.LogLevel)
    cache := setupCache(config.CacheURL)

    userService := NewUserService(db, logger, cache)
    user, _ := userService.GetUser(1)
}

// Acceptable globals:
// - Package-level errors (var ErrNotFound = errors.New(...))
// - sync.Once for lazy initialization
// - Compiled regexes
```

---

### **4. Interface Pollution (5-6 min)**

**The Anti-Pattern:**
```go
// BAD: Interface for every struct
type UserServiceInterface interface {
    GetUser(id int) (*User, error)
    CreateUser(user *User) error
    UpdateUser(user *User) error
    DeleteUser(id int) error
}

type UserService struct{}

func (s *UserService) GetUser(id int) (*User, error) { ... }
// ... more methods

// Only one implementation exists!

// BAD: Exporting interfaces from producer package
package repository

type UserRepository interface {  // Exported but unnecessary
    Find(id int) (*User, error)
}

type PostgresUserRepository struct{}
```

**The Fix:**
```go
// GOOD: Return concrete types
func NewUserService() *UserService {
    return &UserService{}
}

// GOOD: Define interfaces at point of use (consumer)
package handler

// Interface defined where it's used
type userGetter interface {
    GetUser(id int) (*User, error)
}

type Handler struct {
    users userGetter
}

// GOOD: Only create interface when needed
// - Multiple implementations exist
// - Testing requires mocking
// - Package boundary crossing

// Accept interfaces, return structs
func NewHandler(users userGetter) *Handler {
    return &Handler{users: users}
}
```

---

### **5. Nil Pointer Paranoia / Over-Checking (4-5 min)**

**The Anti-Pattern:**
```go
// BAD: Nil checks everywhere
func ProcessUser(user *User) error {
    if user == nil {
        return errors.New("user is nil")
    }
    if user.Profile == nil {
        return errors.New("profile is nil")
    }
    if user.Profile.Address == nil {
        return errors.New("address is nil")
    }
    if user.Profile.Address.City == nil {
        return errors.New("city is nil")
    }

    city := *user.Profile.Address.City
    // ...
}
```

**The Fix:**
```go
// GOOD: Design to avoid nil
type User struct {
    Profile Profile  // Value, not pointer - never nil
}

type Profile struct {
    Address Address  // Value, not pointer
}

type Address struct {
    City string  // Value, not pointer
}

// GOOD: Use constructor to ensure valid state
func NewUser(name string) *User {
    return &User{
        Name: name,
        Profile: Profile{
            Address: Address{
                City: "Unknown",
            },
        },
    }
}

// GOOD: Check at boundaries, trust internal code
func (h *Handler) HandleRequest(r *http.Request) {
    // Validate input at boundary
    user, err := parseUser(r)
    if err != nil {
        // Handle invalid input
        return
    }

    // Internal code can trust user is valid
    h.service.ProcessUser(user)
}

// When pointer is intentional (optional field)
type Config struct {
    Timeout *time.Duration  // nil means "use default"
}

func (c *Config) GetTimeout() time.Duration {
    if c.Timeout == nil {
        return 30 * time.Second
    }
    return *c.Timeout
}
```

---

### **6. Error String Matching (4-5 min)**

**The Anti-Pattern:**
```go
// BAD: String matching for error handling
func HandleError(err error) {
    if err.Error() == "user not found" {
        // Handle not found
    }
    if strings.Contains(err.Error(), "timeout") {
        // Handle timeout
    }
    if strings.HasPrefix(err.Error(), "validation") {
        // Handle validation
    }
}
```

**Problems:**
- Fragile (error message changes break code)
- No compile-time safety
- Doesn't work with wrapped errors

**The Fix:**
```go
// GOOD: Sentinel errors
var (
    ErrNotFound   = errors.New("user not found")
    ErrTimeout    = errors.New("operation timed out")
    ErrValidation = errors.New("validation failed")
)

func HandleError(err error) {
    if errors.Is(err, ErrNotFound) {
        // Handle not found
    }
    if errors.Is(err, ErrTimeout) {
        // Handle timeout
    }
}

// GOOD: Custom error types
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func HandleError(err error) {
    var valErr *ValidationError
    if errors.As(err, &valErr) {
        fmt.Printf("Invalid field: %s\n", valErr.Field)
    }
}
```

---

### **7. Goroutine Leaks (5-6 min)**

**The Anti-Pattern:**
```go
// BAD: Goroutine that can't exit
func startWorker() {
    go func() {
        for {
            // Process forever
            item := <-workQueue  // Blocks forever if queue closes
            process(item)
        }
    }()
}

// BAD: Unbounded channel producer
func producer() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; ; i++ {
            ch <- i  // Blocks forever if no consumer
        }
    }()
    return ch
}

// BAD: Fire and forget with unbuffered channel
func process() {
    ch := make(chan result)
    go func() {
        r := doWork()
        ch <- r  // Blocks forever if main doesn't read
    }()

    // Timeout - goroutine leaks!
    select {
    case r := <-ch:
        return r
    case <-time.After(timeout):
        return nil  // Goroutine still blocked on send!
    }
}
```

**The Fix:**
```go
// GOOD: Goroutine with cancellation
func startWorker(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return  // Clean exit
            case item := <-workQueue:
                process(item)
            }
        }
    }()
}

// GOOD: Buffered channel for fire-and-forget
func process(ctx context.Context) *result {
    ch := make(chan *result, 1)  // Buffered!
    go func() {
        r := doWork()
        ch <- r  // Won't block even if nobody reads
    }()

    select {
    case r := <-ch:
        return r
    case <-ctx.Done():
        return nil  // Goroutine can still complete
    }
}

// GOOD: WaitGroup for cleanup
func processAll(items []Item) {
    var wg sync.WaitGroup
    for _, item := range items {
        wg.Add(1)
        go func(i Item) {
            defer wg.Done()
            process(i)
        }(item)
    }
    wg.Wait()  // Ensure all goroutines complete
}
```

---

### **8. Premature Optimization (4-5 min)**

**The Anti-Pattern:**
```go
// BAD: Complex "optimization" without measurement
func processData(data []byte) {
    // "Optimized" with sync.Pool
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    buf.Reset()

    // Pre-allocated slice
    result := make([]byte, 0, len(data)*2)

    // Manual loop "faster than range"
    for i := 0; i < len(data); i++ {
        // ...
    }
}

// When this simple version works fine:
func processDataSimple(data []byte) {
    var buf bytes.Buffer
    buf.Write(data)
    // ...
}
```

**The Fix:**
```go
// GOOD: Write clear code first
func ProcessItems(items []Item) []Result {
    results := make([]Result, 0, len(items))
    for _, item := range items {
        result := process(item)
        results = append(results, result)
    }
    return results
}

// GOOD: Optimize with evidence
// 1. Profile first: go test -bench . -cpuprofile cpu.out
// 2. Identify bottlenecks: go tool pprof cpu.out
// 3. Optimize only hot paths
// 4. Measure improvement

// Acceptable early optimizations:
// - Pre-allocate when size is known
// - Use strings.Builder for concatenation
// - Choose appropriate data structure
```

---

### **9. Mutex Misuse (4-5 min)**

**The Anti-Pattern:**
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

**The Fix:**
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

---

### **10. Init Function Abuse (3-4 min)**

**The Anti-Pattern:**
```go
// BAD: Complex init with errors
func init() {
    db, err := sql.Open("postgres", os.Getenv("DB_URL"))
    if err != nil {
        panic(err)  // Crashes on startup
    }
    globalDB = db

    config, err := loadConfig()
    if err != nil {
        panic(err)
    }
    globalConfig = config
}
```

**The Fix:**
```go
// GOOD: Explicit initialization in main
func main() {
    config, err := loadConfig()
    if err != nil {
        log.Fatalf("loading config: %v", err)
    }

    db, err := setupDatabase(config)
    if err != nil {
        log.Fatalf("connecting to database: %v", err)
    }
    defer db.Close()

    server := NewServer(config, db)
    server.Run()
}

// Acceptable init uses:
// - Register drivers: sql.Register, http.Handle
// - Compile regexes
// - Set package-level computed constants
```

---

### **11. Recap & Best Practices (2-3 min)**

**Summary:**
1. Use context only for cancellation and request-scoped values
2. Inject dependencies explicitly
3. Define interfaces at point of use
4. Design types to avoid nil
5. Use sentinel errors and custom error types
6. Always provide goroutine exit paths
7. Profile before optimizing
8. Keep mutex scope minimal
9. Initialize explicitly in main

---

## **Supplementary Materials**

**Anti-Pattern Checklist:**
```
[ ] Context.Value for dependencies?
[ ] Global mutable state?
[ ] Interface for single implementation?
[ ] String matching for errors?
[ ] Goroutines without exit path?
[ ] Mutex with value receiver?
[ ] Complex logic in init()?
```

---

This tutorial identifies the most common Go anti-patterns and provides clear guidance on how to avoid them.
