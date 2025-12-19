# Struct Embedding - The Basics

**Duration:** 7-8 minutes

## Topics to cover:
- [Embedding syntax](https://go.dev/ref/spec#Struct_types) (anonymous fields)
- [Field promotion](https://go.dev/ref/spec#Selectors)
- [Method promotion](https://go.dev/ref/spec#Selectors)
- Difference between embedding and composition

## Code Examples

```go
// Embedding - anonymous field (no field name)

type User struct {
    ID       int
    Username string
    Email    string
}

func (u User) GetDisplayName() string {
    return fmt.Sprintf("@%s", u.Username)
}

func (u User) SendEmail(subject string) {
    fmt.Printf("Sending '%s' to %s\n", subject, u.Email)
}

// Admin embeds User
type Admin struct {
    User              // Embedded struct - NO field name
    Permissions []string
}

// Usage - field promotion
admin := Admin{
    User: User{
        ID:       1,
        Username: "admin",
        Email:    "admin@example.com",
    },
    Permissions: []string{"read", "write", "delete"},
}

// Can access User fields directly (promoted)
fmt.Println(admin.Username)  // Not admin.User.Username
fmt.Println(admin.Email)     // Not admin.User.Email

// Can also access through type name
fmt.Println(admin.User.Username)  // Still works

// Method promotion - User methods available on Admin
fmt.Println(admin.GetDisplayName())  // Promoted method
admin.SendEmail("Welcome")           // Promoted method

// Admin-specific fields
fmt.Println(admin.Permissions)
```

## Customer Example

```go
type Customer struct {
    User         // Embedded
    OrderCount   int
    LoyaltyPoints int
}

customer := Customer{
    User: User{
        ID:       2,
        Username: "john_doe",
        Email:    "john@example.com",
    },
    OrderCount:   15,
    LoyaltyPoints: 150,
}

// All User fields and methods promoted
fmt.Println(customer.Username)        // Promoted field
fmt.Println(customer.GetDisplayName()) // Promoted method

// Customer-specific functionality
func (c Customer) GetTier() string {
    if c.OrderCount > 10 {
        return "Gold"
    }
    return "Silver"
}

fmt.Println(customer.GetTier())  // Customer's own method
```

## Key teaching points:
- [Embedding](https://go.dev/ref/spec#Struct_types) = anonymous field (type without name)
- Embedded fields are ["promoted"](https://go.dev/ref/spec#Selectors) to outer struct
- Can access embedded fields/methods [directly](https://go.dev/ref/spec#Selectors)
- Can still access through [type name](https://go.dev/ref/spec#Selectors) if needed
- Looks like inheritance but it's [composition](https://go.dev/doc/faq#inheritance)
- The embedded struct doesn't know it's embedded
