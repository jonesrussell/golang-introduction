## **Video Tutorial Plan: Go Structs - Definition, Initialization, and Methods**

### **Video Metadata**
- **Title:** Go Structs: Definition, Initialization, and Methods
- **Duration Target:** 30-40 minutes
- **Difficulty:** Beginner to Intermediate
- **Prerequisites:** Go Basics (variables, types, functions)

---

## **Video Structure**

### **1. Introduction (2-3 min)**
- Welcome and what viewers will learn
- Why structs? (Organizing related data, modeling real-world entities)
- Comparison to other languages: structs vs classes
- Go philosophy: composition over inheritance
- Show the final example: Building a User management system

---

### **2. What Are Structs? (3-4 min)**

**Topics to cover:**
- Structs as custom types
- Grouping related data together
- Type definition syntax
- When to use structs vs primitives

**Code Example:**
```go
// Without structs - messy and error-prone
firstName := "Jane"
lastName := "Smith"
email := "jane@example.com"
age := 28

// With structs - organized and type-safe
type User struct {
    FirstName string
    LastName  string
    Email     string
    Age       int
}
```

**Key teaching points:**
- Structs create custom types
- Fields (properties) are declared with name and type
- Exported fields start with capital letters (public)
- Lowercase fields are package-private
- Structs are value types (copied by default)

---

### **3. Defining Structs (5-6 min)**

**Topics to cover:**
- Basic struct definition
- Field naming conventions
- Field tags (preview for JSON)
- Nested structs
- Anonymous fields

**Code Examples:**
```go
// Basic struct definition
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// Struct with multiple field types
type Product struct {
    ID          int
    Name        string
    Price       float64
    InStock     bool
    Tags        []string
}

// Struct with field tags (for JSON, validation, etc.)
type User struct {
    ID        int    `json:"id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    CreatedAt string `json:"created_at"`
}

// Nested structs
type Address struct {
    Street  string
    City    string
    ZipCode string
}

type Employee struct {
    Name    string
    Email   string
    Address Address  // Nested struct
}

// Anonymous struct (inline, one-off use)
config := struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080,
}
```

**Key teaching points:**
- Type definition creates a new type in your package
- Fields are accessed with dot notation
- Field tags are metadata (used by encoding/json, validation libraries)
- Nested structs model "has-a" relationships
- Anonymous structs are useful for one-off data structures
- Convention: one struct per logical concept

---

### **4. Struct Initialization (7-8 min)**

**Topics to cover:**
- Zero value initialization
- Struct literals
- Named field initialization
- Positional initialization (avoid)
- Pointer to struct with `&`
- Constructor functions (idiomatic pattern)

**Code Examples:**
```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// 1. Zero value - all fields get their type's zero value
var p1 Person
fmt.Printf("%+v\n", p1)  // {FirstName: LastName: Age:0}

// 2. Struct literal with named fields (RECOMMENDED)
p2 := Person{
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
}

// 3. Partial initialization (remaining fields get zero values)
p3 := Person{
    FirstName: "Jane",
    LastName:  "Smith",
    // Age will be 0
}

// 4. Positional initialization (AVOID - brittle to field reordering)
p4 := Person{"Bob", "Wilson", 25}  // Works but not recommended

// 5. Pointer to struct (common for large structs or when mutating)
p5 := &Person{
    FirstName: "Alice",
    LastName:  "Johnson",
    Age:       28,
}

// 6. Using new() - returns pointer with zero values
p6 := new(Person)
p6.FirstName = "Charlie"

// 7. Constructor function (IDIOMATIC PATTERN)
func NewPerson(firstName, lastName string, age int) *Person {
    return &Person{
        FirstName: firstName,
        LastName:  lastName,
        Age:       age,
    }
}

