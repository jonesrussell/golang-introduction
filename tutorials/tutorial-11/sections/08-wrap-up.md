# Wrap-up

**Duration:** 2-3 minutes

## Key Takeaways:
- Use structured fields, not string formatting
- Create contextual child loggers
- Configure appropriately for environment
- Log errors with full context
- Use middleware for request logging

## Homework:
1. Add structured logging to existing project
2. Create custom encoder for specific format
3. Implement log level API endpoint
4. Set up log aggregation (ELK, Loki)

## Zap Cheat Sheet

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
