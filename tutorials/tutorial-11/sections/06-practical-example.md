# Practical Example: Application Logging

**Duration:** 8-10 minutes

```go
package main

import (
    "context"
    "encoding/json"
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
