# The Problem: Code Reuse Without Inheritance

**Duration:** 4-5 minutes

## Topics to cover:
- Traditional OOP inheritance (what Go doesn't have)
- The diamond problem and why inheritance is complex
- Go's alternative: composition and embedding

## Code Example - The Problem

```go
// Imagine we want to model different types of users
// In traditional OOP (pseudocode):
// class User {
//     name, email
// }
// class Admin extends User {
//     permissions
// }
// class Customer extends User {
//     orderHistory
// }

// Without embedding, we'd repeat fields:
type Admin struct {
    Name        string  // Repeated
    Email       string  // Repeated
    Permissions []string
}

type Customer struct {
    Name         string  // Repeated
    Email        string  // Repeated
    OrderHistory []Order
}

// This violates DRY principle and is hard to maintain
// What if we want to add a PhoneNumber field to all users?
```

## Key teaching points:
- Go has no class inheritance
- Go has no extends or super keywords
- This is intentional - inheritance creates tight coupling
- Go provides composition instead
- Embedding is Go's answer to code reuse
