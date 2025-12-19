# Next Steps & Wrap-up

**Duration:** 2-3 minutes

## Recap:
- Why DI matters
- Constructor injection
- Testing with mocks
- Interface design
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