// Usage
p7 := NewPerson("David", "Brown", 35)
```

**Key teaching points:**
- Named field initialization is most readable and maintainable
- Trailing comma is required on multi-line initialization
- Zero values make structs safe to use even when empty
- Pointer initialization with `&` avoids copying large structs
- Constructor functions allow validation and default values
- Convention: `New` prefix for constructors (NewUser, NewProduct, etc.)

---

### **5. Accessing and Modifying Struct Fields (4-5 min)**

**Topics to cover:**
- Dot notation
- Automatic pointer dereferencing
- Value vs pointer receivers (preview)
- Comparing structs

**Code Examples:**
```go
type Book struct {
    Title  string
    Author string
    Pages  int
}

// Create a book
book := Book{
    Title:  "The Go Programming Language",
    Author: "Donovan & Kernighan",
    Pages:  380,
}

// Accessing fields
fmt.Println(book.Title)   // Read
book.Pages = 400          // Modify

// Working with pointers - automatic dereferencing
bookPtr := &Book{
    Title:  "Learning Go",
    Author: "Jon Bodner",
    Pages:  350,
}

// Go automatically dereferences, these are equivalent:
fmt.Println(bookPtr.Title)
fmt.Println((*bookPtr).Title)

// Comparing structs (all fields must be comparable)
book1 := Book{Title: "Go", Author: "Pike", Pages: 200}
book2 := Book{Title: "Go", Author: "Pike", Pages: 200}
book3 := Book{Title: "Rust", Author: "Klabnik", Pages: 500}

fmt.Println(book1 == book2)  // true (all fields match)
fmt.Println(book1 == book3)  // false

// Structs with slices/maps cannot be compared with ==
type Library struct {
    Name  string
    Books []Book  // Contains slice - not comparable!
}

// This would be a compilation error:
// lib1 := Library{Name: "City Library"}
// lib2 := Library{Name: "City Library"}
// fmt.Println(lib1 == lib2)  // ERROR!
```

**Key teaching points:**
- Dot notation works the same for values and pointers
- Go's automatic dereferencing reduces pointer boilerplate
- Struct comparison requires all fields to be comparable
- Slices, maps, and functions make structs non-comparable
- Use reflect.DeepEqual() for complex comparisons

---

### **6. Methods on Structs (8-10 min)**

**Topics to cover:**
- Method definition syntax
- Value receivers vs pointer receivers
- When to use each receiver type
- Method chaining
- Methods vs functions

**Code Examples:**
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver (does NOT modify original)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with value receiver
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Method with pointer receiver (CAN modify original)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Method with pointer receiver for consistency
func (r *Rectangle) String() string {
    return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

// Usage
rect := Rectangle{Width: 10, Height: 5}

// Value receiver methods
area := rect.Area()           // 50
perimeter := rect.Perimeter() // 30

// Pointer receiver method - modifies original
rect.Scale(2)
fmt.Println(rect.Width)  // 20
fmt.Println(rect.Height) // 10

// Go allows calling pointer receiver methods on values (and vice versa)
rect2 := Rectangle{Width: 5, Height: 3}
rect2.Scale(3)  // Go automatically does (&rect2).Scale(3)

pRect := &Rectangle{Width: 7, Height: 4}
area2 := pRect.Area()  // Go automatically does (*pRect).Area()
```

**More examples:**
```go
type BankAccount struct {
    Owner   string
    Balance float64
}

// Constructor
func NewBankAccount(owner string, initialBalance float64) *BankAccount {
    return &BankAccount{
        Owner:   owner,
        Balance: initialBalance,
    }
}

// Pointer receiver - modifies state
func (ba *BankAccount) Deposit(amount float64) {
    if amount > 0 {
        ba.Balance += amount
    }
}

// Pointer receiver - modifies state
func (ba *BankAccount) Withdraw(amount float64) bool {
    if amount > 0 && amount <= ba.Balance {
        ba.Balance -= amount
        return true
    }
    return false
}

// Value receiver - just reads data
func (ba BankAccount) GetBalance() float64 {
    return ba.Balance
}

// Value receiver - formats output
func (ba BankAccount) String() string {
    return fmt.Sprintf("%s's account: $%.2f", ba.Owner, ba.Balance)
}

// Usage
account := NewBankAccount("Alice", 1000)
account.Deposit(500)
account.Withdraw(200)
fmt.Println(account)  // Alice's account: $1300.00
```

