# Visibility Rules

**Duration:** 5-6 minutes

## Exported vs Unexported

```go
// user/user.go
package user

// Exported - used by other packages
type User struct {
    ID    int      // Exported field
    Name  string   // Exported field
    email string   // Unexported - only this package
}

// Exported function
func New(name, email string) *User {
    return &User{
        Name:  name,
        email: email,
    }
}

// Exported method
func (u *User) GetEmail() string {
    return u.email  // Can access unexported from same package
}

// Unexported function - helper
func validateEmail(email string) bool {
    return strings.Contains(email, "@")
}
```

## From Another Package

```go
// main.go
package main

import "github.com/project/user"

func main() {
    u := user.New("Alice", "alice@example.com")
    fmt.Println(u.Name)          // OK - exported field
    // fmt.Println(u.email)      // ERROR - unexported
    fmt.Println(u.GetEmail())    // OK - via exported method
}
```

## internal/ Enforcement

```
project/
├── internal/
│   └── secret/
│       └── secret.go    # package secret
├── cmd/
│   └── app/
│       └── main.go      # Can import internal/secret
└── go.mod

# Another project cannot import:
# import "github.com/project/internal/secret"  // ERROR!
```
