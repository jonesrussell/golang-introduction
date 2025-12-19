# The Problem: Tight Coupling

**Duration:** 5-6 minutes

## The Problem

```go
// TIGHTLY COUPLED - Hard to test
type UserService struct {
    // Direct dependency on concrete database
}

func (s *UserService) GetUser(id int) (*User, error) {
    // Directly using global database connection
    db, _ := sql.Open("postgres", "connection-string")
    defer db.Close()

    var user User
    err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Name, &user.Email)
    return &user, err
}
```

## Why This Is Problematic:
- Tests require real database
- Can't test error conditions easily
- Slow tests (real I/O)
- Flaky tests (external dependencies)
- Hidden dependencies
- Can't swap implementations

## Key teaching points:
- Tight coupling makes code hard to test
- Direct dependencies on concrete types reduce flexibility
- [Dependency injection](https://go.dev/doc/effective_go#interfaces_and_types) solves these problems
