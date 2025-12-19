# Interface Design for DI

**Duration:** 6-7 minutes

## Topics to cover:
- Small, focused [interfaces](https://go.dev/ref/spec#Interface_types)
- Consumer-defined interfaces
- Interface segregation

## BAD: Large Interface (Hard to Mock)

```go
type DatabaseInterface interface {
    Query(sql string, args ...interface{}) (*sql.Rows, error)
    Exec(sql string, args ...interface{}) (sql.Result, error)
    Begin() (*sql.Tx, error)
    Prepare(sql string) (*sql.Stmt, error)
    Ping() error
    Close() error
    // ... 20 more methods
}
```

## GOOD: Small, Focused Interfaces

```go
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
```

## BEST: Consumer-Defined Interfaces

```go
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
```

## Interface Segregation

```go
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

## Key teaching points:
- Small [interfaces](https://go.dev/ref/spec#Interface_types) = easier mocking
- Consumer defines what it needs
- One implementation can satisfy many interfaces
- Don't export interfaces unnecessarily
