## **Video Tutorial Plan: Understanding Go Interfaces**

### **Video Metadata**
- **Title:** Understanding Go Interfaces: From Theory to Real-World Usage
- **Duration Target:** 40-50 minutes
- **Difficulty:** Intermediate
- **Prerequisites:** Go Basics, Structs, Methods, Pointers

---

## **Video Structure**

### **1. Introduction (3-4 min)**
- Welcome and what viewers will learn
- Why interfaces are Go's "killer feature"
- Comparison to other languages (explicit vs implicit)
- Go philosophy: "Accept interfaces, return structs"
- Preview: Building a plugin system

---

### **2. What Are Interfaces? (5-6 min)**

**Topics to cover:**
- Interfaces define behavior, not data
- Method signatures only
- Implicit implementation
- Any type that has the methods satisfies the interface

**Code Examples:**
```go
package main

import "fmt"

// Interface definition - just method signatures
type Speaker interface {
    Speak() string
}

// Dog implements Speaker (implicitly - no "implements" keyword!)
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

// Cat also implements Speaker
type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

// Robot implements Speaker too
type Robot struct {
    Model string
}

func (r Robot) Speak() string {
    return "Beep boop!"
}

// Function accepts any Speaker
func MakeSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    robot := Robot{Model: "R2D2"}

    // All satisfy Speaker interface
    MakeSpeak(dog)    // Woof!
    MakeSpeak(cat)    // Meow!
    MakeSpeak(robot)  // Beep boop!

    // Can store in slice of interface type
    speakers := []Speaker{dog, cat, robot}
    for _, s := range speakers {
        fmt.Println(s.Speak())
    }
}
```

**Key teaching points:**
- Interface = set of method signatures
- No `implements` keyword - implementation is implicit
- If a type has the methods, it satisfies the interface
- Enables polymorphism without inheritance
- Types don't need to know about interfaces they implement

---

### **3. Interface Satisfaction Rules (5-6 min)**

**Topics to cover:**
- Method sets and receivers
- Pointer vs value receiver implications
- All methods must match exactly

**Code Examples:**
```go
type Writer interface {
    Write(data []byte) error
}

// Value receiver - both Type and *Type satisfy interface
type Buffer struct {
    data []byte
}

func (b Buffer) Write(data []byte) error {
    // Note: can't actually modify b.data here (value receiver)
    return nil
}

var _ Writer = Buffer{}   // OK
var _ Writer = &Buffer{}  // OK

// Pointer receiver - only *Type satisfies interface
type File struct {
    path string
}

func (f *File) Write(data []byte) error {
    // Can modify f here (pointer receiver)
    return nil
}

// var _ Writer = File{}   // ERROR! File doesn't have Write method
var _ Writer = &File{}     // OK

// Rule: If method has pointer receiver, must use pointer to satisfy interface
// Rule: If method has value receiver, both value and pointer work

// Method signature must match EXACTLY
type Processor interface {
    Process(input string) (string, error)
}

type MyProcessor struct{}

// This satisfies Processor:
func (p MyProcessor) Process(input string) (string, error) {
    return input, nil
}

// This does NOT satisfy Processor (different signature):
// func (p MyProcessor) Process(input string) string { ... }
```

**Common mistake:**
```go
type Stringer interface {
    String() string
}

type User struct {
    Name string
}

// Pointer receiver
func (u *User) String() string {
    return u.Name
}

func PrintString(s Stringer) {
    fmt.Println(s.String())
}

func main() {
    user := User{Name: "Alice"}

    // PrintString(user)   // ERROR: User doesn't implement Stringer
    PrintString(&user)     // OK: *User implements Stringer
}
```

**Key teaching points:**
- Value receiver: Type and *Type both satisfy
- Pointer receiver: only *Type satisfies
- All methods must have exact matching signatures
- Use compile-time check: `var _ Interface = Type{}`

---

