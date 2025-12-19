# Practical Example: Building a User System

**Duration:** 8-10 minutes

## Build Together

A complete User management system demonstrating all struct concepts.

```go runnable
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
}
```

## Walk Through:
- Define User struct with appropriate fields
- Constructor with validation (returns error)
- Value receiver methods for reading data
- Pointer receiver methods for mutations
- Repository pattern for managing collections
- Practical usage showing all concepts together
- Error handling in constructors
