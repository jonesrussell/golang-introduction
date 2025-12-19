# Sentinel Errors

**Duration:** 4-5 minutes

## Topics to cover:
- What are sentinel errors?
- When to use them
- Package-level error variables
- Comparing with [`errors.Is`](https://pkg.go.dev/errors#Is)

## Code Examples

```go
// Sentinel errors - predefined error values
// Convention: ErrXxx naming

package user

import "errors"

// Package-level sentinel errors
var (
    ErrNotFound       = errors.New("user not found")
    ErrDuplicateEmail = errors.New("email already exists")
    ErrInvalidAge     = errors.New("age must be positive")
    ErrUnauthorized   = errors.New("unauthorized access")
)

func GetUser(id int) (*User, error) {
    user := db.Find(id)
    if user == nil {
        return nil, ErrNotFound
    }
    return user, nil
}

func CreateUser(u *User) error {
    if exists := db.FindByEmail(u.Email); exists != nil {
        return ErrDuplicateEmail
    }
    if u.Age < 0 {
        return ErrInvalidAge
    }
    return db.Save(u)
}

// Caller code
func main() {
    user, err := user.GetUser(123)
    if err != nil {
        switch {
        case errors.Is(err, user.ErrNotFound):
            fmt.Println("User doesn't exist")
        case errors.Is(err, user.ErrUnauthorized):
            fmt.Println("You don't have permission")
        default:
            fmt.Println("Unexpected error:", err)
        }
    }
}
```

## Wrapping Sentinel Errors

```go
// Wrapping sentinel errors
func GetUserWithContext(id int) (*User, error) {
    user, err := db.Find(id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            // Wrap with our sentinel error
            return nil, fmt.Errorf("%w: id=%d", ErrNotFound, id)
        }
        return nil, fmt.Errorf("database error: %w", err)
    }
    return user, nil
}
```

## Standard Library Sentinel Errors

```go
import (
    "io"
    "os"
)

func ProcessFile(path string) error {
    f, err := os.Open(path)
    if errors.Is(err, os.ErrNotExist) {
        // File doesn't exist
    }
    if errors.Is(err, os.ErrPermission) {
        // No permission
    }

    // Reading
    buf := make([]byte, 1024)
    _, err = f.Read(buf)
    if errors.Is(err, io.EOF) {
        // End of file (not really an error)
    }
}
```

## When to Use Sentinel vs Custom Type

```go
// GOOD: Sentinel errors
// - Error has no dynamic data
// - Callers need to check for specific error
// - Error is part of your API contract

var ErrNotFound = errors.New("not found")

// BETTER: Custom error type
// - Error needs context (which user? which file?)
// - Multiple pieces of information needed

type NotFoundError struct {
    Resource string
    ID       int
}

// AVOID: String comparison
// BAD:
if err.Error() == "not found" {  // Fragile!
    // ...
}

// GOOD:
if errors.Is(err, ErrNotFound) {  // Robust
    // ...
}
```

## Key teaching points:
- Sentinel errors are package-level variables
- Use [`errors.Is()`](https://pkg.go.dev/errors#Is) to compare, not `==` (for wrapped errors)
- Name convention: `ErrXxx`
- Document sentinel errors as part of API
- Choose sentinel vs custom type based on needs
