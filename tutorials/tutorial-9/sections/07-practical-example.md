# Practical Example: Complete Application

**Duration:** 10-12 minutes

## Build Together

A complete user management application.

```go runnable
package main

import (
    "context"
    "errors"
    "fmt"
    "time"
)

// ========================================
// Domain
// ========================================

var (
    ErrNotFound       = errors.New("not found")
    ErrAlreadyExists  = errors.New("already exists")
    ErrInvalidInput   = errors.New("invalid input")
)

type User struct {
    ID        int
    Email     string
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// ========================================
// Interfaces (defined by consumers)
// ========================================

type UserRepository interface {
    FindByID(ctx context.Context, id int) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Save(ctx context.Context, user *User) error
    Delete(ctx context.Context, id int) error
}

type Logger interface {
    Info(msg string, fields ...interface{})
    Error(msg string, fields ...interface{})
    With(fields ...interface{}) Logger
}

type Cache interface {
    Get(ctx context.Context, key string) (interface{}, bool)
    Set(ctx context.Context, key string, value interface{}, ttl time.Duration)
    Delete(ctx context.Context, key string)
}

type EventPublisher interface {
    Publish(ctx context.Context, event interface{}) error
}

// ========================================
// Events
// ========================================

type UserCreatedEvent struct {
    UserID    int
    Email     string
    Timestamp time.Time
}

type UserDeletedEvent struct {
    UserID    int
    Timestamp time.Time
}

// ========================================
// Service
// ========================================

type UserService struct {
    repo      UserRepository
    cache     Cache
    events    EventPublisher
    logger    Logger
    cacheTTL  time.Duration
}

type UserServiceOption func(*UserService)

func WithUserCache(cache Cache, ttl time.Duration) UserServiceOption {
    return func(s *UserService) {
        s.cache = cache
        s.cacheTTL = ttl
    }
}

func WithEventPublisher(pub EventPublisher) UserServiceOption {
    return func(s *UserService) {
        s.events = pub
    }
}

func WithUserLogger(logger Logger) UserServiceOption {
    return func(s *UserService) {
        s.logger = logger
    }
}

func NewUserService(repo UserRepository, opts ...UserServiceOption) *UserService {
    s := &UserService{
        repo:     repo,
        cache:    &noOpCache{},
        events:   &noOpPublisher{},
        logger:   &noOpLogger{},
        cacheTTL: 5 * time.Minute,
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}

func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
    // Check cache
    cacheKey := fmt.Sprintf("user:%d", id)
    if cached, ok := s.cache.Get(ctx, cacheKey); ok {
        s.logger.Info("cache hit", "user_id", id)
        return cached.(*User), nil
    }

    // Get from repository
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        s.logger.Error("failed to get user", "user_id", id, "error", err)
        return nil, fmt.Errorf("getting user: %w", err)
    }

    // Update cache
    s.cache.Set(ctx, cacheKey, user, s.cacheTTL)
    s.logger.Info("user retrieved", "user_id", id)

    return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, email, name string) (*User, error) {
    // Validate
    if email == "" || name == "" {
        return nil, fmt.Errorf("%w: email and name required", ErrInvalidInput)
    }

    // Check if exists
    existing, err := s.repo.FindByEmail(ctx, email)
    if err != nil && !errors.Is(err, ErrNotFound) {
        return nil, fmt.Errorf("checking email: %w", err)
    }
    if existing != nil {
        return nil, fmt.Errorf("email %s: %w", email, ErrAlreadyExists)
    }

    // Create user
    user := &User{
        Email:     email,
        Name:      name,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    if err := s.repo.Save(ctx, user); err != nil {
        s.logger.Error("failed to save user", "error", err)
        return nil, fmt.Errorf("saving user: %w", err)
    }

    // Publish event
    event := UserCreatedEvent{
        UserID:    user.ID,
        Email:     user.Email,
        Timestamp: time.Now(),
    }
    if err := s.events.Publish(ctx, event); err != nil {
        s.logger.Error("failed to publish event", "error", err)
    }

    s.logger.Info("user created", "user_id", user.ID)
    return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
    if err := s.repo.Delete(ctx, id); err != nil {
        s.logger.Error("failed to delete user", "user_id", id, "error", err)
        return fmt.Errorf("deleting user: %w", err)
    }

    // Invalidate cache
    cacheKey := fmt.Sprintf("user:%d", id)
    s.cache.Delete(ctx, cacheKey)

    // Publish event
    event := UserDeletedEvent{
        UserID:    id,
        Timestamp: time.Now(),
    }
    s.events.Publish(ctx, event)

    s.logger.Info("user deleted", "user_id", id)
    return nil
}

// ========================================
// No-op implementations (defaults)
// ========================================

type noOpCache struct{}
func (n *noOpCache) Get(ctx context.Context, key string) (interface{}, bool) { return nil, false }
func (n *noOpCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) {}
func (n *noOpCache) Delete(ctx context.Context, key string) {}

type noOpPublisher struct{}
func (n *noOpPublisher) Publish(ctx context.Context, event interface{}) error { return nil }

type noOpLogger struct{}
func (n *noOpLogger) Info(msg string, fields ...interface{}) {}
func (n *noOpLogger) Error(msg string, fields ...interface{}) {}
func (n *noOpLogger) With(fields ...interface{}) Logger { return n }

// ========================================
// Test Implementation
// ========================================

type InMemoryUserRepository struct {
    users  map[int]*User
    nextID int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
    return &InMemoryUserRepository{
        users:  make(map[int]*User),
        nextID: 1,
    }
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id int) (*User, error) {
    user, ok := r.users[id]
    if !ok {
        return nil, ErrNotFound
    }
    return user, nil
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
    for _, user := range r.users {
        if user.Email == email {
            return user, nil
        }
    }
    return nil, ErrNotFound
}

func (r *InMemoryUserRepository) Save(ctx context.Context, user *User) error {
    if user.ID == 0 {
        user.ID = r.nextID
        r.nextID++
    }
    r.users[user.ID] = user
    return nil
}

func (r *InMemoryUserRepository) Delete(ctx context.Context, id int) error {
    if _, ok := r.users[id]; !ok {
        return ErrNotFound
    }
    delete(r.users, id)
    return nil
}

// ========================================
// Main
// ========================================

func main() {
    fmt.Println("=== Dependency Injection Demo ===\n")

    // Create dependencies
    repo := NewInMemoryUserRepository()

    // Create service with minimal dependencies
    service := NewUserService(repo)

    ctx := context.Background()

    // Create user
    user, err := service.CreateUser(ctx, "alice@example.com", "Alice")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Created user: %+v\n", user)

    // Get user
    fetched, err := service.GetUser(ctx, user.ID)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Fetched user: %+v\n", fetched)

    // Try duplicate
    _, err = service.CreateUser(ctx, "alice@example.com", "Alice 2")
    fmt.Printf("Duplicate error: %v\n", err)

    // Delete
    err = service.DeleteUser(ctx, user.ID)
    fmt.Printf("Delete result: %v\n", err)

    // Get deleted (should fail)
    _, err = service.GetUser(ctx, user.ID)
    fmt.Printf("Get deleted: %v\n", err)
}
```
