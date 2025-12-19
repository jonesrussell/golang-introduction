# Practical Patterns with Embedding

**Duration:** 8-10 minutes

## Topics to cover:
- Mixin pattern
- Decorator pattern
- Base functionality pattern
- When to use embedding vs composition

## Mixin Pattern

```go
// Mixin pattern - adding common functionality

// Audit mixin - adds audit fields to any struct
type Auditable struct {
    CreatedAt time.Time
    CreatedBy string
    UpdatedAt time.Time
    UpdatedBy string
}

func (a *Auditable) SetCreated(by string) {
    a.CreatedAt = time.Now()
    a.CreatedBy = by
}

func (a *Auditable) SetUpdated(by string) {
    a.UpdatedAt = time.Now()
    a.UpdatedBy = by
}

// Soft delete mixin
type SoftDeletable struct {
    DeletedAt *time.Time
    DeletedBy string
}

func (s *SoftDeletable) Delete(by string) {
    now := time.Now()
    s.DeletedAt = &now
    s.DeletedBy = by
}

func (s SoftDeletable) IsDeleted() bool {
    return s.DeletedAt != nil
}

// Use mixins by embedding
type Product struct {
    Auditable      // Mixin
    SoftDeletable  // Mixin
    ID             int
    Name           string
    Price          float64
}

type Order struct {
    Auditable      // Reuse same mixins
    SoftDeletable
    ID             int
    ProductID      int
    Quantity       int
}

// Usage
product := Product{
    ID:    1,
    Name:  "Laptop",
    Price: 999.99,
}

product.SetCreated("admin")
product.SetUpdated("admin")
product.Delete("admin")

if product.IsDeleted() {
    fmt.Println("Product is soft deleted")
}

fmt.Printf("Created: %v by %s\n", 
    product.CreatedAt, product.CreatedBy)
```

## Decorator Pattern

```go
// Decorator pattern - adding behavior

type Logger interface {
    Log(message string)
}

type SimpleLogger struct{}

func (s SimpleLogger) Log(message string) {
    fmt.Println(message)
}

// Decorator that adds timestamps
type TimestampLogger struct {
    Logger  // Embedded interface
}

func (t TimestampLogger) Log(message string) {
    // Add timestamp before delegating
    timestamped := fmt.Sprintf("[%s] %s", 
        time.Now().Format("15:04:05"), message)
    t.Logger.Log(timestamped)
}

// Decorator that adds level
type LevelLogger struct {
    Logger
    Level string
}

func (l LevelLogger) Log(message string) {
    leveled := fmt.Sprintf("[%s] %s", l.Level, message)
    l.Logger.Log(leveled)
}

// Usage - compose decorators
logger := LevelLogger{
    Logger: TimestampLogger{
        Logger: SimpleLogger{},
    },
    Level: "INFO",
}

logger.Log("Application started")
// Output: [INFO] [15:04:05] Application started
```

## Base Entity Pattern

```go
// Base entity pattern - common database entity behavior

type BaseEntity struct {
    ID        int
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (b *BaseEntity) BeforeSave() {
    now := time.Now()
    if b.CreatedAt.IsZero() {
        b.CreatedAt = now
    }
    b.UpdatedAt = now
}

type Saveable interface {
    BeforeSave()
}

// All entities embed BaseEntity
type User struct {
    BaseEntity
    Username string
    Email    string
}

type Post struct {
    BaseEntity
    Title   string
    Content string
    UserID  int
}

type Comment struct {
    BaseEntity
    Content string
    PostID  int
    UserID  int
}

// Generic save function works with any Saveable
func Save(s Saveable) {
    s.BeforeSave()
    fmt.Println("Saving to database...")
    // Database save logic
}

// Usage
user := &User{
    Username: "alice",
    Email:    "alice@example.com",
}

post := &Post{
    Title:   "Go Embedding",
    Content: "Learn about embedding...",
    UserID:  user.ID,
}

Save(user)  // Calls BeforeSave on BaseEntity
Save(post)  // Same interface, different type

fmt.Printf("User created at: %v\n", user.CreatedAt)
fmt.Printf("Post created at: %v\n", post.CreatedAt)
```

## Key teaching points:
- Mixin pattern: Embed common functionality across multiple types
- Decorator pattern: Add behavior by wrapping
- Base entity: Share common behavior across related types
- DRY: Don't repeat timestamp, audit, soft-delete logic
- Embedding promotes code reuse without inheritance
