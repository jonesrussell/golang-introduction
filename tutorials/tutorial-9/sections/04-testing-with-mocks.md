# Testing with Mocks

**Duration:** 8-10 minutes

## Topics to cover:
- Creating mock implementations
- Testing success cases
- Testing error cases
- Table-driven tests

## Mock Implementation

```go
type MockUserRepository struct {
    users    map[int]*User
    saveErr  error
    findErr  error
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users: make(map[int]*User),
    }
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
    if m.findErr != nil {
        return nil, m.findErr
    }
    user, ok := m.users[id]
    if !ok {
        return nil, ErrNotFound
    }
    return user, nil
}

func (m *MockUserRepository) Save(user *User) error {
    if m.saveErr != nil {
        return m.saveErr
    }
    if user.ID == 0 {
        user.ID = len(m.users) + 1
    }
    m.users[user.ID] = user
    return nil
}

func (m *MockUserRepository) Delete(id int) error {
    delete(m.users, id)
    return nil
}

// Helper methods for setting up test scenarios
func (m *MockUserRepository) SetFindError(err error) {
    m.findErr = err
}

func (m *MockUserRepository) AddUser(user *User) {
    m.users[user.ID] = user
}
```

## Mock Logger

```go
type MockLogger struct {
    Logs []LogEntry
}

type LogEntry struct {
    Level   string
    Message string
    Fields  map[string]interface{}
}

func (m *MockLogger) Info(msg string, fields ...interface{}) {
    m.Logs = append(m.Logs, LogEntry{Level: "info", Message: msg})
}

func (m *MockLogger) Error(msg string, fields ...interface{}) {
    m.Logs = append(m.Logs, LogEntry{Level: "error", Message: msg})
}
```

## Tests

```go
func TestUserService_GetUser_Success(t *testing.T) {
    // Arrange
    repo := NewMockUserRepository()
    repo.AddUser(&User{ID: 1, Name: "Alice", Email: "alice@test.com"})
    logger := &MockLogger{}
    service := NewUserService(repo, logger)

    // Act
    user, err := service.GetUser(1)

    // Assert
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if user.Name != "Alice" {
        t.Errorf("expected Alice, got %s", user.Name)
    }
    if len(logger.Logs) == 0 || logger.Logs[0].Level != "info" {
        t.Error("expected info log")
    }
}

func TestUserService_GetUser_NotFound(t *testing.T) {
    repo := NewMockUserRepository()
    logger := &MockLogger{}
    service := NewUserService(repo, logger)

    _, err := service.GetUser(999)

    if !errors.Is(err, ErrNotFound) {
        t.Errorf("expected ErrNotFound, got %v", err)
    }
}

func TestUserService_GetUser_RepoError(t *testing.T) {
    repo := NewMockUserRepository()
    repo.SetFindError(errors.New("database connection lost"))
    logger := &MockLogger{}
    service := NewUserService(repo, logger)

    _, err := service.GetUser(1)

    if err == nil {
        t.Fatal("expected error")
    }
    // Check error logged
    hasErrorLog := false
    for _, log := range logger.Logs {
        if log.Level == "error" {
            hasErrorLog = true
            break
        }
    }
    if !hasErrorLog {
        t.Error("expected error to be logged")
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
            logger := &MockLogger{}
            service := NewUserService(repo, logger)

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
- Mocks implement same interface
- Set up test scenarios with helper methods
- Test both success and error paths
- Table-driven tests for variations
