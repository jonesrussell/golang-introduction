# Premature Optimization

**Duration:** 4-5 minutes

## The Anti-Pattern

```go
// BAD: Complex "optimization" without measurement
func processData(data []byte) {
    // "Optimized" with sync.Pool
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    buf.Reset()

    // Pre-allocated slice
    result := make([]byte, 0, len(data)*2)

    // Manual loop "faster than range"
    for i := 0; i < len(data); i++ {
        // ...
    }
}

// When this simple version works fine:
func processDataSimple(data []byte) {
    var buf bytes.Buffer
    buf.Write(data)
    // ...
}
```

## The Fix

```go
// GOOD: Write clear code first
func ProcessItems(items []Item) []Result {
    results := make([]Result, 0, len(items))
    for _, item := range items {
        result := process(item)
        results = append(results, result)
    }
    return results
}
```

## Optimization Process

```bash
# 1. Profile first
go test -bench . -cpuprofile cpu.out

# 2. Identify bottlenecks
go tool pprof cpu.out

# 3. Optimize only hot paths
# 4. Measure improvement
```

## Acceptable Early Optimizations:
- Pre-allocate when size is known
- Use `strings.Builder` for concatenation
- Choose appropriate data structure