### **4. The Empty Interface (4-5 min)**

**Topics to cover:**
- `interface{}` and `any` (Go 1.18+)
- Why it accepts any type
- Type assertions
- When to use (and not use) empty interface

**Code Examples:**
```go
// Empty interface - has zero methods
// Every type has at least zero methods, so everything satisfies it
func PrintAnything(v interface{}) {
    fmt.Println(v)
}

// Go 1.18+: 'any' is alias for interface{}
func PrintAny(v any) {
    fmt.Println(v)
}

func main() {
    PrintAnything(42)
    PrintAnything("hello")
    PrintAnything([]int{1, 2, 3})
    PrintAnything(struct{ X int }{X: 10})

    // Type assertion - extract concrete type
    var i interface{} = "hello"

    // Basic assertion (panics if wrong type)
    s := i.(string)
    fmt.Println(s)  // hello

    // Safe assertion with comma-ok
    s, ok := i.(string)
    if ok {
        fmt.Println("It's a string:", s)
    }

    n, ok := i.(int)
    if !ok {
        fmt.Println("Not an int")
    }

    // Type switch - handle multiple types
    describe(42)
    describe("hello")
    describe(true)
    describe([]int{1, 2})
}

func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

**When to use empty interface:**
```go
// GOOD: Generic containers (before Go 1.18 generics)
type Cache struct {
    data map[string]interface{}
}

// GOOD: JSON unmarshaling when structure is unknown
var data interface{}
json.Unmarshal([]byte(`{"key": "value"}`), &data)

// GOOD: Printf-style variadic functions
func Log(format string, args ...interface{}) {
    fmt.Printf(format, args...)
}

// BAD: Avoid when you know the type
func ProcessUser(u interface{}) {  // BAD - just use User type!
    user := u.(User)
    // ...
}

// Since Go 1.18, prefer generics over interface{} where applicable
func First[T any](items []T) T {
    return items[0]
}
```

**Key teaching points:**
- `interface{}` / `any` accepts any type
- Use type assertions to get concrete type back
- Comma-ok pattern prevents panics
- Type switch for multiple type handling
- Prefer specific interfaces over `interface{}`
- Go 1.18 generics often better than `interface{}`

---

### **5. Standard Library Interfaces (6-7 min)**

**Topics to cover:**
- `io.Reader` and `io.Writer`
- `fmt.Stringer`
- `error` interface
- Why small interfaces matter

**Code Examples:**
```go
import (
    "fmt"
    "io"
    "strings"
)

// io.Reader - one of the most important interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Custom type implementing io.Reader
type RepeatReader struct {
    char  byte
    count int
    read  int
}

func (r *RepeatReader) Read(p []byte) (n int, err error) {
    if r.read >= r.count {
        return 0, io.EOF
    }

    toRead := r.count - r.read
    if toRead > len(p) {
        toRead = len(p)
    }

    for i := 0; i < toRead; i++ {
        p[i] = r.char
    }
    r.read += toRead

    return toRead, nil
}

func main() {
    // Our RepeatReader works with any function expecting io.Reader
    reader := &RepeatReader{char: 'A', count: 10}
    data, _ := io.ReadAll(reader)
    fmt.Println(string(data))  // AAAAAAAAAA

    // strings.Reader implements io.Reader
    sr := strings.NewReader("Hello, World!")
    io.Copy(os.Stdout, sr)
}

// fmt.Stringer - custom string representation
type Stringer interface {
    String() string
}

type User struct {
    ID   int
    Name string
}

func (u User) String() string {
    return fmt.Sprintf("User#%d: %s", u.ID, u.Name)
}

func main() {
    user := User{ID: 1, Name: "Alice"}
    fmt.Println(user)  // User#1: Alice (uses String() method)
}

// error interface - extremely simple
type error interface {
    Error() string
}

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

