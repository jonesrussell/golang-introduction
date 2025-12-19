# Recap & Best Practices

**Duration:** 2-3 minutes

## Summary:
1. Use [context](https://pkg.go.dev/context) only for cancellation and request-scoped values
2. Inject dependencies explicitly
3. Define [interfaces](https://go.dev/ref/spec#Interface_types) at point of use
4. Design types to avoid [nil](https://go.dev/ref/spec#The_zero_value)
5. Use [sentinel errors](https://pkg.go.dev/errors#New) and custom error types
6. Always provide [goroutine](https://go.dev/ref/spec#Go_statements) exit paths
7. [Profile before optimizing](https://go.dev/doc/diagnostics#profiling)
8. Keep [mutex](https://pkg.go.dev/sync#Mutex) scope minimal
9. Initialize explicitly in [`main()`](https://go.dev/ref/spec#Program_initialization_and_execution)

## Anti-Pattern Checklist

```
[ ] Context.Value for dependencies?
[ ] Global mutable state?
[ ] Interface for single implementation?
[ ] String matching for errors?
[ ] Goroutines without exit path?
[ ] Mutex with value receiver?
[ ] Complex logic in init()?
```

## Resources:
- [Effective Go](https://go.dev/doc/effective_go): go.dev/doc/effective_go
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments): github.com/golang/go/wiki/CodeReviewComments
- [Go FAQ](https://go.dev/doc/faq): go.dev/doc/faq
