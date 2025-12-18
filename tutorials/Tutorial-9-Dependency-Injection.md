## **Video Tutorial Plan: Dependency Injection in Go**

### **Video Metadata**
- **Title:** Dependency Injection in Go: Writing Testable, Maintainable Code
- **Duration Target:** 40-50 minutes
- **Difficulty:** Advanced
- **Prerequisites:** Interfaces, Structs, Error Handling

---

## **Video Structure**

### **1. Introduction (3-4 min)**
- What is dependency injection?
- Why DI matters for testing and maintainability
- Go's approach: interfaces and constructor injection
- Preview: Building a testable user service

---

### **2. The Problem: Tight Coupling (5-6 min)**

**Code Examples:**
```go
// TIGHTLY COUPLED - Hard to test
type UserService struct {
    // Direct dependency on concrete database
}

func (s *UserService) GetUser(id int) (*User, error) {
    // Directly using global database connection
    db, _ := sql.Open("postgres", "connection-string")
    defer db.Close()

    var user User
    err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Name, &user.Email)
    return &user, err
}

// Problems:
// 1. Can't test without real database
// 2. Can't swap implementations
// 3. Hard to mock for unit tests
// 4. Hidden dependencies
```

**Why this is problematic:**
- Tests require real database
- Can't test error conditions easily
- Slow tests (real I/O)
- Flaky tests (external dependencies)

---

### **3. Constructor Injection Pattern (8-10 min)**

**Topics to cover:**
- Define interfaces for dependencies
- Accept interfaces in constructors
- Store dependencies as fields

**Code Examples:**
```go
// Define interface for what you need
type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
    Delete(id int) error
}

// Service depends on interface, not concrete type
type UserService struct {
    repo   UserRepository
    logger Logger
}

// Constructor accepts interfaces
func NewUserService(repo UserRepository, logger Logger) *UserService {
    return &UserService{
        repo:   repo,
        logger: logger,
    }
}

// Methods use injected dependencies
func (s *UserService) GetUser(id int) (*User, error) {
    s.logger.Info("getting user", "id", id)

    user, err := s.repo.FindByID(id)
    if err != nil {
        s.logger.Error("failed to get user", "id", id, "error", err)
        return nil, fmt.Errorf("getting user %d: %w", id, err)
    }

    return user, nil
}

func (s *UserService) CreateUser(name, email string) (*User, error) {
    user := &User{Name: name, Email: email}

    if err := s.repo.Save(user); err != nil {
        return nil, fmt.Errorf("creating user: %w", err)
    }

    s.logger.Info("created user", "id", user.ID)
    return user, nil
}

// Production implementation
type PostgresUserRepository struct {
    db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
    return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindByID(id int) (*User, error) {
    var user User
    err := r.db.QueryRow(
        "SELECT id, name, email FROM users WHERE id = $1", id,
    ).Scan(&user.ID, &user.Name, &user.Email)
    if err == sql.ErrNoRows {
        return nil, ErrNotFound
    }
    return &user, err
}

// Wiring it up
func main() {
    db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    repo := NewPostgresUserRepository(db)
    logger := NewZapLogger()
    service := NewUserService(repo, logger)

    // Now use service
    user, _ := service.GetUser(1)
}
```

**Key teaching points:**
- Define small interfaces
- Accept interfaces, store interfaces
- Constructors wire dependencies
- Easy to swap implementations

---

### **4. Testing with Mocks (8-10 min)**

**Topics to cover:**
- Creating mock implementations
- Testing success cases
- Testing error cases
- Table-driven tests