func ValidateUser(name string) error {
    if name == "" {
        return ValidationError{Field: "name", Message: "cannot be empty"}
    }
    return nil
}
```

**Interface composition:**
```go
// Composing interfaces
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// This is how standard library defines io.ReadWriteCloser!
```

**Key teaching points:**
- Small interfaces are powerful (`Reader` has 1 method)
- Composition creates larger interfaces
- `fmt.Stringer` customizes print output
- `error` is just an interface with one method
- Standard library interfaces enable ecosystem interop

---

### **6. Interface Design Principles (5-6 min)**

**Topics to cover:**
- Keep interfaces small
- Accept interfaces, return structs
- Define interfaces at point of use
- Interface segregation

**Code Examples:**
```go
// PRINCIPLE 1: Keep interfaces small

// BAD: Large interface (hard to implement/mock)
type Repository interface {
    Create(user User) error
    Update(user User) error
    Delete(id int) error
    FindByID(id int) (*User, error)
    FindByEmail(email string) (*User, error)
    FindAll() ([]User, error)
    Count() (int, error)
    // ... 10 more methods
}

// GOOD: Small, focused interfaces
type UserCreator interface {
    Create(user User) error
}

type UserFinder interface {
    FindByID(id int) (*User, error)
}

type UserUpdater interface {
    Update(user User) error
}

// Compose when needed
type UserService interface {
    UserCreator
    UserFinder
    UserUpdater
}

// PRINCIPLE 2: Accept interfaces, return structs

// BAD: Returns interface (caller doesn't know concrete type)
func NewUserService() UserServiceInterface {
    return &userService{}
}

// GOOD: Returns concrete type (more flexibility)
func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

// Parameter uses interface (can accept any implementation)
func ProcessUsers(finder UserFinder) {
    // Can work with any UserFinder implementation
}

// PRINCIPLE 3: Define interfaces where they're used

// In package 'handlers':
type UserGetter interface {
    GetUser(id int) (*User, error)
}

type UserHandler struct {
    users UserGetter  // Depends on interface, not concrete type
}

// In package 'repository':
type UserRepository struct {
    db *sql.DB
}

func (r *UserRepository) GetUser(id int) (*User, error) {
    // Implementation
}

// repository.UserRepository satisfies handlers.UserGetter
// without needing to import handlers package!

// PRINCIPLE 4: Interface segregation

// BAD: Force implementations to have unused methods
type Animal interface {
    Walk()
    Swim()
    Fly()
}

// Fish can't walk or fly!
type Fish struct{}
func (f Fish) Walk() { panic("fish can't walk") }  // Forced to implement
func (f Fish) Swim() { /* ... */ }
func (f Fish) Fly() { panic("fish can't fly") }    // Forced to implement

// GOOD: Segregated interfaces
type Walker interface {
    Walk()
}

type Swimmer interface {
    Swim()
}

type Flyer interface {
    Fly()
}

// Types implement only what they can do
type Fish struct{}
func (f Fish) Swim() { /* ... */ }

type Bird struct{}
func (b Bird) Walk() { /* ... */ }
func (b Bird) Fly() { /* ... */ }

type Duck struct{}
func (d Duck) Walk() { /* ... */ }
func (d Duck) Swim() { /* ... */ }
func (d Duck) Fly() { /* ... */ }
```

**Key teaching points:**
- Smaller interfaces = more implementations possible
- Consumer defines the interface they need
- Producer returns concrete types
- Don't force types to implement unused methods
- Implicit implementation enables decoupling

---

### **7. Practical Example: Plugin System (10-12 min)**

**Build together:** An extensible notification system

```go
package main

import (
    "fmt"
    "time"
)

// ========================================
// Core Interface - small and focused
// ========================================

// Notifier is the core interface for all notification types
type Notifier interface {
    Notify(message string) error
    Name() string
}

// ========================================
// Concrete Implementations
// ========================================

// EmailNotifier sends notifications via email
type EmailNotifier struct {
    From    string
    To      string
    SMTPHost string
}

