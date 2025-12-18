## **Video Tutorial Plan: Structured Logging with Zap**

### **Video Metadata**
- **Title:** Structured Logging with Zap: Production-Ready Logging Practices
- **Duration Target:** 30-40 minutes
- **Difficulty:** Intermediate to Advanced
- **Prerequisites:** Go Basics, Interfaces, Error Handling

---

## **Video Structure**

### **1. Introduction (3-4 min)**
- Why structured logging matters
- Printf vs structured logging
- Zap overview and performance
- Preview: Production logging setup

---

### **2. The Problem with Printf (4-5 min)**

**The Anti-Pattern:**
```go
// BAD: Unstructured logging
log.Printf("User %d logged in from %s at %s", userID, ip, time.Now())
log.Printf("Error: %v", err)
log.Printf("Request completed in %dms", elapsed)

// Problems:
// 1. Hard to parse programmatically
// 2. Inconsistent formats
// 3. Can't filter by fields
// 4. Poor for log aggregation
```

**Structured logging benefits:**
- Machine-parseable (JSON)
- Queryable fields
- Consistent format
- Easy aggregation
- Better performance

---

### **3. Zap Basics (6-7 min)**

**Topics:**
- Logger types (Logger vs SugaredLogger)
- Creating loggers
- Log levels
- Adding fields

**Code Examples:**
```go
import "go.uber.org/zap"

// Production logger (JSON, fast)
logger, _ := zap.NewProduction()
defer logger.Sync()

// Development logger (human-readable)
logger, _ := zap.NewDevelopment()

// Logger vs SugaredLogger
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

// Log levels
logger.Debug("debug message")   // Development only
logger.Info("info message")     // Normal operations
logger.Warn("warning message")  // Potential issues
logger.Error("error message")   // Errors
logger.Fatal("fatal message")   // Exits program
logger.Panic("panic message")   // Panics

// With fields (creates child logger)
userLogger := logger.With(
    zap.Int("userID", 123),
    zap.String("component", "auth"),
)
userLogger.Info("login successful")
userLogger.Info("password changed")
```

**Output:**
```json
{"level":"info","ts":1702900000,"caller":"main.go:15","msg":"user logged in","userID":123,"ip":"192.168.1.1"}
```

---

### **4. Custom Configuration (5-6 min)**

**Code Examples:**
```go
// Custom logger configuration
config := zap.Config{
    Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
    Development: false,
    Encoding:    "json",  // or "console"
    EncoderConfig: zapcore.EncoderConfig{
        TimeKey:        "timestamp",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "message",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.MillisDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    },
    OutputPaths:      []string{"stdout", "/var/log/app.log"},
    ErrorOutputPaths: []string{"stderr"},
}

logger, _ := config.Build()

// Dynamic log level
atomicLevel := zap.NewAtomicLevel()
atomicLevel.SetLevel(zap.InfoLevel)

// Change at runtime
atomicLevel.SetLevel(zap.DebugLevel)

// Environment-based configuration
func NewLogger(env string) (*zap.Logger, error) {
    switch env {
    case "production":
        return zap.NewProduction()
    case "development":
        return zap.NewDevelopment()
    default:
        return zap.NewNop(), nil  // No-op logger for testing
    }
}
```

---

### **5. Logging Best Practices (6-7 min)**

**Code Examples:**
```go
// 1. Use structured fields, not string formatting
// BAD:
logger.Info(fmt.Sprintf("user %d performed %s", userID, action))

// GOOD:
logger.Info("user action",
    zap.Int("userID", userID),
    zap.String("action", action),
)

// 2. Create contextual loggers
type UserService struct {
    logger *zap.Logger
}

func NewUserService(baseLogger *zap.Logger) *UserService {
    return &UserService{
        logger: baseLogger.With(zap.String("component", "user-service")),
    }
}

func (s *UserService) GetUser(id int) (*User, error) {
    log := s.logger.With(zap.Int("userID", id))
    log.Debug("fetching user")

    user, err := s.repo.Find(id)
    if err != nil {
        log.Error("failed to fetch user", zap.Error(err))
        return nil, err
    }

    log.Info("user fetched successfully")
    return user, nil
}

// 3. Log errors with context
func processOrder(orderID string) error {
    logger := baseLogger.With(zap.String("orderID", orderID))

    if err := validateOrder(orderID); err != nil {
        logger.Error("order validation failed",
            zap.Error(err),
            zap.String("stage", "validation"),
        )
        return fmt.Errorf("validation: %w", err)
    }

    if err := chargePayment(orderID); err != nil {
        logger.Error("payment failed",
            zap.Error(err),
            zap.String("stage", "payment"),
        )
        return fmt.Errorf("payment: %w", err)
    }

    logger.Info("order processed successfully")
    return nil
}

// 4. Request logging middleware
func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            requestID := uuid.New().String()

            reqLogger := logger.With(
                zap.String("requestID", requestID),
                zap.String("method", r.Method),
                zap.String("path", r.URL.Path),
                zap.String("remoteAddr", r.RemoteAddr),
            )

            reqLogger.Info("request started")

            // Wrap response writer to capture status
            wrapped := &responseWriter{ResponseWriter: w, status: 200}

            next.ServeHTTP(wrapped, r)

            reqLogger.Info("request completed",
                zap.Int("status", wrapped.status),
                zap.Duration("duration", time.Since(start)),
            )
        })
    }
}
```

