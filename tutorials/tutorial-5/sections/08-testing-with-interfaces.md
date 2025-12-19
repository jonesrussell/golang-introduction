# Testing with Interfaces

**Duration:** 5-6 minutes

## Topics to cover:
- Mock implementations
- Interface-based testing
- Dependency injection for testability

## Code Examples

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
func TestUserService_GetUser_Success(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    mockRepo.users[1] = &User{ID: 1, Name: "Alice", Email: "alice@test.com"}

    service := NewUserService(mockRepo)

    // Act
    user, err := service.GetUser(1)

    // Assert
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
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

## Table-Driven Tests

```go
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
            service := NewUserService(repo)

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

## Key teaching points:
- Interfaces enable mock implementations
- Test behavior, not implementation
- Mock can simulate error conditions
- No need for mocking frameworks
- Production code depends on interfaces
- Tests inject mock implementations
