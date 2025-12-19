# Zap Basics

**Duration:** 6-7 minutes

## Topics:
- Logger types (Logger vs SugaredLogger)
- Creating loggers
- Log levels
- Adding fields

## Code Examples

```go
import "go.uber.org/zap"

// Production logger (JSON, fast)
logger, _ := zap.NewProduction()
defer logger.Sync()

// Development logger (human-readable)
logger, _ := zap.NewDevelopment()
```

## Logger vs SugaredLogger

```go
// Logger: Strongly typed, fastest
logger.Info("user logged in",
    zap.Int("userID", 123),
    zap.String("ip", "192.168.1.1"),
)

// SugaredLogger: Printf-style, slightly slower
sugar := logger.Sugar()
sugar.Infow("user logged in",
    "userID", 123,
    "ip", "192.168.1.1",
)
sugar.Infof("User %d logged in", 123)
```

## Log Levels

```go
logger.Debug("debug message")   // Development only
logger.Info("info message")     // Normal operations
logger.Warn("warning message")  // Potential issues
logger.Error("error message")   // Errors
logger.Fatal("fatal message")   // Exits program
logger.Panic("panic message")   // Panics
```

## Child Loggers with Fields

```go
// With fields (creates child logger)
userLogger := logger.With(
    zap.Int("userID", 123),
    zap.String("component", "auth"),
)
userLogger.Info("login successful")
userLogger.Info("password changed")
```

## Output Example

```json
{"level":"info","ts":1702900000,"caller":"main.go:15","msg":"user logged in","userID":123,"ip":"192.168.1.1"}
```

## Key teaching points:
- [zap.Logger](https://pkg.go.dev/go.uber.org/zap#Logger) is type-safe and fastest
- [zap.SugaredLogger](https://pkg.go.dev/go.uber.org/zap#SugaredLogger) is more convenient but slower
- Use [child loggers](https://pkg.go.dev/go.uber.org/zap#Logger.With) for contextual logging
- Always call [`Sync()`](https://pkg.go.dev/go.uber.org/zap#Logger.Sync) before program exit
