# Package Design Principles

**Duration:** 5-6 minutes

## Guidelines

```go
// 1. Single responsibility
// BAD: package utils (grab bag of unrelated functions)
// GOOD: package validation, package encryption, package http

// 2. Name by what it provides, not what it contains
// BAD: package models, package types, package common
// GOOD: package user, package order, package payment

// 3. Avoid stutter
// BAD:
package user
type UserService struct{}  // user.UserService stutters

// GOOD:
package user
type Service struct{}  // user.Service

// 4. Package-level documentation
// user/doc.go
/*
Package user provides functionality for managing user accounts.

It supports creating, updating, and authenticating users.
Sessions are managed through the Session type.

Basic usage:

    svc := user.NewService(db)
    user, err := svc.Create("alice@example.com")

For authentication:

    session, err := svc.Authenticate(email, password)
*/
package user

// 5. Minimize exported surface
// Export only what consumers need
// Start unexported, export when needed

// 6. Accept interfaces, return structs
func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}
```

## Key teaching points:
- Packages should have [single responsibility](https://go.dev/doc/effective_go#package-names)
- Name packages by what they provide, not what they contain
- Avoid [stutter](https://go.dev/doc/effective_go#package-names) in package names
- Document packages with [package comments](https://go.dev/doc/effective_go#commentary)
- Minimize [exported surface](https://go.dev/ref/spec#Exported_identifiers)
