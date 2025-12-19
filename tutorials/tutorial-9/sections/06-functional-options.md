# Functional Options Pattern

**Duration:** 5-6 minutes

## Topics to cover:
- Optional dependencies
- Configuration options
- Self-documenting constructors

## Code Examples

```go
// Option function type
type Option func(*Service)

// Service with optional dependencies
type Service struct {
    repo      Repository
    cache     Cache
    logger    Logger
    timeout   time.Duration
    retries   int
}

// Option functions
func WithCache(cache Cache) Option {
    return func(s *Service) {
        s.cache = cache
    }
}

func WithLogger(logger Logger) Option {
    return func(s *Service) {
        s.logger = logger
    }
}

func WithTimeout(timeout time.Duration) Option {
    return func(s *Service) {
        s.timeout = timeout
    }
}

func WithRetries(retries int) Option {
    return func(s *Service) {
        s.retries = retries
    }
}

// Constructor applies options
func NewService(repo Repository, opts ...Option) *Service {
    // Defaults
    s := &Service{
        repo:    repo,
        cache:   &NoOpCache{},    // Default: no caching
        logger:  &NoOpLogger{},   // Default: no logging
        timeout: 30 * time.Second,
        retries: 3,
    }

    // Apply options
    for _, opt := range opts {
        opt(s)
    }

    return s
}
```

## Usage - Very Readable!

```go
func main() {
    repo := NewPostgresRepo(db)

    // Minimal setup
    svc1 := NewService(repo)

    // With all options
    svc2 := NewService(repo,
        WithCache(NewRedisCache(redis)),
        WithLogger(NewZapLogger()),
        WithTimeout(10*time.Second),
        WithRetries(5),
    )

    // Just what you need
    svc3 := NewService(repo,
        WithLogger(NewZapLogger()),
    )
}
```
