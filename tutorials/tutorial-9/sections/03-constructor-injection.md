# Constructor Injection Pattern

**Duration:** 8-10 minutes

## Topics to cover:
- Define [interfaces](https://go.dev/ref/spec#Interface_types) for dependencies
- Accept interfaces in constructors
- Store dependencies as fields

## Code Examples

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
```

## Production Implementation

```go
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
```

## Wiring It Up

```go
func main() {
    db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    repo := NewPostgresUserRepository(db)
    logger := NewZapLogger()
    service := NewUserService(repo, logger)

    // Now use service
    user, _ := service.GetUser(1)
}
```

## Key teaching points:
- Define small [interfaces](https://go.dev/ref/spec#Interface_types)
- ["Accept interfaces, return structs"](https://go.dev/doc/effective_go#interfaces_and_types)
- Constructors wire dependencies
- Easy to swap implementations
