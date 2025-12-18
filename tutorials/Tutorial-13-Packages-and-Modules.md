## **Video Tutorial Plan: Go Packages and Modules**

### **Video Metadata**
- **Title:** Go Packages and Modules: Organization and Visibility
- **Duration Target:** 35-45 minutes
- **Difficulty:** Intermediate
- **Prerequisites:** Go Basics

---

## **Video Structure**

### **1. Introduction (3-4 min)**
- Packages vs modules
- Go module system overview
- Why organization matters
- Preview: Building a well-structured project

---

### **2. Package Basics (6-7 min)**

**Topics:**
- Package declaration
- Imports
- Exported vs unexported
- Package initialization

**Code Examples:**
```go
// math/calculator.go
package math  // Package declaration

import (
    "fmt"       // Standard library
    "strings"   // Multiple imports

    "github.com/user/project/internal/util"  // Project import
)

// Exported (capital letter) - accessible from other packages
func Add(a, b int) int {
    return a + b
}

// Unexported (lowercase) - only this package
func validateInput(n int) bool {
    return n >= 0
}

// Exported constant
const MaxValue = 1000

// Unexported constant
const defaultPrecision = 2

// Exported type
type Calculator struct {
    precision int
}

// Unexported type
type operation func(int, int) int

// Package-level variables
var (
    ErrOverflow = errors.New("overflow")   // Exported
    cache       = make(map[string]int)     // Unexported
)

// init() runs when package is imported
func init() {
    fmt.Println("math package initialized")
    // Setup, register, etc.
}
```

**Import rules:**
```go
// Standard format
import "fmt"

// Grouped imports
import (
    "fmt"
    "os"
)

// Alias
import (
    f "fmt"  // Use as f.Println()
)

// Blank import (side effects only)
import (
    _ "github.com/lib/pq"  // Registers postgres driver
)

// Dot import (avoid - pollutes namespace)
import (
    . "fmt"  // Use Println() instead of fmt.Println()
)
```

---

### **3. Go Modules (7-8 min)**

**Topics:**
- Module initialization
- go.mod file
- Dependencies
- Versioning

**Code Examples:**
```bash
# Initialize a new module
go mod init github.com/username/project

# Add a dependency
go get github.com/spf13/cobra@latest
go get github.com/spf13/cobra@v1.7.0

# Update dependencies
go get -u ./...

# Remove unused dependencies
go mod tidy

# Download dependencies
go mod download

# Verify dependencies
go mod verify

# View dependency graph
go mod graph
```

**go.mod file:**
```go
module github.com/username/project

go 1.21

require (
    github.com/spf13/cobra v1.7.0
    github.com/spf13/viper v1.16.0
)

require (
    // Indirect dependencies (transitive)
    github.com/fsnotify/fsnotify v1.6.0 // indirect
)

// Replace for local development
replace github.com/username/otherproject => ../otherproject

// Exclude problematic versions
exclude github.com/old/package v1.0.0
```

**go.sum file:**
- Cryptographic checksums
- Ensures reproducible builds
- Commit to version control

---

### **4. Project Layout (8-10 min)**

**Standard layout:**
```
project/
├── cmd/                    # Main applications
│   ├── api/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/               # Private packages
│   ├── auth/
│   ├── database/
│   └── handlers/
├── pkg/                    # Public packages (library)
│   ├── client/
│   └── models/
├── api/                    # API definitions (OpenAPI, proto)
├── web/                    # Web assets
├── configs/                # Configuration files
├── scripts/                # Build/deploy scripts
├── test/                   # Additional test data
├── docs/                   # Documentation
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

**Key directories:**
```go
// cmd/ - Entry points
// cmd/api/main.go
package main

import "github.com/username/project/internal/app"

func main() {
    app.Run()
}

// internal/ - Private to this module
// internal/database/db.go
package database

// Can only be imported by github.com/username/project/...
// NOT importable by other modules!

// pkg/ - Public library code
// pkg/client/client.go
package client

// Importable by anyone
type Client struct {
    baseURL string
}
```

---

### **5. Visibility Rules (5-6 min)**

**Code Examples:**
```go
// Exported vs Unexported

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

// From another package:
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

**internal/ enforcement:**
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

---

### **6. Package Design Principles (5-6 min)**

**Guidelines:**
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

---

### **7. Practical Example: Well-Organized Project (8-10 min)**

```
bookstore/
├── cmd/
│   └── bookstore/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   ├── book.go
│   │   └── author.go
│   ├── repository/
│   │   ├── book.go
│   │   └── author.go
│   └── service/
│       ├── book.go
│       └── author.go
├── pkg/
│   └── api/
│       └── client.go
├── go.mod
└── go.sum
```