func NewEmailNotifier(from, to, host string) *EmailNotifier {
    return &EmailNotifier{From: from, To: to, SMTPHost: host}
}

func (e *EmailNotifier) Notify(message string) error {
    fmt.Printf("[EMAIL] From: %s, To: %s\n", e.From, e.To)
    fmt.Printf("[EMAIL] Message: %s\n", message)
    // In real code: send via SMTP
    return nil
}

func (e *EmailNotifier) Name() string {
    return "Email"
}

// SlackNotifier sends notifications to Slack
type SlackNotifier struct {
    WebhookURL string
    Channel    string
}

func NewSlackNotifier(webhookURL, channel string) *SlackNotifier {
    return &SlackNotifier{WebhookURL: webhookURL, Channel: channel}
}

func (s *SlackNotifier) Notify(message string) error {
    fmt.Printf("[SLACK] Channel: %s\n", s.Channel)
    fmt.Printf("[SLACK] Message: %s\n", message)
    // In real code: POST to webhook URL
    return nil
}

func (s *SlackNotifier) Name() string {
    return "Slack"
}

// SMSNotifier sends notifications via SMS
type SMSNotifier struct {
    PhoneNumber string
    APIKey      string
}

func NewSMSNotifier(phone, apiKey string) *SMSNotifier {
    return &SMSNotifier{PhoneNumber: phone, APIKey: apiKey}
}

func (s *SMSNotifier) Notify(message string) error {
    fmt.Printf("[SMS] To: %s\n", s.PhoneNumber)
    fmt.Printf("[SMS] Message: %s\n", message)
    // In real code: call SMS API
    return nil
}

func (s *SMSNotifier) Name() string {
    return "SMS"
}

// ConsoleNotifier for development/testing
type ConsoleNotifier struct{}

func NewConsoleNotifier() *ConsoleNotifier {
    return &ConsoleNotifier{}
}

func (c *ConsoleNotifier) Notify(message string) error {
    fmt.Printf("[CONSOLE] %s: %s\n", time.Now().Format("15:04:05"), message)
    return nil
}

func (c *ConsoleNotifier) Name() string {
    return "Console"
}

// ========================================
// Optional Interface - for additional capabilities
// ========================================

// Validator can validate before sending
type Validator interface {
    Validate() error
}

// EmailNotifier implements Validator
func (e *EmailNotifier) Validate() error {
    if e.To == "" {
        return fmt.Errorf("email recipient cannot be empty")
    }
    return nil
}

// ========================================
// Notification Manager
// ========================================

type NotificationManager struct {
    notifiers []Notifier
}

func NewNotificationManager() *NotificationManager {
    return &NotificationManager{
        notifiers: make([]Notifier, 0),
    }
}

// Register adds a notifier (accepts interface)
func (nm *NotificationManager) Register(n Notifier) {
    // Check if notifier implements Validator
    if v, ok := n.(Validator); ok {
        if err := v.Validate(); err != nil {
            fmt.Printf("Warning: %s notifier validation failed: %v\n", n.Name(), err)
            return
        }
    }
    nm.notifiers = append(nm.notifiers, n)
    fmt.Printf("Registered: %s notifier\n", n.Name())
}

// NotifyAll sends message to all registered notifiers
func (nm *NotificationManager) NotifyAll(message string) error {
    fmt.Printf("\n--- Sending notification to %d channels ---\n", len(nm.notifiers))

    var lastErr error
    for _, n := range nm.notifiers {
        if err := n.Notify(message); err != nil {
            fmt.Printf("Error with %s: %v\n", n.Name(), err)
            lastErr = err
        }
    }
    return lastErr
}

// GetNotifier returns notifier by name (type assertion example)
func (nm *NotificationManager) GetNotifier(name string) Notifier {
    for _, n := range nm.notifiers {
        if n.Name() == name {
            return n
        }
    }
    return nil
}