**Key teaching points:**
- Value receivers: `func (r Rectangle)` - gets a copy, cannot modify original
- Pointer receivers: `func (r *Rectangle)` - gets pointer, can modify original
- Go automatically converts between values and pointers for method calls
- **When to use pointer receivers:**
  - Method needs to modify the receiver
  - Struct is large (avoid copying)
  - Consistency: if any method uses pointer receiver, all should
- **When to use value receivers:**
  - Method only reads data
  - Struct is small (a few fields)
  - You want immutability guarantees
- Convention: be consistent within a type (all pointer or all value)

---

### **7. Practical Example: Building a User System (8-10 min)**

**Build together:** A complete User management system

```go
package main

import (
    "fmt"
    "strings"
    "time"
)

// User represents a system user
type User struct {
    ID        int
    Username  string
    Email     string
    CreatedAt time.Time
    IsActive  bool
}

// NewUser creates a new user with validation
func NewUser(id int, username, email string) (*User, error) {
    // Validation
    if username == "" {
        return nil, fmt.Errorf("username cannot be empty")
    }
    if !strings.Contains(email, "@") {
        return nil, fmt.Errorf("invalid email format")
    }
    
    return &User{
        ID:        id,
        Username:  username,
        Email:     email,
        CreatedAt: time.Now(),
        IsActive:  true,
    }, nil
}

// Methods with value receivers (read-only)

func (u User) GetDisplayName() string {
    return fmt.Sprintf("@%s", u.Username)
}

func (u User) GetAccountAge() time.Duration {
    return time.Since(u.CreatedAt)
}

func (u User) String() string {
    status := "Active"
    if !u.IsActive {
        status = "Inactive"
    }
    return fmt.Sprintf("User %d: %s (%s) - %s", 
        u.ID, u.Username, u.Email, status)
}

// Methods with pointer receivers (modify state)

func (u *User) Deactivate() {
    u.IsActive = false
}

func (u *User) Activate() {
    u.IsActive = true
}

func (u *User) UpdateEmail(newEmail string) error {
    if !strings.Contains(newEmail, "@") {
        return fmt.Errorf("invalid email format")
    }
    u.Email = newEmail
    return nil
}

// UserRepository manages multiple users
type UserRepository struct {
    users  []*User
    nextID int
}

// NewUserRepository creates a new repository
func NewUserRepository() *UserRepository {
    return &UserRepository{
        users:  make([]*User, 0),
        nextID: 1,
    }
}

func (ur *UserRepository) AddUser(username, email string) (*User, error) {
    user, err := NewUser(ur.nextID, username, email)
    if err != nil {
        return nil, err
    }
    
    ur.users = append(ur.users, user)
    ur.nextID++
    return user, nil
}

func (ur *UserRepository) FindByID(id int) *User {
    for _, user := range ur.users {
        if user.ID == id {
            return user
        }
    }
    return nil
}

func (ur *UserRepository) FindByUsername(username string) *User {
    for _, user := range ur.users {
        if user.Username == username {
            return user
        }
    }
    return nil
}

func (ur *UserRepository) ListActiveUsers() []*User {
    active := make([]*User, 0)
    for _, user := range ur.users {
        if user.IsActive {
            active = append(active, user)
        }
    }
    return active
}

func (ur UserRepository) GetUserCount() int {
    return len(ur.users)
}

func main() {
    // Create repository
    repo := NewUserRepository()
    
    // Add users
    user1, err := repo.AddUser("alice", "alice@example.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    user2, _ := repo.AddUser("bob", "bob@example.com")
    user3, _ := repo.AddUser("charlie", "charlie@example.com")
    
    // Display user info
    fmt.Println(user1)
    fmt.Println(user1.GetDisplayName())
    
    // Modify user
    user2.Deactivate()
    err = user3.UpdateEmail("charlie.new@example.com")
    if err != nil {
        fmt.Println("Update failed:", err)
    }
    
    // Query repository
    fmt.Printf("\nTotal users: %d\n", repo.GetUserCount())
    
    found := repo.FindByUsername("alice")
    if found != nil {
        fmt.Println("Found:", found)
    }
    
    // List active users
    fmt.Println("\nActive users:")
    for _, user := range repo.ListActiveUsers() {
        fmt.Printf("  %s\n", user.GetDisplayName())
    }
    
    // Show account ages
    fmt.Println("\nAccount ages:")
    time.Sleep(100 * time.Millisecond)  // Small delay to show time passage
    for _, user := range repo.ListActiveUsers() {
        fmt.Printf("  %s: %v\n", user.Username, user.GetAccountAge())
    }
}
```

