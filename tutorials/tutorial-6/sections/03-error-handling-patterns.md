# Error Handling Patterns

**Duration:** 6-7 minutes

## Topics to cover:
- The if err != nil pattern
- Early returns
- Error propagation
- Don't ignore errors!

## Code Examples

```go
// Pattern 1: Basic error checking
func ReadFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err  // Propagate error
    }
    defer file.Close()

    // Continue with file operations
    return nil
}

// Pattern 2: Handle and continue
func ProcessFiles(paths []string) {
    for _, path := range paths {
        err := processFile(path)
        if err != nil {
            fmt.Printf("Warning: %s - %v\n", path, err)
            continue  // Handle error and continue
        }
        fmt.Printf("Processed: %s\n", path)
    }
}

// Pattern 3: Multiple error checks
func CreateUser(name, email string) (*User, error) {
    // Validate name
    if name == "" {
        return nil, errors.New("name cannot be empty")
    }

    // Validate email
    if !strings.Contains(email, "@") {
        return nil, errors.New("invalid email format")
    }

    // Create in database
    user, err := db.Create(name, email)
    if err != nil {
        return nil, err
    }

    // Send welcome email
    err = sendWelcomeEmail(user.Email)
    if err != nil {
        // Log but don't fail user creation
        log.Printf("Warning: welcome email failed: %v", err)
    }

    return user, nil
}

// Pattern 4: Cleanup on error
func ProcessWithCleanup() error {
    resource, err := acquireResource()
    if err != nil {
        return err
    }
    defer resource.Release()  // Always cleanup

    err = doWork(resource)
    if err != nil {
        return err  // resource.Release() still called
    }

    return nil
}

// Anti-pattern: Ignoring errors
func BadExample() {
    file, _ := os.Open("file.txt")  // BAD: ignoring error!
    // If file.txt doesn't exist, file is nil
    // Next line will panic
    file.Read(make([]byte, 100))
}

// Pattern 5: Inline error declaration
func ProcessData() error {
    if err := validateInput(); err != nil {
        return err
    }

    if data, err := fetchData(); err != nil {
        return err
    } else {
        return processData(data)
    }
}
```

## Key teaching points:
- Always check returned errors
- Use early returns to reduce nesting
- `_` to ignore is almost always wrong
- [Defer](https://go.dev/ref/spec#Defer_statements) ensures cleanup even on error
- Log non-critical errors, return critical ones
