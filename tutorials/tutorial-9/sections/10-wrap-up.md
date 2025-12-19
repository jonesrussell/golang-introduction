# Next Steps & Wrap-up

**Duration:** 2-3 minutes

## Recap:
- Why DI matters
- [Constructor injection](https://go.dev/doc/effective_go#composite_literals)
- Testing with mocks
- [Interface design](https://go.dev/ref/spec#Interface_types)
- Functional options
- Complete application example

## Homework:
1. Refactor existing code to use DI
2. Add caching layer with DI
3. Build testable HTTP handlers
4. Implement with Google Wire

## Cheat Sheet

```
Constructor Injection:
  func NewService(dep Interface) *Service

Functional Options:
  type Option func(*Service)
  func NewService(opts ...Option) *Service

Interface Location:
  Consumer defines interface it needs

Testing:
  Mock implements same interface
  Test success and error paths
```

## Resources:
- [Effective Go - Interfaces](https://go.dev/doc/effective_go#interfaces_and_types): go.dev/doc/effective_go#interfaces_and_types
- [Go Blog: "Testable Examples in Go"](https://go.dev/blog/examples): go.dev/blog/examples
- [Go Testing Package](https://pkg.go.dev/testing): pkg.go.dev/testing
