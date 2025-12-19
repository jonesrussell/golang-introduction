# Logging Best Practices

**Duration:** 6-7 minutes

## 1. Use Structured Fields, Not String Formatting

```go
// BAD:
logger.Info(fmt.Sprintf("user %d performed %s", userID, action))

// GOOD:
logger.Info("user action",
    zap.Int("userID", userID),
    zap.String("action", action),
)
```

## 2. Create Contextual Loggers

```go
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
```

## 3. Log Errors with Context

```go
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
```

## 4. Request Logging Middleware

```go
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