```go
// cmd/bookstore/main.go
package main

import (
    "log"
    "os"

    "github.com/user/bookstore/internal/app"
    "github.com/user/bookstore/internal/config"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    application, err := app.New(cfg)
    if err != nil {
        log.Fatal(err)
    }

    if err := application.Run(); err != nil {
        log.Fatal(err)
    }
}

// internal/config/config.go
package config

type Config struct {
    Port        int
    DatabaseURL string
    LogLevel    string
}

func Load() (*Config, error) {
    return &Config{
        Port:        getEnvInt("PORT", 8080),
        DatabaseURL: os.Getenv("DATABASE_URL"),
        LogLevel:    getEnv("LOG_LEVEL", "info"),
    }, nil
}

// internal/app/app.go
package app

import (
    "github.com/user/bookstore/internal/config"
    "github.com/user/bookstore/internal/handler"
    "github.com/user/bookstore/internal/repository"
    "github.com/user/bookstore/internal/service"
)

type App struct {
    config *config.Config
    server *http.Server
}

func New(cfg *config.Config) (*App, error) {
    // Initialize layers
    db, err := sql.Open("postgres", cfg.DatabaseURL)
    if err != nil {
        return nil, err
    }

    // Repositories
    bookRepo := repository.NewBookRepository(db)
    authorRepo := repository.NewAuthorRepository(db)

    // Services
    bookSvc := service.NewBookService(bookRepo)
    authorSvc := service.NewAuthorService(authorRepo)

    // Handlers
    bookHandler := handler.NewBookHandler(bookSvc)
    authorHandler := handler.NewAuthorHandler(authorSvc)

    // Router
    mux := http.NewServeMux()
    bookHandler.Register(mux)
    authorHandler.Register(mux)

    return &App{
        config: cfg,
        server: &http.Server{
            Addr:    fmt.Sprintf(":%d", cfg.Port),
            Handler: mux,
        },
    }, nil
}

func (a *App) Run() error {
    return a.server.ListenAndServe()
}

// internal/service/book.go
package service

type BookService struct {
    repo BookRepository
}

// Interface defined where used
type BookRepository interface {
    FindByID(id int) (*Book, error)
    Save(book *Book) error
}

func NewBookService(repo BookRepository) *BookService {
    return &BookService{repo: repo}
}

// internal/repository/book.go
package repository

type BookRepository struct {
    db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
    return &BookRepository{db: db}
}

func (r *BookRepository) FindByID(id int) (*Book, error) {
    // SQL query
}

// pkg/api/client.go - Public API client
package api

// Client for external consumers
type Client struct {
    baseURL    string
    httpClient *http.Client
}

func NewClient(baseURL string) *Client {
    return &Client{
        baseURL:    baseURL,
        httpClient: &http.Client{Timeout: 30 * time.Second},
    }
}

func (c *Client) GetBook(id int) (*Book, error) {
    // HTTP call
}
```

---

### **8. Common Patterns (4-5 min)**

```go
// 1. Option pattern for configuration
type ServerOption func(*Server)

func WithPort(port int) ServerOption {
    return func(s *Server) { s.port = port }
}

func NewServer(opts ...ServerOption) *Server {
    s := &Server{port: 8080}
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// 2. Registry pattern
package plugins

var registry = make(map[string]Plugin)

func Register(name string, p Plugin) {
    registry[name] = p
}

func Get(name string) Plugin {
    return registry[name]
}

// 3. Package-level errors
package user

var (
    ErrNotFound    = errors.New("user not found")
    ErrInvalidData = errors.New("invalid user data")
)

// 4. Package initialization order
// a.go: var A = initA()   // Called first (alphabetically)
// b.go: var B = initB()   // Called second
// Both init() functions run after var initialization
```

---

### **9. Wrap-up (2-3 min)**

**Key takeaways:**
- Use meaningful package names
- Leverage internal/ for private code
- Export only what's necessary
- Follow standard project layout
- Document packages properly

**Homework:**
1. Reorganize an existing project
2. Create a reusable library
3. Publish to pkg.go.dev
4. Review popular Go projects for patterns

---

## **Supplementary Materials**

**Cheat Sheet:**
```
# Module commands
go mod init <module>
go mod tidy
go get <package>@<version>

# Visibility
Exported:   CapitalLetter
Unexported: lowercase

# Project structure
cmd/          Entry points
internal/     Private packages
pkg/          Public library
```

**Resources:**
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

---

This tutorial provides comprehensive guidance on organizing Go code professionally.
