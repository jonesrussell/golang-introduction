# Performance Tips

**Duration:** 3-4 minutes

## 1. Use Logger not SugaredLogger in Hot Paths

```go
logger.Info("fast", zap.Int("key", 123))  // Faster
sugar.Infow("slower", "key", 123)          // Allocation for interface{}
```

## 2. Check Level Before Expensive Operations

```go
if logger.Core().Enabled(zap.DebugLevel) {
    logger.Debug("expensive", zap.Any("data", expensiveCompute()))
}
```

## 3. Pre-Build Loggers with Common Fields

```go
requestLogger := logger.With(
    zap.String("requestID", id),
    zap.String("userID", userID),
)
// Reuse requestLogger for the entire request
```

## 4. Use Sampling in Production

```go
config := zap.NewProductionConfig()
config.Sampling = &zap.SamplingConfig{
    Initial:    100,   // Log first 100 per second
    Thereafter: 100,   // Then sample 1 in 100
}
```