**Code Examples:**
```go
// Mock implementation
type MockUserRepository struct {
    users    map[int]*User
    saveErr  error
    findErr  error
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users: make(map[int]*User),
    }
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
    if m.findErr != nil {
        return nil, m.findErr
    }
    user, ok := m.users[id]
    if !ok {
        return nil, ErrNotFound
    }
    return user, nil
}

func (m *MockUserRepository) Save(user *User) error {
    if m.saveErr != nil {
        return m.saveErr
    }
    if user.ID == 0 {
        user.ID = len(m.users) + 1
    }
    m.users[user.ID] = user
    return nil
}

func (m *MockUserRepository) Delete(id int) error {
    delete(m.users, id)
    return nil
}

// Helper methods for setting up test scenarios
func (m *MockUserRepository) SetFindError(err error) {
    m.findErr = err
}

func (m *MockUserRepository) AddUser(user *User) {
    m.users[user.ID] = user
}

// Mock logger
type MockLogger struct {
    Logs []LogEntry
}

type LogEntry struct {
    Level   string
    Message string
    Fields  map[string]interface{}
}

func (m *MockLogger) Info(msg string, fields ...interface{}) {
    m.Logs = append(m.Logs, LogEntry{Level: "info", Message: msg})
}

func (m *MockLogger) Error(msg string, fields ...interface{}) {
    m.Logs = append(m.Logs, LogEntry{Level: "error", Message: msg})
}

// Tests
func TestUserService_GetUser_Success(t *testing.T) {
    // Arrange
    repo := NewMockUserRepository()
    repo.AddUser(&User{ID: 1, Name: "Alice", Email: "alice@test.com"})
    logger := &MockLogger{}
    service := NewUserService(repo, logger)

    // Act
    user, err := service.GetUser(1)

    // Assert
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if user.Name != "Alice" {
        t.Errorf("expected Alice, got %s", user.Name)
    }
    if len(logger.Logs) == 0 || logger.Logs[0].Level != "info" {
        t.Error("expected info log")
    }
}

func TestUserService_GetUser_NotFound(t *testing.T) {
    repo := NewMockUserRepository()
    logger := &MockLogger{}
    service := NewUserService(repo, logger)

    _, err := service.GetUser(999)

    if !errors.Is(err, ErrNotFound) {
        t.Errorf("expected ErrNotFound, got %v", err)
    }
}

func TestUserService_GetUser_RepoError(t *testing.T) {
    repo := NewMockUserRepository()
    repo.SetFindError(errors.New("database connection lost"))
    logger := &MockLogger{}
    service := NewUserService(repo, logger)

    _, err := service.GetUser(1)

    if err == nil {
        t.Fatal("expected error")
    }
    // Check error logged
    hasErrorLog := false
    for _, log := range logger.Logs {
        if log.Level == "error" {
            hasErrorLog = true
            break
        }
    }
    if !hasErrorLog {
        t.Error("expected error to be logged")
    }
}

// Table-driven tests
func TestUserService_CreateUser(t *testing.T) {
    tests := []struct {
        name      string
        userName  string
        email     string
        saveErr   error
        wantErr   bool
    }{
        {
            name:     "success",
            userName: "Bob",
            email:    "bob@test.com",
            wantErr:  false,
        },
        {
            name:     "repository error",
            userName: "Bob",
            email:    "bob@test.com",
            saveErr:  errors.New("db error"),
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            repo := NewMockUserRepository()
            repo.saveErr = tt.saveErr
            logger := &MockLogger{}
            service := NewUserService(repo, logger)

            user, err := service.CreateUser(tt.userName, tt.email)

            if tt.wantErr && err == nil {
                t.Error("expected error")
            }
            if !tt.wantErr && err != nil {
                t.Errorf("unexpected error: %v", err)
            }
            if !tt.wantErr && user.Name != tt.userName {
                t.Errorf("expected %s, got %s", tt.userName, user.Name)
            }
        })
    }
}
```

**Key teaching points:**
- Mocks implement same interface
- Set up test scenarios with helper methods
- Test both success and error paths
- Table-driven tests for variations

---

### **5. Interface Design for DI (6-7 min)**

**Topics to cover:**
- Small, focused interfaces
- Consumer-defined interfaces
- Interface segregation

**Code Examples:**
```go
// BAD: Large interface (hard to mock)
type DatabaseInterface interface {
    Query(sql string, args ...interface{}) (*sql.Rows, error)
    Exec(sql string, args ...interface{}) (sql.Result, error)
    Begin() (*sql.Tx, error)
    Prepare(sql string) (*sql.Stmt, error)
    Ping() error
    Close() error
    // ... 20 more methods
}

// GOOD: Small, focused interface
type UserFinder interface {
    FindByID(id int) (*User, error)
}

type UserSaver interface {
    Save(user *User) error
}

// Combine when needed
type UserRepository interface {
    UserFinder
    UserSaver
    Delete(id int) error
}

// BEST: Define interface where it's used (consumer-side)

// In service package:
package service

// Only the methods this service needs
type userRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
}

type UserService struct {
    repo userRepository  // lowercase = private interface
}

// In repository package:
package repository

// Concrete implementation
type PostgresRepository struct {
    db *sql.DB
}

func (r *PostgresRepository) FindByID(id int) (*User, error) { ... }
func (r *PostgresRepository) Save(user *User) error { ... }
func (r *PostgresRepository) Delete(id int) error { ... }
func (r *PostgresRepository) FindAll() ([]*User, error) { ... }

// PostgresRepository has MORE methods than service needs
// Service only depends on what it uses

// INTERFACE SEGREGATION
// Different services need different things:

// AuthService only needs to find users
type userFinder interface {
    FindByEmail(email string) (*User, error)
}

type AuthService struct {
    users userFinder
}

// ReportService needs read-only access
type userReader interface {
    FindAll() ([]*User, error)
    Count() (int, error)
}

type ReportService struct {
    users userReader
}

// AdminService needs full access
type userAdmin interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
    Delete(id int) error
}

type AdminService struct {
    users userAdmin
}

// One concrete repository satisfies all three interfaces!
```