**Walk through:**
- Define User struct with appropriate fields
- Constructor with validation (returns error)
- Value receiver methods for reading data
- Pointer receiver methods for mutations
- Repository pattern for managing collections
- Practical usage showing all concepts together
- Error handling in constructors

---

### **8. Best Practices & Code Smells (5-6 min)**

**Cover these important patterns:**

```go
// ✅ GOOD: Constructor with validation
func NewProduct(name string, price float64) (*Product, error) {
    if price < 0 {
        return nil, fmt.Errorf("price cannot be negative")
    }
    return &Product{Name: name, Price: price}, nil
}

// ❌ BAD: No validation, invalid state possible
func NewProduct(name string, price float64) *Product {
    return &Product{Name: name, Price: price}  // Could have negative price!
}

// ✅ GOOD: Consistent receiver types (all pointers)
type Counter struct {
    count int
}

func (c *Counter) Increment() { c.count++ }
func (c *Counter) Decrement() { c.count-- }
func (c *Counter) Value() int { return c.count }  // Even for read-only!

// ❌ BAD: Mixed receiver types
func (c *Counter) Increment() { c.count++ }
func (c Counter) Value() int { return c.count }  // Inconsistent!

// ✅ GOOD: Named field initialization
person := Person{
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
}

// ❌ BAD: Positional initialization (brittle)
person := Person{"John", "Doe", 30}

// ✅ GOOD: Pointer for structs that will be modified
user := &User{Name: "Alice"}
user.Activate()

// ✅ GOOD: Value for small, immutable structs
point := Point{X: 10, Y: 20}
distance := point.DistanceFromOrigin()

// ✅ GOOD: Export only what's necessary
type user struct {  // Lowercase = private
    id       int
    password string  // Keep private!
}

func (u user) GetID() int { return u.id }  // Controlled access

// ❌ BAD: Everything public (no encapsulation)
type User struct {
    ID       int
    Password string  // Should be private!
}

// ✅ GOOD: Zero value is useful
type Config struct {
    MaxRetries int  // 0 is sensible default
    Timeout    int  // 0 could mean no timeout
}

// Consider providing defaults via constructor
func NewConfig() *Config {
    return &Config{
        MaxRetries: 3,
        Timeout:    30,
    }
}
```

**Common mistakes:**
```go
// Mistake 1: Forgetting pointer receiver when modifying
func (u User) Deactivate() {  // Value receiver!
    u.IsActive = false  // Modifies copy, not original!
}

// Mistake 2: Nil pointer dereference
var user *User
fmt.Println(user.Username)  // PANIC! user is nil

// Safe approach:
if user != nil {
    fmt.Println(user.Username)
}

// Mistake 3: Comparing structs with non-comparable fields
type User struct {
    Name    string
    Friends []string  // Slice is not comparable!
}

// This won't compile:
// if user1 == user2 { ... }

// Mistake 4: Large struct by value (expensive copying)
type HugeStruct struct {
    Data [1000000]int
}

func process(h HugeStruct) {  // Copies entire array!
    // ...
}

// Better:
func process(h *HugeStruct) {  // Just copies pointer
    // ...
}
```

---

### **9. When to Use Structs vs Other Types (3-4 min)**

**Decision guide:**

