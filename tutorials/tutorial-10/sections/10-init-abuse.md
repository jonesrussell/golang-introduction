# Init Function Abuse

**Duration:** 3-4 minutes

## The Anti-Pattern

```go
// BAD: Complex init with errors
func init() {
    db, err := sql.Open("postgres", os.Getenv("DB_URL"))
    if err != nil {
        panic(err)  // Crashes on startup
    }
    globalDB = db

    config, err := loadConfig()
    if err != nil {
        panic(err)
    }
    globalConfig = config
}
```

## The Fix

```go
// GOOD: Explicit initialization in main
func main() {
    config, err := loadConfig()
    if err != nil {
        log.Fatalf("loading config: %v", err)
    }

    db, err := setupDatabase(config)
    if err != nil {
        log.Fatalf("connecting to database: %v", err)
    }
    defer db.Close()

    server := NewServer(config, db)
    server.Run()
}
```

## Acceptable Init Uses:
- Register drivers: `sql.Register`, `http.Handle`
- Compile regexes
- Set package-level computed constants

## Key teaching points:
- Avoid complex logic in [`init()`](https://go.dev/ref/spec#Package_initialization)
- Initialize explicitly in [`main()`](https://go.dev/ref/spec#Program_initialization_and_execution)
- Use [`init()`](https://go.dev/ref/spec#Package_initialization) only for simple, error-free setup
- Prefer explicit initialization for better error handling