// ========================================
// Decorator Pattern with Interfaces
// ========================================

// LoggingNotifier wraps any Notifier with logging
type LoggingNotifier struct {
    wrapped Notifier
}

func WithLogging(n Notifier) *LoggingNotifier {
    return &LoggingNotifier{wrapped: n}
}

func (l *LoggingNotifier) Notify(message string) error {
    start := time.Now()
    fmt.Printf("[LOG] Starting %s notification...\n", l.wrapped.Name())

    err := l.wrapped.Notify(message)

    elapsed := time.Since(start)
    if err != nil {
        fmt.Printf("[LOG] %s failed after %v: %v\n", l.wrapped.Name(), elapsed, err)
    } else {
        fmt.Printf("[LOG] %s completed in %v\n", l.wrapped.Name(), elapsed)
    }
    return err
}

func (l *LoggingNotifier) Name() string {
    return l.wrapped.Name() + " (logged)"
}

// RetryNotifier wraps any Notifier with retry logic
type RetryNotifier struct {
    wrapped    Notifier
    maxRetries int
}

func WithRetry(n Notifier, maxRetries int) *RetryNotifier {
    return &RetryNotifier{wrapped: n, maxRetries: maxRetries}
}

func (r *RetryNotifier) Notify(message string) error {
    var err error
    for i := 0; i <= r.maxRetries; i++ {
        err = r.wrapped.Notify(message)
        if err == nil {
            return nil
        }
        fmt.Printf("[RETRY] Attempt %d/%d failed for %s\n", i+1, r.maxRetries+1, r.wrapped.Name())
    }
    return err
}

func (r *RetryNotifier) Name() string {
    return r.wrapped.Name()
}

// ========================================
// Main - Demonstration
// ========================================

func main() {
    fmt.Println("=== Notification Plugin System ===\n")

    // Create manager
    manager := NewNotificationManager()

    // Register various notifiers
    manager.Register(NewConsoleNotifier())
    manager.Register(NewEmailNotifier("system@company.com", "admin@company.com", "smtp.company.com"))
    manager.Register(NewSlackNotifier("https://hooks.slack.com/...", "#alerts"))
    manager.Register(NewSMSNotifier("+1234567890", "api-key-123"))

    // This one will fail validation
    manager.Register(NewEmailNotifier("system@company.com", "", "smtp.company.com"))

    // Send to all
    manager.NotifyAll("Server CPU usage exceeded 90%!")

    fmt.Println("\n=== Using Decorators ===\n")

    // Wrap with logging
    loggedEmail := WithLogging(NewEmailNotifier("alerts@company.com", "ops@company.com", "smtp.company.com"))
    loggedEmail.Notify("Database connection lost")

    // Wrap with retry and logging
    retryLoggedSlack := WithLogging(
        WithRetry(
            NewSlackNotifier("https://hooks.slack.com/...", "#critical"),
            3,
        ),
    )
    retryLoggedSlack.Notify("Critical: Payment service down!")

    fmt.Println("\n=== Type Assertions ===\n")

    // Get specific notifier and use type-specific methods
    notifier := manager.GetNotifier("Email")
    if email, ok := notifier.(*EmailNotifier); ok {
        fmt.Printf("Email notifier found: sending to %s\n", email.To)
    }

    // Check if notifier implements optional interface
    if validator, ok := notifier.(Validator); ok {
        if err := validator.Validate(); err != nil {
            fmt.Println("Validation error:", err)
        } else {
            fmt.Println("Notifier is valid")
        }
    }
}
```

**Walk through:**
- Core `Notifier` interface (small, focused)
- Multiple implementations (Email, Slack, SMS, Console)
- Optional `Validator` interface
- Type assertions to check capabilities
- Decorator pattern (logging, retry wrappers)
- Manager works with any Notifier
- Easy to add new notification types

---

### **8. Testing with Interfaces (5-6 min)**

**Topics to cover:**
- Mock implementations
- Interface-based testing
- Dependency injection for testability

**Code Examples:**
```go
// Production code
type UserRepository interface {
    GetByID(id int) (*User, error)
    Save(user *User) error
}

