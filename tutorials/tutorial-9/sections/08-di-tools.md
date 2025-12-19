# Dependency Injection Tools

**Duration:** 4-5 minutes

## Topics to cover:
- Google Wire
- Manual vs generated DI
- When to use tools

## Google Wire - Compile-Time DI

```go
// wire.go (build tag: wireinject)
//go:build wireinject

package main

import "github.com/google/wire"

func InitializeApp() (*App, error) {
    wire.Build(
        NewConfig,
        NewDatabase,
        NewUserRepository,
        NewUserService,
        NewHTTPServer,
        NewApp,
    )
    return nil, nil
}
```

## Providers

```go
func NewDatabase(cfg *Config) (*sql.DB, error) {
    return sql.Open("postgres", cfg.DatabaseURL)
}

func NewUserRepository(db *sql.DB) *PostgresUserRepository {
    return &PostgresUserRepository{db: db}
}

func NewUserService(repo UserRepository, logger Logger) *UserService {
    return &UserService{repo: repo, logger: logger}
}

// Wire generates wire_gen.go with actual wiring code
```

## When to Use DI Tools:
- Large applications with many dependencies
- Complex dependency graphs
- Want compile-time safety

## When Manual DI Is Fine:
- Small to medium applications
- Simple dependency graphs
- Prefer explicit over magic
