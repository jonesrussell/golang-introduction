# Practical Example: Well-Organized Project

**Duration:** 8-10 minutes

## Project Structure

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

## Entry Point

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
```

## Configuration

```go
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
```

## Application Setup

```go
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
```

## Service Layer

```go
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
```

## Public API Client

```go
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