type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*User, error) {
    return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(name, email string) (*User, error) {
    user := &User{Name: name, Email: email}
    if err := s.repo.Save(user); err != nil {
        return nil, err
    }
    return user, nil
}

// Test code - mock implementation
type MockUserRepository struct {
    users map[int]*User
    saveError error
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users: make(map[int]*User),
    }
}

func (m *MockUserRepository) GetByID(id int) (*User, error) {
    user, ok := m.users[id]
    if !ok {
        return nil, fmt.Errorf("user not found: %d", id)
    }
    return user, nil
}

func (m *MockUserRepository) Save(user *User) error {
    if m.saveError != nil {
        return m.saveError
    }
    m.users[user.ID] = user
    return nil
}

// Helper to set up error scenarios
func (m *MockUserRepository) SetSaveError(err error) {
    m.saveError = err
}

// Tests
func TestUserService_GetUser(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    mockRepo.users[1] = &User{ID: 1, Name: "Alice", Email: "alice@test.com"}

    service := NewUserService(mockRepo)

    // Act
    user, err := service.GetUser(1)

    // Assert
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }
    if user.Name != "Alice" {
        t.Errorf("expected Alice, got %s", user.Name)
    }
}

func TestUserService_GetUser_NotFound(t *testing.T) {
    mockRepo := NewMockUserRepository()
    service := NewUserService(mockRepo)

    _, err := service.GetUser(999)

    if err == nil {
        t.Error("expected error for non-existent user")
    }
}

func TestUserService_CreateUser_SaveError(t *testing.T) {
    mockRepo := NewMockUserRepository()
    mockRepo.SetSaveError(fmt.Errorf("database error"))
    service := NewUserService(mockRepo)

    _, err := service.CreateUser("Bob", "bob@test.com")

    if err == nil {
        t.Error("expected error when save fails")
    }
}
```

**Key teaching points:**
- Interfaces enable mock implementations
- Test behavior, not implementation
- Mock can simulate error conditions
- No need for mocking frameworks
- Production code depends on interfaces
- Tests inject mock implementations

---

### **9. Common Interface Mistakes (4-5 min)**

**Topics to cover:**
- Interface pollution
- Premature abstraction
- Type assertion abuse

**Code Examples:**
```go
// MISTAKE 1: Interface pollution (too many interfaces)

// BAD: Every type has its own interface
type UserInterface interface {
    GetName() string
}

type User struct {
    Name string
}

func (u User) GetName() string {
    return u.Name
}

// Just use the concrete type!
func ProcessUser(u User) {  // Not UserInterface
    fmt.Println(u.Name)
}

// MISTAKE 2: Premature abstraction

// BAD: Creating interface before you need it
type Logger interface {
    Log(msg string)
}

// You only have one implementation!
type FileLogger struct{}

func (f FileLogger) Log(msg string) { /* ... */ }

// Don't create interface until you need multiple implementations
// or need to mock for testing

// MISTAKE 3: Returning interface instead of concrete type

// BAD: Returns interface
func NewService() ServiceInterface {
    return &service{}
}

// GOOD: Returns concrete type
func NewService() *Service {
    return &Service{}
}

// MISTAKE 4: Accepting concrete when interface would work

// BAD: Only works with this specific type
func ProcessFile(f *os.File) error {
    // reads from file
}

// GOOD: Works with any reader
func ProcessReader(r io.Reader) error {
    // reads from any source
}

// MISTAKE 5: Interface{} abuse

// BAD: Using interface{} when type is known
func BadProcess(data interface{}) {
    user := data.(User)  // Why not just accept User?
    // ...
}

