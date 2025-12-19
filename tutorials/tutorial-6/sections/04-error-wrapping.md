# Error Wrapping

**Duration:** 7-8 minutes

## Topics to cover:
- Why wrap errors? (context!)
- `fmt.Errorf` with `%w` verb
- `errors.Unwrap`
- `errors.Is` for comparison
- `errors.As` for type assertion

## Code Examples

```go
import (
    "errors"
    "fmt"
    "os"
)

// Define sentinel errors
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrInvalidInput = errors.New("invalid input")
)

// Wrapping errors adds context
func ReadConfig(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        // Wrap with context - %w preserves original error
        return nil, fmt.Errorf("reading config file %s: %w", path, err)
    }
    return data, nil
}

func LoadConfig() (*Config, error) {
    data, err := ReadConfig("/etc/app/config.json")
    if err != nil {
        // Add more context at each layer
        return nil, fmt.Errorf("loading configuration: %w", err)
    }

    config, err := parseConfig(data)
    if err != nil {
        return nil, fmt.Errorf("parsing configuration: %w", err)
    }

    return config, nil
}

// errors.Is - checks if error is (or wraps) a target error
func main() {
    _, err := LoadConfig()
    if err != nil {
        // Check for specific underlying error
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("Config file not found - using defaults")
            // Use default config
        } else if errors.Is(err, os.ErrPermission) {
            fmt.Println("Permission denied - check file permissions")
        } else {
            fmt.Printf("Configuration error: %v\n", err)
        }
    }

    // Full error chain:
    // "loading configuration: reading config file /etc/app/config.json: open /etc/app/config.json: no such file or directory"
}
```

## errors.As Example

```go
// errors.As - extracts specific error type from chain
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

func ValidateUser(u *User) error {
    if u.Email == "" {
        return &ValidationError{Field: "email", Message: "required"}
    }
    return nil
}

func CreateUser(u *User) error {
    if err := ValidateUser(u); err != nil {
        return fmt.Errorf("user creation failed: %w", err)
    }
    // ...
    return nil
}

func main() {
    err := CreateUser(&User{})
    if err != nil {
        var valErr *ValidationError
        if errors.As(err, &valErr) {
            // Can access ValidationError fields
            fmt.Printf("Fix the %s field: %s\n", valErr.Field, valErr.Message)
        } else {
            fmt.Println("Error:", err)
        }
    }
}
```

## Key teaching points:
- `%w` wraps error, preserving the chain
- `%v` formats error but breaks the chain
- `errors.Is()` checks entire chain for match
- `errors.As()` extracts typed error from chain
- Add context at each layer for debugging
- Error messages should flow: "outer: middle: inner: root"
