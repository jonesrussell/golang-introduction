# Interface Pollution

**Duration:** 5-6 minutes

## The Anti-Pattern

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

## The Fix

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

// Accept interfaces, return structs
func NewHandler(users userGetter) *Handler {
    return &Handler{users: users}
}
```

## When to Create Interfaces:
- Multiple implementations exist
- Testing requires mocking
- Package boundary crossing

## Key teaching points:
- Don't create [interfaces](https://go.dev/ref/spec#Interface_types) for single implementations
- ["Accept interfaces, return structs"](https://go.dev/doc/effective_go#interfaces_and_types)
- Define [interfaces](https://go.dev/ref/spec#Interface_types) at point of use
- Wait until you need the abstraction
