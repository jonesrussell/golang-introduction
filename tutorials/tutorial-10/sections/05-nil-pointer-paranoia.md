# Nil Pointer Paranoia / Over-Checking

**Duration:** 4-5 minutes

## The Anti-Pattern

```go
// BAD: Nil checks everywhere
func ProcessUser(user *User) error {
    if user == nil {
        return errors.New("user is nil")
    }
    if user.Profile == nil {
        return errors.New("profile is nil")
    }
    if user.Profile.Address == nil {
        return errors.New("address is nil")
    }
    if user.Profile.Address.City == nil {
        return errors.New("city is nil")
    }

    city := *user.Profile.Address.City
    // ...
}
```

## The Fix

```go
// GOOD: Design to avoid nil
type User struct {
    Profile Profile  // Value, not pointer - never nil
}

type Profile struct {
    Address Address  // Value, not pointer
}

type Address struct {
    City string  // Value, not pointer
}

// GOOD: Use constructor to ensure valid state
func NewUser(name string) *User {
    return &User{
        Name: name,
        Profile: Profile{
            Address: Address{
                City: "Unknown",
            },
        },
    }
}

// GOOD: Check at boundaries, trust internal code
func (h *Handler) HandleRequest(r *http.Request) {
    // Validate input at boundary
    user, err := parseUser(r)
    if err != nil {
        // Handle invalid input
        return
    }

    // Internal code can trust user is valid
    h.service.ProcessUser(user)
}
```

## When Pointer Is Intentional (Optional Field)

```go
type Config struct {
    Timeout *time.Duration  // nil means "use default"
}

func (c *Config) GetTimeout() time.Duration {
    if c.Timeout == nil {
        return 30 * time.Second
    }
    return *c.Timeout
}
```

## Key teaching points:
- Design types to avoid [nil](https://go.dev/ref/spec#The_zero_value) when possible
- Use [value types](https://go.dev/ref/spec#Types) instead of pointers for required fields
- Check at boundaries, trust internal code
- Use pointers only when [nil is meaningful](https://go.dev/ref/spec#The_zero_value)
