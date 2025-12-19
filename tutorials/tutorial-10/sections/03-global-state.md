# Global State

**Duration:** 6-7 minutes

## The Anti-Pattern

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

## Problems:
- Hidden dependencies
- Impossible to test
- Race conditions
- Initialization order issues

## The Fix

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
```

## Acceptable Globals:
- Package-level errors: `var ErrNotFound = errors.New(...)`
- [`sync.Once`](https://pkg.go.dev/sync#Once) for lazy initialization
- Compiled regexes

## Key teaching points:
- Avoid global mutable state
- Use [dependency injection](https://go.dev/doc/effective_go#interfaces_and_types) for dependencies
- Global constants and [sentinel errors](https://pkg.go.dev/errors#New) are acceptable
- [sync.Once](https://pkg.go.dev/sync#Once) for thread-safe lazy initialization