**Key teaching points:**
- Small interfaces = easier mocking
- Consumer defines what it needs
- One implementation can satisfy many interfaces
- Don't export interfaces unnecessarily

---

### **6. Functional Options Pattern (5-6 min)**

**Topics to cover:**
- Optional dependencies
- Configuration options
- Self-documenting constructors

**Code Examples:**
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

// Usage - very readable!
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

---

### **7. Practical Example: Complete Application (10-12 min)**

**Build together:** A complete user management application

```go
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
    log := s.logger.With("method", "GetUser", "user_id", id)

    // Check cache
    cacheKey := fmt.Sprintf("user:%d", id)
    if cached, ok := s.cache.Get(ctx, cacheKey); ok {
        log.Info("cache hit")
        return cached.(*User), nil
    }

    // Get from repository
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        log.Error("failed to get user", "error", err)
        return nil, fmt.Errorf("getting user: %w", err)
    }

    // Update cache
    s.cache.Set(ctx, cacheKey, user, s.cacheTTL)
    log.Info("user retrieved")

    return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, email, name string) (*User, error) {
    log := s.logger.With("method", "CreateUser", "email", email)

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
        log.Error("failed to save user", "error", err)
        return nil, fmt.Errorf("saving user: %w", err)
    }

    // Publish event
    event := UserCreatedEvent{
        UserID:    user.ID,
        Email:     user.Email,
        Timestamp: time.Now(),
    }
    if err := s.events.Publish(ctx, event); err != nil {
        log.Error("failed to publish event", "error", err)
        // Don't fail the operation for event publishing
    }

    log.Info("user created", "user_id", user.ID)
    return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
    log := s.logger.With("method", "DeleteUser", "user_id", id)

    // Delete from repo
    if err := s.repo.Delete(ctx, id); err != nil {
        log.Error("failed to delete user", "error", err)
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

    log.Info("user deleted")
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

---

### **8. Dependency Injection Tools (4-5 min)**

**Topics to cover:**
- Google Wire
- Manual vs generated DI
- When to use tools

**Code Examples:**
```go
// Google Wire - compile-time DI

// wire.go (build tag: wireinject)
//go:build wireinject

package main

import "github.com/google/wire"

func InitializeApp() (*App, error) {
    wire.Build(
        NewConfig,
        NewDatabase,
        NewUserRepository,
        NewUserService,
        NewHTTPServer,
        NewApp,
    )
    return nil, nil
}

// Providers
func NewDatabase(cfg *Config) (*sql.DB, error) {
    return sql.Open("postgres", cfg.DatabaseURL)
}

func NewUserRepository(db *sql.DB) *PostgresUserRepository {
    return &PostgresUserRepository{db: db}
}

func NewUserService(repo UserRepository, logger Logger) *UserService {
    return &UserService{repo: repo, logger: logger}
}

// Wire generates wire_gen.go with actual wiring code

// When to use DI tools:
// - Large applications with many dependencies
// - Complex dependency graphs
// - Want compile-time safety

// When manual DI is fine:
// - Small to medium applications
// - Simple dependency graphs
// - Prefer explicit over magic
```

---

### **9. Best Practices Summary (3-4 min)**

```go
// 1. Accept interfaces, return structs
func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

// 2. Define interfaces where they're used
type service struct {
    repo userRepository  // Interface defined in this package
}

// 3. Keep interfaces small
type Saver interface {
    Save(v interface{}) error
}

// 4. Use constructor injection
func NewService(dep1 Dep1, dep2 Dep2) *Service

// 5. Provide no-op defaults for optional deps
func NewService(repo Repository, opts ...Option) *Service

// 6. Don't inject everything
// Inject: External services, databases, configuration
// Don't inject: Utility functions, time.Now(), etc.

// 7. Use context for request-scoped values
func (s *Service) DoThing(ctx context.Context) error

// 8. Test with mocks, not real implementations
repo := &MockRepository{}
service := NewService(repo)
```

---

### **10. Next Steps & Wrap-up (2-3 min)**

**Recap:**
- Why DI matters
- Constructor injection
- Testing with mocks
- Interface design
- Functional options
- Complete application example

**Homework:**
1. Refactor existing code to use DI
2. Add caching layer with DI
3. Build testable HTTP handlers
4. Implement with Google Wire

---

## **Supplementary Materials**

**Cheat Sheet:**
```
Constructor Injection:
  func NewService(dep Interface) *Service

Functional Options:
  type Option func(*Service)
  func NewService(opts ...Option) *Service

Interface Location:
  Consumer defines interface it needs

Testing:
  Mock implements same interface
  Test success and error paths
```

---

This tutorial covers DI comprehensively with practical patterns and a complete example showing how all pieces fit together.
