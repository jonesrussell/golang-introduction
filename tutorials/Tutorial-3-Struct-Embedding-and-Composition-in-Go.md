## **Video Tutorial Plan: Struct Embedding and Composition in Go**

### **Video Metadata**
- **Title:** Struct Embedding and Composition in Go
- **Duration Target:** 35-45 minutes
- **Difficulty:** Intermediate
- **Prerequisites:** Go Structs (definition, initialization, methods)

---

## **Video Structure**

### **1. Introduction (3-4 min)**
- Welcome and what viewers will learn
- Why composition over inheritance?
- Go's philosophy: "favor composition over inheritance"
- Comparison: How other languages do it (Java/Python classes vs Go structs)
- Show the final example: Building an employee management system
- Preview: How embedding provides code reuse without inheritance complexity

---

### **2. The Problem: Code Reuse Without Inheritance (4-5 min)**

**Topics to cover:**
- Traditional OOP inheritance (what Go doesn't have)
- The diamond problem and why inheritance is complex
- Go's alternative: composition and embedding

**Code Example - The Problem:**
```go
// Imagine we want to model different types of users
// In traditional OOP (pseudocode):
// class User {
//     name, email
// }
// class Admin extends User {
//     permissions
// }
// class Customer extends User {
//     orderHistory
// }

// Without embedding, we'd repeat fields:
type Admin struct {
    Name        string  // Repeated
    Email       string  // Repeated
    Permissions []string
}

type Customer struct {
    Name         string  // Repeated
    Email        string  // Repeated
    OrderHistory []Order
}

// This violates DRY principle and is hard to maintain
// What if we want to add a PhoneNumber field to all users?
```

**Key teaching points:**
- Go has no class inheritance
- Go has no extends or super keywords
- This is intentional - inheritance creates tight coupling
- Go provides composition instead
- Embedding is Go's answer to code reuse

---

### **3. Basic Struct Composition (5-6 min)**

**Topics to cover:**
- Composition via explicit fields
- Has-a relationships
- Accessing nested fields
- When explicit composition is appropriate

**Code Examples:**
```go
// Basic composition - one struct contains another

type Address struct {
    Street  string
    City    string
    State   string
    ZipCode string
}

type Person struct {
    FirstName string
    LastName  string
    Address   Address  // Explicit field - composition
}

// Usage - explicit field access
person := Person{
    FirstName: "John",
    LastName:  "Doe",
    Address: Address{
        Street:  "123 Main St",
        City:    "Springfield",
        State:   "IL",
        ZipCode: "62701",
    },
}

// Accessing nested fields
fmt.Println(person.Address.Street)  // Must go through Address field
fmt.Println(person.Address.City)

// Methods on composed struct
func (a Address) FullAddress() string {
    return fmt.Sprintf("%s, %s, %s %s", 
        a.Street, a.City, a.State, a.ZipCode)
}

// Must access through field name
fmt.Println(person.Address.FullAddress())

// Another example - explicit composition
type Engine struct {
    Horsepower int
    Type       string
}

type Car struct {
    Brand  string
    Model  string
    Engine Engine  // Car HAS-AN Engine
}

car := Car{
    Brand: "Toyota",
    Model: "Camry",
    Engine: Engine{
        Horsepower: 203,
        Type:       "V6",
    },
}

fmt.Printf("%s %s has %d HP\n", 
    car.Brand, car.Model, car.Engine.Horsepower)
```

**Key teaching points:**
- Regular composition uses named fields
- Represents clear "has-a" relationships
- Must explicitly reference the field name
- Good when the relationship is explicit (Car has Engine)
- Fields and methods are accessed through the field name

---

### **4. Struct Embedding - The Basics (7-8 min)**

**Topics to cover:**
- Embedding syntax (anonymous fields)
- Field promotion
- Method promotion
- Difference between embedding and composition

**Code Examples:**
```go
// Embedding - anonymous field (no field name)

type User struct {
    ID       int
    Username string
    Email    string
}

func (u User) GetDisplayName() string {
    return fmt.Sprintf("@%s", u.Username)
}

func (u User) SendEmail(subject string) {
    fmt.Printf("Sending '%s' to %s\n", subject, u.Email)
}

// Admin embeds User
type Admin struct {
    User              // Embedded struct - NO field name
    Permissions []string
}

// Usage - field promotion
admin := Admin{
    User: User{
        ID:       1,
        Username: "admin",
        Email:    "admin@example.com",
    },
    Permissions: []string{"read", "write", "delete"},
}

// Can access User fields directly (promoted)
fmt.Println(admin.Username)  // Not admin.User.Username
fmt.Println(admin.Email)     // Not admin.User.Email

// Can also access through type name
fmt.Println(admin.User.Username)  // Still works

// Method promotion - User methods available on Admin
fmt.Println(admin.GetDisplayName())  // Promoted method
admin.SendEmail("Welcome")           // Promoted method

// Admin-specific fields
fmt.Println(admin.Permissions)
```

**More examples:**
```go
type Customer struct {
    User         // Embedded
    OrderCount   int
    LoyaltyPoints int
}

customer := Customer{
    User: User{
        ID:       2,
        Username: "john_doe",
        Email:    "john@example.com",
    },
    OrderCount:   15,
    LoyaltyPoints: 150,
}

// All User fields and methods promoted
fmt.Println(customer.Username)        // Promoted field
fmt.Println(customer.GetDisplayName()) // Promoted method

// Customer-specific functionality
func (c Customer) GetTier() string {
    if c.OrderCount > 10 {
        return "Gold"
    }
    return "Silver"
}

fmt.Println(customer.GetTier())  // Customer's own method
```

**Key teaching points:**
- Embedding = anonymous field (type without name)
- Embedded fields are "promoted" to outer struct
- Can access embedded fields/methods directly
- Can still access through type name if needed
- Looks like inheritance but it's composition
- The embedded struct doesn't know it's embedded

---

### **5. Multiple Embedding (6-7 min)**

**Topics to cover:**
- Embedding multiple structs
- Field/method name conflicts
- Resolution order
- Explicitly accessing embedded types

**Code Examples:**
```go
type Timestamps struct {
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (t Timestamps) Age() time.Duration {
    return time.Since(t.CreatedAt)
}

type Metadata struct {
    Tags        []string
    Description string
}

func (m Metadata) HasTag(tag string) bool {
    for _, t := range m.Tags {
        if t == tag {
            return true
        }
    }
    return false
}

// Embed multiple structs
type Article struct {
    Timestamps  // Embedded
    Metadata    // Embedded
    Title       string
    Content     string
    Author      string
}

// Usage
article := Article{
    Timestamps: Timestamps{
        CreatedAt: time.Now().Add(-24 * time.Hour),
        UpdatedAt: time.Now(),
    },
    Metadata: Metadata{
        Tags:        []string{"golang", "tutorial"},
        Description: "Learn about struct embedding",
    },
    Title:   "Go Structs",
    Content: "...",
    Author:  "Russell",
}

// Can access all promoted fields
fmt.Println(article.CreatedAt)   // From Timestamps
fmt.Println(article.Tags)        // From Metadata
fmt.Println(article.Title)       // From Article

// Can call all promoted methods
fmt.Println(article.Age())              // From Timestamps
fmt.Println(article.HasTag("golang"))   // From Metadata

// Handling conflicts - same field/method name in multiple embedded structs
type A struct {
    Name string
}

func (a A) Display() {
    fmt.Println("From A:", a.Name)
}

type B struct {
    Name string
}

func (b B) Display() {
    fmt.Println("From B:", b.Name)
}

type C struct {
    A  // Both have Name field and Display method
    B
}

c := C{
    A: A{Name: "A's name"},
    B: B{Name: "B's name"},
}

// Ambiguous - won't compile:
// fmt.Println(c.Name)     // ERROR: ambiguous selector c.Name
// c.Display()             // ERROR: ambiguous selector c.Display

// Must be explicit when there's a conflict:
fmt.Println(c.A.Name)  // OK - explicit
fmt.Println(c.B.Name)  // OK - explicit
c.A.Display()          // OK - explicit
c.B.Display()          // OK - explicit

// No conflict if outer struct has the same field
type D struct {
    A
    B
    Name string  // Shadows embedded Name fields
}

d := D{
    A:    A{Name: "A's name"},
    B:    B{Name: "B's name"},
    Name: "D's name",
}

fmt.Println(d.Name)    // D's name (outer struct wins)
fmt.Println(d.A.Name)  // A's name (still accessible)
fmt.Println(d.B.Name)  // B's name (still accessible)
```

**Key teaching points:**
- Can embed multiple structs
- All fields/methods from all embedded types are promoted
- Name conflicts cause compile errors (not runtime!)
- Must explicitly specify which embedded type when ambiguous
- Outer struct fields shadow embedded fields with same name
- This is safer than inheritance (explicit is better than implicit)

---

### **6. Embedding and Interfaces (7-8 min)**

**Topics to cover:**
- How embedded structs satisfy interfaces
- Interface composition via embedding
- Wrapper pattern with embedding
- Overriding embedded methods

**Code Examples:**
```go
// Interface satisfaction through embedding

type Notifier interface {
    Notify(message string)
}

type EmailNotifier struct {
    Email string
}

func (e EmailNotifier) Notify(message string) {
    fmt.Printf("Email to %s: %s\n", e.Email, message)
}

// User embeds EmailNotifier and satisfies Notifier interface
type User struct {
    Name string
    EmailNotifier  // Embedded - promotes Notify method
}

// User automatically satisfies Notifier interface
func SendNotification(n Notifier, msg string) {
    n.Notify(msg)
}

user := User{
    Name: "Alice",
    EmailNotifier: EmailNotifier{
        Email: "alice@example.com",
    },
}

SendNotification(user, "Welcome!")  // Works! User satisfies Notifier

// Overriding embedded methods
type AdminUser struct {
    User
    Permissions []string
}

// Override the Notify method
func (a AdminUser) Notify(message string) {
    fmt.Println("[ADMIN NOTIFICATION]")
    a.User.Notify(message)  // Can still call embedded version
    fmt.Printf("Admin %s has been notified\n", a.Name)
}

admin := AdminUser{
    User: User{
        Name: "Bob",
        EmailNotifier: EmailNotifier{
            Email: "bob@example.com",
        },
    },
    Permissions: []string{"admin"},
}

SendNotification(admin, "System alert")  // Uses overridden method

// Interface embedding
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Compose interfaces by embedding
type ReadWriter interface {
    Reader  // Embedded interface
    Writer  // Embedded interface
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// This is how io.ReadWriteCloser is defined in standard library!

// Any type that has Read, Write, and Close methods satisfies ReadWriteCloser
type File struct {
    name string
    data []byte
}

func (f *File) Read(p []byte) (int, error) {
    // Implementation
    return 0, nil
}

func (f *File) Write(p []byte) (int, error) {
    // Implementation
    return 0, nil
}

func (f *File) Close() error {
    // Implementation
    return nil
}

var rwc ReadWriteCloser = &File{name: "test.txt"}
```

**Key teaching points:**
- Embedded struct methods count toward interface satisfaction
- Can override embedded methods by defining same method on outer struct
- Can still call original embedded method explicitly
- Interface embedding creates composite interfaces
- This is how standard library composes interfaces (io.Reader, io.Writer, etc.)
- Embedding provides delegation pattern

---

### **7. Practical Patterns with Embedding (8-10 min)**

**Topics to cover:**
- Mixin pattern
- Decorator pattern
- Base functionality pattern
- When to use embedding vs composition

**Code Example 1: Mixin Pattern**
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

**Code Example 2: Decorator Pattern**
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

**Code Example 3: Base Entity Pattern**
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

**Key teaching points:**
- Mixin pattern: Embed common functionality across multiple types
- Decorator pattern: Add behavior by wrapping
- Base entity: Share common behavior across related types
- DRY: Don't repeat timestamp, audit, soft-delete logic
- Embedding promotes code reuse without inheritance

---

### **8. When to Use Embedding vs Composition (5-6 min)**

**Decision guide:**

```go
// Use EMBEDDING when:
// 1. You want promoted fields/methods (convenience)
// 2. You're implementing mixin behavior
// 3. The embedded type is a "base" or "common" functionality
// 4. You want the outer type to satisfy interfaces of embedded type

type User struct {
    Auditable  // ✅ Good - mixin behavior
    ID         int
    Name       string
}

// Use EXPLICIT COMPOSITION when:
// 1. The relationship is clearly "has-a"
// 2. You want explicit access (clarity over convenience)
// 3. You might have multiple of the same type
// 4. The field has semantic meaning

type Car struct {
    Engine Engine  // ✅ Good - Car HAS-AN Engine
    Brand  string
}

type House struct {
    PrimaryBathroom Bathroom  // ✅ Good - semantic meaning
    Bedrooms        []Bedroom // Multiple bedrooms
}

// ANTI-PATTERNS to avoid:

// ❌ BAD: Embedding just to avoid typing
type Person struct {
    Address  // BAD - this should be explicit field
    Name string
}
// Problem: person.Street is confusing - is it Person's street?
// Better: person.Address.Street - clear relationship

// ❌ BAD: Embedding unrelated types
type Config struct {
    sync.Mutex  // BAD - Config is not a specialized Mutex
    Settings map[string]string
}
// Problem: Exposes Lock/Unlock at wrong level
// Better: Have a mutex as private field

// ✅ GOOD: Embedding mutex in correct context
type SafeCounter struct {
    sync.Mutex  // OK - SafeCounter IS synchronized
    count int
}

func (sc *SafeCounter) Increment() {
    sc.Lock()
    defer sc.Unlock()
    sc.count++
}

// ❌ BAD: Embedding for code reuse when composition is clearer
type AdminUser struct {
    User  // Confusing - is AdminUser a specialized User?
    AdminLevel int
}
// Better: explicit composition or separate types

// ✅ GOOD: Clear "is-a-specialized-version" relationship
type Buffer struct {
    bytes.Buffer  // OK - adding functionality to Buffer
}

func (b *Buffer) WriteJSON(v interface{}) error {
    data, err := json.Marshal(v)
    if err != nil {
        return err
    }
    b.Write(data)
    return nil
}
```

**Decision flowchart:**
```
┌─────────────────────────────────────────┐
│ Does the outer type need ALL behaviors  │
│ of the inner type?                      │
└────────────┬────────────────────────────┘
             │
         No ─┼─ Yes
             │       │
             │       ▼
             │   ┌───────────────────────────┐
             │   │ Is it a "base" or "mixin" │
             │   │ pattern?                  │
             │   └─────────┬─────────────────┘
             │             │
             │         Yes─┼─ No
             │             │    │
             ▼             ▼    ▼
     Use explicit    Use      Consider
     composition   embedding  explicit
                             composition
```

**Key teaching points:**
- Embedding is powerful but can be overused
- Prefer explicit composition for clarity
- Use embedding for mixins and base patterns
- Don't embed just to save typing
- Think about the conceptual relationship
- "Has-a" → composition; "Is-a-kind-of" → maybe embedding

---

### **9. Complete Practical Example (10-12 min)**

**Build together:** Employee Management System

```go
package main

import (
    "fmt"
    "time"
)

// ========================================
// Mixins - Reusable functionality
// ========================================

type Timestamps struct {
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (t *Timestamps) Touch() {
    t.UpdatedAt = time.Now()
    if t.CreatedAt.IsZero() {
        t.CreatedAt = t.UpdatedAt
    }
}

type Identifiable struct {
    ID int
}

// ========================================
// Base Person type
// ========================================

type Person struct {
    Timestamps   // Mixin
    Identifiable // Mixin
    FirstName    string
    LastName     string
    Email        string
    Phone        string
}

func (p Person) FullName() string {
    return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p Person) ContactInfo() string {
    return fmt.Sprintf("%s (%s)", p.Email, p.Phone)
}

// ========================================
// Employee types with different roles
// ========================================

type Employee struct {
    Person           // Embedded - promotes all Person fields/methods
    EmployeeID       string
    Department       string
    HireDate         time.Time
    Salary           float64
}

func (e Employee) YearsOfService() int {
    return int(time.Since(e.HireDate).Hours() / 24 / 365)
}

func (e Employee) String() string {
    return fmt.Sprintf("Employee[%s]: %s - %s", 
        e.EmployeeID, e.FullName(), e.Department)
}

type Manager struct {
    Employee            // Embedded - Manager IS an Employee
    TeamSize     int
    DirectReports []string
}

func (m Manager) CanApprove(amount float64) bool {
    // Managers can approve up to $10k per team member
    return amount <= float64(m.TeamSize) * 10000
}

// Override String method
func (m Manager) String() string {
    return fmt.Sprintf("Manager[%s]: %s - %s (Team: %d)", 
        m.EmployeeID, m.FullName(), m.Department, m.TeamSize)
}

type Developer struct {
    Employee              // Embedded
    ProgrammingLanguages []string
    SeniorityLevel      string
}

func (d Developer) HasSkill(language string) bool {
    for _, lang := range d.ProgrammingLanguages {
        if lang == language {
            return true
        }
    }
    return false
}

func (d Developer) String() string {
    return fmt.Sprintf("Developer[%s]: %s - %s %s", 
        d.EmployeeID, d.FullName(), d.SeniorityLevel, d.Department)
}

// ========================================
// Contractor - different from Employee
// ========================================

type Contractor struct {
    Person              // Embedded Person (not Employee!)
    ContractID   string
    HourlyRate   float64
    EndDate      time.Time
}

func (c Contractor) IsActive() bool {
    return time.Now().Before(c.EndDate)
}

func (c Contractor) String() string {
    status := "Active"
    if !c.IsActive() {
        status = "Expired"
    }
    return fmt.Sprintf("Contractor[%s]: %s - %s", 
        c.ContractID, c.FullName(), status)
}

// ========================================
// Interfaces that work with embedded types
// ========================================

type Worker interface {
    FullName() string
    String() string
}

type Payable interface {
    CalculatePay() float64
}

func (e Employee) CalculatePay() float64 {
    return e.Salary / 12  // Monthly salary
}

func (c Contractor) CalculatePay() float64 {
    return c.HourlyRate * 160  // Assume 160 hours/month
}

// ========================================
// Department management
// ========================================

type Department struct {
    Timestamps
    Name      string
    Budget    float64
    Workers   []Worker  // Can hold any Worker type
}

func (d *Department) AddWorker(w Worker) {
    d.Workers = append(d.Workers, w)
    d.Touch()
    fmt.Printf("Added %s to %s department\n", w.FullName(), d.Name)
}

func (d Department) ListWorkers() {
    fmt.Printf("\n%s Department (%d workers):\n", d.Name, len(d.Workers))
    for i, worker := range d.Workers {
        fmt.Printf("  %d. %s\n", i+1, worker.String())
    }
}

func (d Department) TotalPayroll() float64 {
    total := 0.0
    for _, worker := range d.Workers {
        if p, ok := worker.(Payable); ok {
            total += p.CalculatePay()
        }
    }
    return total
}

// ========================================
// Main - Demonstration
// ========================================

func main() {
    fmt.Println("=== Employee Management System ===\n")

    // Create Engineering Department
    engineering := &Department{
        Name:   "Engineering",
        Budget: 500000,
    }

    // Create a Manager
    manager := Manager{
        Employee: Employee{
            Person: Person{
                Identifiable: Identifiable{ID: 1},
                FirstName:    "Alice",
                LastName:     "Johnson",
                Email:        "alice@company.com",
                Phone:        "555-0101",
            },
            EmployeeID: "E001",
            Department: "Engineering",
            HireDate:   time.Now().AddDate(-3, 0, 0),
            Salary:     120000,
        },
        TeamSize:      5,
        DirectReports: []string{"E002", "E003", "E004"},
    }
    manager.Touch()

    // Create Developers
    dev1 := Developer{
        Employee: Employee{
            Person: Person{
                Identifiable: Identifiable{ID: 2},
                FirstName:    "Bob",
                LastName:     "Smith",
                Email:        "bob@company.com",
                Phone:        "555-0102",
            },
            EmployeeID: "E002",
            Department: "Engineering",
            HireDate:   time.Now().AddDate(-2, 0, 0),
            Salary:     95000,
        },
        ProgrammingLanguages: []string{"Go", "Python", "JavaScript"},
        SeniorityLevel:      "Senior",
    }
    dev1.Touch()

    dev2 := Developer{
        Employee: Employee{
            Person: Person{
                Identifiable: Identifiable{ID: 3},
                FirstName:    "Charlie",
                LastName:     "Brown",
                Email:        "charlie@company.com",
                Phone:        "555-0103",
            },
            EmployeeID: "E003",
            Department: "Engineering",
            HireDate:   time.Now().AddDate(-1, 0, 0),
            Salary:     85000,
        },
        ProgrammingLanguages: []string{"Go", "Rust"},
        SeniorityLevel:      "Mid-level",
    }
    dev2.Touch()

    // Create a Contractor
    contractor := Contractor{
        Person: Person{
            Identifiable: Identifiable{ID: 4},
            FirstName:    "Diana",
            LastName:     "Williams",
            Email:        "diana@contractor.com",
            Phone:        "555-0104",
        },
        ContractID: "C001",
        HourlyRate: 85,
        EndDate:    time.Now().AddDate(0, 6, 0),
    }
    contractor.Touch()

    // Add workers to department
    engineering.AddWorker(manager)
    engineering.AddWorker(dev1)
    engineering.AddWorker(dev2)
    engineering.AddWorker(contractor)

    // List all workers
    engineering.ListWorkers()

    // Calculate payroll
    fmt.Printf("\nTotal monthly payroll: $%.2f\n", 
        engineering.TotalPayroll())

    // Demonstrate promoted methods
    fmt.Printf("\n=== Promoted Method Access ===\n")
    fmt.Printf("Manager full name: %s\n", manager.FullName())  // From Person
    fmt.Printf("Manager contact: %s\n", manager.ContactInfo()) // From Person
    fmt.Printf("Manager ID: %d\n", manager.ID)                 // From Identifiable
    fmt.Printf("Years of service: %d\n", manager.YearsOfService()) // From Employee

    // Demonstrate method override
    fmt.Printf("\n=== Method Override ===\n")
    fmt.Println(manager.String())  // Manager's custom String
    fmt.Println(dev1.String())     // Developer's custom String

    // Demonstrate type-specific functionality
    fmt.Printf("\n=== Type-Specific Methods ===\n")
    if manager.CanApprove(40000) {
        fmt.Printf("%s can approve $40,000 expense\n", manager.FullName())
    }

    if dev1.HasSkill("Go") {
        fmt.Printf("%s knows Go!\n", dev1.FullName())
    }

    if contractor.IsActive() {
        fmt.Printf("%s contract is active until %s\n", 
            contractor.FullName(), 
            contractor.EndDate.Format("2006-01-02"))
    }

    // Demonstrate timestamp tracking
    fmt.Printf("\n=== Timestamp Tracking ===\n")
    fmt.Printf("Manager created: %s\n", 
        manager.CreatedAt.Format("2006-01-02 15:04:05"))
    fmt.Printf("Dev1 last updated: %s\n", 
        dev1.UpdatedAt.Format("2006-01-02 15:04:05"))
}
```

**Walk through:**
- Mixin structs for common functionality (Timestamps, Identifiable)
- Base Person struct with common identity fields
- Employee embeds Person (promotes all fields/methods)
- Manager and Developer embed Employee (multi-level embedding)
- Contractor embeds Person directly (different hierarchy)
- Interfaces work with embedded types
- Method promotion in action
- Method overriding
- Type assertions for Payable interface
- Real-world modeling with embedding

**Output explanation:**
```
=== Employee Management System ===

Added Alice Johnson to Engineering department
Added Bob Smith to Engineering department
Added Charlie Brown to Engineering department
Added Diana Williams to Engineering department

Engineering Department (4 workers):
  1. Manager[E001]: Alice Johnson - Engineering (Team: 5)
  2. Developer[E002]: Bob Smith - Senior Engineering
  3. Developer[E003]: Charlie Brown - Mid-level Engineering
  4. Contractor[C001]: Diana Williams - Active

Total monthly payroll: $33833.33

=== Promoted Method Access ===
Manager full name: Alice Johnson
Manager contact: alice@company.com (555-0101)
Manager ID: 1
Years of service: 3

=== Method Override ===
Manager[E001]: Alice Johnson - Engineering (Team: 5)
Developer[E002]: Bob Smith - Senior Engineering

=== Type-Specific Methods ===
Alice Johnson can approve $40,000 expense
Bob Smith knows Go!
Diana Williams contract is active until 2025-06-17

=== Timestamp Tracking ===
Manager created: 2024-12-17 10:30:15
Dev1 last updated: 2024-12-17 10:30:15
```

---

### **10. Common Pitfalls & Best Practices (5-6 min)**

**Cover these important mistakes:**

```go
// ❌ PITFALL 1: Embedding to avoid typing (wrong motivation)
type User struct {
    Address  // BAD - unclear semantic relationship
    Name string
}

// ✅ BETTER: Explicit composition for clarity
type User struct {
    Name    string
    Address Address  // Clear: User has an Address
}

// ❌ PITFALL 2: Over-embedding creates confusion
type Admin struct {
    User
    Manager
    DatabaseConnection  // Too many embedded types!
    Logger
}
// Hard to understand what Admin IS

// ✅ BETTER: Selective embedding + composition
type Admin struct {
    User                 // Admin IS a specialized User
    ManagedTeam   []User // Composition for clarity
    db            *sql.DB  // Private, not promoted
    logger        Logger   // Private, not promoted
}

// ❌ PITFALL 3: Embedding interfaces when not needed
type MyService struct {
    io.Reader  // Why? Just for Read method?
    config Config
}

// ✅ BETTER: Explicit field if you need it
type MyService struct {
    reader io.Reader  // Clear: service uses a reader
    config Config
}

// ❌ PITFALL 4: Circular embedding
type A struct {
    B  // A embeds B
}

type B struct {
    A  // B embeds A - CIRCULAR!
}
// Won't compile - infinite size

// ❌ PITFALL 5: Embedding with name conflicts (unintentional)
type Author struct {
    Name string
}

type Book struct {
    Author  // Embeds Author
    Name string  // Conflicts! Which Name?
}

book := Book{
    Author: Author{Name: "Tolkien"},
    Name:   "The Hobbit",
}

// book.Name is the Book's Name (outer wins)
// book.Author.Name is the Author's Name

// ✅ BETTER: Use different field names or explicit composition
type Book struct {
    Title  string
    Author Author  // Explicit - no confusion
}

// ❌ PITFALL 6: Exposing internals through embedding
type Cache struct {
    sync.Mutex  // BAD - exposes Lock/Unlock publicly
    data map[string]interface{}
}

// Now users can do: cache.Lock() - should be internal!

// ✅ BETTER: Private field for sync
type Cache struct {
    mu   sync.Mutex  // Private
    data map[string]interface{}
}

func (c *Cache) Get(key string) interface{} {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.data[key]
}

// ✅ EXCEPTION: sync.Mutex embedding is OK for types that ARE synchronization primitives
type SafeCounter struct {
    sync.Mutex  // OK - SafeCounter IS synchronized
    count int
}

// ❌ PITFALL 7: Forgetting embedded pointer vs value semantics
type Base struct {
    value int
}

func (b *Base) SetValue(v int) {
    b.value = v
}

// With value embedding
type Derived1 struct {
    Base  // Value embedding
}

d1 := Derived1{}
d1.SetValue(10)  // Works but creates copy issues
// d1.Base is a value, but SetValue needs pointer

// With pointer embedding
type Derived2 struct {
    *Base  // Pointer embedding - must initialize!
}

d2 := Derived2{Base: &Base{}}  // Must allocate
d2.SetValue(10)  // Works correctly

// ✅ BETTER: Be explicit about pointer semantics
```

**Best Practices Summary:**
```go
// ✅ DO: Use embedding for mixins
type Auditable struct { /* ... */ }
type Product struct {
    Auditable  // Adds audit functionality
    // ...
}

// ✅ DO: Use embedding for base functionality
type BaseEntity struct { /* ... */ }
type User struct {
    BaseEntity
    // ...
}

// ✅ DO: Use composition for "has-a" relationships
type Car struct {
    Engine Engine  // Car HAS an Engine
}

// ✅ DO: Keep embedding shallow (1-2 levels max)
type Admin struct {
    User  // OK - one level
}

// ✅ DO: Document why you're embedding
type Product struct {
    Auditable  // Embedding for audit mixin pattern
    // ...
}

// ❌ DON'T: Embed just to save typing
// ❌ DON'T: Embed more than 3-4 types in one struct
// ❌ DON'T: Embed when composition is clearer
// ❌ DON'T: Create deep embedding hierarchies (3+ levels)
```

---

### **11. Comparison to Other Languages (3-4 min)**

**Quick comparison:**

```go
// Go - Composition via Embedding
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println("Some sound")
}

type Dog struct {
    Animal  // Composition, not inheritance!
    Breed string
}

func (d Dog) Speak() {  // Override
    fmt.Println("Woof!")
}

// Python - Class Inheritance
/*
class Animal:
    def __init__(self, name):
        self.name = name
    
    def speak(self):
        print("Some sound")

class Dog(Animal):  # Inheritance
    def __init__(self, name, breed):
        super().__init__(name)  # Call parent constructor
        self.breed = breed
    
    def speak(self):  # Override
        print("Woof!")
*/

// Java - Class Inheritance
/*
class Animal {
    String name;
    void speak() {
        System.out.println("Some sound");
    }
}

class Dog extends Animal {  // Inheritance
    String breed;
    
    @Override
    void speak() {
        System.out.println("Woof!");
    }
}
*/
```

**Key differences:**
- Go: No inheritance keyword, no `extends`, no `super`
- Go: Composition is explicit and visible
- Go: No virtual methods - method resolution is simple
- Go: Can't accidentally break parent class
- Go: Interface satisfaction is implicit
- Go: Simpler mental model - just nested structs

---

### **12. Next Steps & Wrap-up (3-4 min)**

**Recap what was covered:**
- Basic composition with named fields
- Struct embedding (anonymous fields)
- Field and method promotion
- Multiple embedding and conflicts
- Embedding with interfaces
- Practical patterns (mixins, decorators, base entities)
- When to use embedding vs composition
- Common pitfalls to avoid

**Preview next topics:**
- Interfaces in depth
- Polymorphism in Go
- Type assertions and type switches
- Interface composition
- Error handling patterns

**Homework/Practice suggestions:**
1. **Easy:** Create a Shape hierarchy (Shape → Rectangle, Circle) using embedding
2. **Medium:** Build a notification system with different notifier types
3. **Challenge:** Create a plugin system where plugins embed common functionality
4. **Advanced:** Implement a middleware chain using embedding

**Resources:**
- Effective Go on Embedding: golang.org/doc/effective_go#embedding
- Go blog on embedding: blog.golang.org/json-and-go
- Your GitHub repo with complete employee system

---

## **Production Notes**

### **Screen Setup:**
- Code editor: 70% of screen
- Terminal output: 30% of screen
- Use split view when comparing embedding vs composition
- Font size: 18-20pt minimum

### **Teaching Techniques:**
- Start simple (basic composition) → build to complex (multiple embedding)
- Show the problem before the solution
- Use real-world examples (employees, products, notifications)
- Demonstrate both value and pointer embedding
- Intentionally create name conflicts to show compiler errors
- Run code frequently to show promoted fields/methods in action

### **Code Quality Reminders:**
- Use embedding for mixins and base patterns
- Prefer composition for has-a relationships
- Keep embedding shallow (1-2 levels)
- Document why you're embedding
- Be consistent with pointer vs value semantics

### **Engagement:**
- "What fields can we access on manager?" (demonstrate promotion)
- "What happens if we have two embedded types with Name field?"
- "Pause and try creating your own embedded struct"
- Compare with inheritance from languages students might know

### **Visual Aids:**
- Diagram: Struct embedding vs composition memory layout
- Diagram: Method promotion and resolution order
- Diagram: Employee hierarchy with embedded types
- Flowchart: "Should I use embedding?" decision tree

---

## **Supplementary Materials to Provide**

1. **GitHub Repository:**
   - Complete employee management system
   - All code examples from video
   - Practice exercises with solutions
   - README with embedding best practices

2. **Cheat Sheet (PDF/Gist):**
   ```
   Composition:        field Type
   Embedding:          Type (no field name)
   Field promotion:    outer.Field (from embedded)
   Method promotion:   outer.Method() (from embedded)
   Explicit access:    outer.Type.Field
   Multiple embedding: All types promoted
   Conflicts:          Must use Type.Field explicitly
   Override:           Define same method on outer type
   ```

3. **Practice Exercises:**
   - **Easy:** Animal/Dog/Cat hierarchy
   - **Medium:** Vehicle system (Car, Truck, Motorcycle)
   - **Challenge:** Blog system (Post, Comment, Author with audit mixins)
   - **Advanced:** Middleware pipeline with logging, auth, metrics

4. **Decision Flowchart:**
   "Embedding vs Composition" decision tree (PDF)
   
5. **Anti-Pattern Guide:**
   - List of common embedding mistakes
   - Before/After refactoring examples
   - When embedding goes wrong

6. **Real-World Examples:**
   - How standard library uses embedding (io.Reader, http.Handler)
   - Popular Go projects using embedding patterns
   - Database ORM patterns with embedding

---

This tutorial builds naturally on the previous structs video while introducing Go's unique approach to composition. The employee management system is a comprehensive example that ties all concepts together. The emphasis on when NOT to use embedding is as important as when to use it.