---

### **6. Practical Example: Application Logging (8-10 min)**

```go
package main

import (
    "context"
    "net/http"
    "os"
    "os/signal"
    "time"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// Global logger (initialized once)
var logger *zap.Logger

func initLogger(env string) {
    var config zap.Config

    if env == "production" {
        config = zap.NewProductionConfig()
        config.EncoderConfig.TimeKey = "timestamp"
        config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    } else {
        config = zap.NewDevelopmentConfig()
        config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    }

    var err error
    logger, err = config.Build()
    if err != nil {
        panic(err)
    }
}

// Service with injected logger
type OrderService struct {
    logger *zap.Logger
    repo   OrderRepository
}

func NewOrderService(baseLogger *zap.Logger, repo OrderRepository) *OrderService {
    return &OrderService{
        logger: baseLogger.With(zap.String("service", "orders")),
        repo:   repo,
    }
}

func (s *OrderService) CreateOrder(ctx context.Context, order *Order) error {
    log := s.logger.With(
        zap.String("orderID", order.ID),
        zap.Int("customerID", order.CustomerID),
    )

    log.Info("creating order",
        zap.Int("itemCount", len(order.Items)),
        zap.Float64("total", order.Total),
    )

    start := time.Now()
    if err := s.repo.Save(ctx, order); err != nil {
        log.Error("failed to save order",
            zap.Error(err),
            zap.Duration("duration", time.Since(start)),
        )
        return err
    }

    log.Info("order created successfully",
        zap.Duration("duration", time.Since(start)),
    )
    return nil
}

// HTTP handler with request logging
type Handler struct {
    logger  *zap.Logger
    service *OrderService
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
    requestID := r.Header.Get("X-Request-ID")
    log := h.logger.With(zap.String("requestID", requestID))

    log.Debug("parsing request body")

    var order Order
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        log.Warn("invalid request body", zap.Error(err))
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    if err := h.service.CreateOrder(r.Context(), &order); err != nil {
        log.Error("failed to create order", zap.Error(err))
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    log.Info("order endpoint completed")
    w.WriteHeader(http.StatusCreated)
}

func main() {
    env := os.Getenv("APP_ENV")
    initLogger(env)
    defer logger.Sync()

    logger.Info("application starting",
        zap.String("env", env),
        zap.String("version", "1.0.0"),
    )

    // Setup services
    repo := NewOrderRepository(db)
    service := NewOrderService(logger, repo)
    handler := &Handler{logger: logger, service: service}

    // Setup routes
    mux := http.NewServeMux()
    mux.HandleFunc("/orders", handler.CreateOrder)

    server := &http.Server{
        Addr:    ":8080",
        Handler: LoggingMiddleware(logger)(mux),
    }

    // Graceful shutdown
    go func() {
        logger.Info("server starting", zap.String("addr", ":8080"))
        if err := server.ListenAndServe(); err != http.ErrServerClosed {
            logger.Fatal("server error", zap.Error(err))
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit

    logger.Info("shutting down gracefully")
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        logger.Error("shutdown error", zap.Error(err))
    }

    logger.Info("application stopped")
}
```

---

### **7. Performance Tips (3-4 min)**

```go
// 1. Use Logger not SugaredLogger in hot paths
logger.Info("fast", zap.Int("key", 123))  // Faster
sugar.Infow("slower", "key", 123)          // Allocation for interface{}

// 2. Check level before expensive operations
if logger.Core().Enabled(zap.DebugLevel) {
    logger.Debug("expensive", zap.Any("data", expensiveCompute()))
}

// 3. Pre-build loggers with common fields
requestLogger := logger.With(
    zap.String("requestID", id),
    zap.String("userID", userID),
)
// Reuse requestLogger for the entire request

// 4. Use sampling in production
config := zap.NewProductionConfig()
config.Sampling = &zap.SamplingConfig{
    Initial:    100,   // Log first 100 per second
    Thereafter: 100,   // Then sample 1 in 100
}
```

---

### **8. Wrap-up (2-3 min)**

**Key takeaways:**
- Use structured fields, not string formatting
- Create contextual child loggers
- Configure appropriately for environment
- Log errors with full context
- Use middleware for request logging

**Homework:**
1. Add structured logging to existing project
2. Create custom encoder for specific format
3. Implement log level API endpoint
4. Set up log aggregation (ELK, Loki)

---

## **Supplementary Materials**

**Zap Cheat Sheet:**
```go
// Logger creation
logger, _ := zap.NewProduction()
logger, _ := zap.NewDevelopment()
sugar := logger.Sugar()

// Logging
logger.Info("msg", zap.String("key", "val"))
logger.Error("msg", zap.Error(err))
sugar.Infow("msg", "key", "val")

// Child logger
child := logger.With(zap.String("ctx", "value"))

// Common fields
zap.String("key", "value")
zap.Int("key", 123)
zap.Error(err)
zap.Duration("key", time.Second)
zap.Time("key", time.Now())
zap.Any("key", obj)
```

---

This tutorial provides practical guidance on implementing production-ready logging with Zap.
