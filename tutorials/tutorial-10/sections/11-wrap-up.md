# Recap & Best Practices

**Duration:** 2-3 minutes

## Summary:
1. Use context only for cancellation and request-scoped values
2. Inject dependencies explicitly
3. Define interfaces at point of use
4. Design types to avoid nil
5. Use sentinel errors and custom error types
6. Always provide goroutine exit paths
7. Profile before optimizing
8. Keep mutex scope minimal
9. Initialize explicitly in main

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