// GOOD: Use actual type
func GoodProcess(user User) {
    // ...
}

// MISTAKE 6: Large interfaces

// BAD: Forces implementations to have many methods
type Repository interface {
    Create(v interface{}) error
    Read(id int) (interface{}, error)
    Update(v interface{}) error
    Delete(id int) error
    List() ([]interface{}, error)
    Count() (int, error)
    Search(query string) ([]interface{}, error)
    // ... more methods
}

// GOOD: Small, focused interfaces
type Creator interface {
    Create(v interface{}) error
}

type Reader interface {
    Read(id int) (interface{}, error)
}

// Compose when needed
type ReadWriter interface {
    Reader
    Creator
}
```

**Key teaching points:**
- Don't create interfaces for single implementations
- Wait until you need the abstraction
- Return concrete types, accept interfaces
- Keep interfaces small
- Avoid `interface{}` when you know the type
- Don't wrap every type in an interface

---

### **10. Next Steps & Wrap-up (2-3 min)**

**Recap what was covered:**
- Interface basics and implicit implementation
- Interface satisfaction rules (pointer vs value)
- Empty interface and type assertions
- Standard library interfaces
- Interface design principles
- Plugin system example
- Testing with interfaces
- Common mistakes

**Preview next topics:**
- Error handling patterns
- Concurrency (interfaces with goroutines)
- Generics (Go 1.18+)

**Homework/Practice suggestions:**
1. **Easy:** Implement `fmt.Stringer` for a custom type
2. **Medium:** Create a `Shape` interface with `Area()` and `Perimeter()`
3. **Challenge:** Build a cache with pluggable storage backends
4. **Advanced:** Implement a middleware chain using interfaces

**Resources:**
- Effective Go on Interfaces
- Go Blog: "Go Data Structures: Interfaces"
- Your GitHub repo with plugin system code

---

## **Production Notes**

### **Screen Setup:**
- Code editor: 70% of screen
- Terminal output: 30% of screen
- Split view for interface definition vs implementation
- Font size: 18-20pt minimum

### **Teaching Techniques:**
- Start with concrete problem, then show interface solution
- Show multiple implementations side by side
- Demonstrate implicit implementation (no `implements`)
- Show compile errors when interface not satisfied
- Use `var _ Interface = Type{}` for compile checks

### **Visual Aids:**
- Diagram: Interface as "contract" between types
- Diagram: Multiple types implementing same interface
- Diagram: Interface composition
- Animation: Type assertion flow

### **Engagement:**
- "Does this type satisfy the interface?" quizzes
- "What's wrong with this interface design?" exercises
- Live coding: add new notifier type
- Compare to Java/C# explicit interfaces

---

## **Supplementary Materials to Provide**

1. **GitHub Repository:**
   - Complete notification system
   - All code examples from video
   - Practice exercises with solutions
   - README with interface best practices

2. **Cheat Sheet (PDF/Gist):**
   ```
   Interface definition:     type Name interface { Methods }
   Implicit implementation:  Just implement the methods
   Type assertion:           value.(Type)
   Safe type assertion:      v, ok := value.(Type)
   Type switch:              switch v := value.(type) { }
   Empty interface:          interface{} or any
   Compose interfaces:       type RW interface { Reader; Writer }
   Compile check:            var _ Interface = (*Type)(nil)
   ```

3. **Practice Exercises:**
   - **Easy:** Implement Stringer for Point{X, Y}
   - **Medium:** Payment processor interface (Credit, PayPal, Crypto)
   - **Challenge:** HTTP middleware system
   - **Advanced:** Plugin loader with interface discovery

4. **Design Guidelines PDF:**
   - When to create an interface
   - Interface naming conventions
   - Small interface examples from stdlib
   - Anti-patterns to avoid

---

This tutorial covers interfaces comprehensively while focusing on practical usage. The plugin system example shows real-world application, and the testing section demonstrates why interfaces matter for code quality.
