# Custom Configuration

**Duration:** 5-6 minutes

## Custom Logger Configuration

```go
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
```

## Dynamic Log Level

```go
atomicLevel := zap.NewAtomicLevel()
atomicLevel.SetLevel(zap.InfoLevel)

// Change at runtime
atomicLevel.SetLevel(zap.DebugLevel)
```

## Environment-Based Configuration

```go
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