```go
// Use struct when: Grouping related data
type Address struct {
    Street, City, State, Zip string
}

// Use map when: Dynamic keys, not fixed structure
userScores := map[string]int{
    "alice": 95,
    "bob":   87,
}

// Use struct when: Fixed fields with different types
type Event struct {
    Name      string
    Timestamp time.Time
    Attendees []string
    Metadata  map[string]interface{}
}

// Use interface when: Behavior matters more than data
type Writer interface {
    Write([]byte) (int, error)
}

// Use struct when: You need methods with state
type Database struct {
    conn *sql.DB
}

func (db *Database) Query(sql string) (*Rows, error) {
    // Uses db.conn
}
```

**Guidelines:**
- Struct: Fixed schema, type safety, methods
- Map: Dynamic keys, uniform value types
- Interface: Define contracts, polymorphism
- Slice: Ordered collection of same type

---

### **10. Advanced Preview & Next Steps (3-4 min)**

**Brief preview of upcoming topics:**

```go
// Struct embedding (next video)
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person        // Embedded struct
    EmployeeID int
    Department string
}

// JSON encoding/decoding
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

data, _ := json.Marshal(user)

// Interfaces with structs
type Stringer interface {
    String() string
}

// Any type with String() method satisfies Stringer
```

**Recap what was covered:**
- Defining structs
- Initialization patterns (literals, constructors)
- Value vs pointer receivers
- Methods and when to use each
- Best practices and common mistakes

**Homework/Practice suggestions:**
1. **Easy:** Create a Book struct with methods for GetFullTitle, IsLongBook
2. **Medium:** Build a TodoList struct that manages Todo items
3. **Challenge:** Create a BankAccount system with Account and Transaction structs

**Resources:**
- Effective Go: golang.org/doc/effective_go
- Go by Example: gobyexample.com/structs
- Your GitHub repo with complete User system code

---

## **Production Notes**

### **Screen Setup:**
- Code editor on left (75% screen)
- Terminal/output on right (25%)
- Use split view for comparing value vs pointer receivers side-by-side
- Font size: 18-20pt minimum

### **Teaching Techniques:**
- Start simple (basic struct) then build complexity
- Show mistakes intentionally (forgetting pointer receiver, nil pointer)
- Use real-world examples (User, Product, BankAccount)
- Run code after each concept to demonstrate behavior
- Use `fmt.Printf("%+v\n", struct)` to show struct contents clearly

### **Code Quality Reminders:**
- Always use named field initialization
- Be consistent with receiver types
- Validate in constructors
- Return errors from constructors when validation fails
- Use pointer receivers for: modification, large structs, consistency

### **Engagement:**
- "Pause here and predict what happens when we modify this"
- "Try creating your own struct for [X]"
- "Notice how Go prevents this common bug..."
- Show before/after refactoring

### **Visual Aids:**
- Diagram: Value vs Pointer in memory
- Diagram: Method call with value receiver (copy created)
- Diagram: Method call with pointer receiver (original modified)

---

## **Supplementary Materials to Provide**

1. **GitHub Repository:**
   - Complete User management system
   - All code examples from video
   - Practice exercises with solutions
   - README with struct best practices

2. **Cheat Sheet (PDF/Gist):**
   ```
   Struct Definition:      type Name struct { ... }
   Initialization:         obj := Name{Field: value}
   Constructor:            func NewName() *Name { ... }
   Value Receiver:         func (n Name) Method() { ... }
   Pointer Receiver:       func (n *Name) Method() { ... }
   ```

3. **Practice Exercises:**
   - **Easy:** Product struct with Price, Name, InStock
   - **Medium:** Library system with Book and Member structs
   - **Challenge:** Inventory system with Product, Order, Customer
   - **Advanced:** Implement a simple cache with TTL using structs

4. **Decision Tree Flowchart:**
   "Should I use a pointer receiver?" flowchart
   - Does method modify receiver? → Yes: Use pointer
   - Is struct large (> a few fields)? → Yes: Use pointer
   - Do other methods use pointer receiver? → Yes: Use pointer for consistency
   - Otherwise → Value receiver is fine

---

This tutorial builds on the basics while introducing object-oriented patterns in Go. The practical example ties everything together and shows how structs form the foundation of Go applications. The emphasis on value vs pointer receivers is critical for avoiding bugs.
