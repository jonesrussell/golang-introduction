# Interface Design Principles

**Duration:** 5-6 minutes

## Topics to cover:
- Keep interfaces small
- Accept interfaces, return structs
- Define interfaces at point of use
- Interface segregation

## Code Examples

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

## Key teaching points:
- Smaller interfaces = more implementations possible
- Consumer defines the interface they need
- Producer returns concrete types
- Don't force types to implement unused methods
- Implicit implementation enables decoupling
