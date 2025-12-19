# Best Practices Summary

**Duration:** 3-4 minutes

## Guidelines

```go
// 1. Accept interfaces, return structs
func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

// 2. Define interfaces where they're used
type service struct {
    repo userRepository  // Interface defined in this package
}

// 3. Keep interfaces small
type Saver interface {
    Save(v interface{}) error
}

// 4. Use constructor injection
func NewService(dep1 Dep1, dep2 Dep2) *Service

// 5. Provide no-op defaults for optional deps
func NewService(repo Repository, opts ...Option) *Service

// 6. Don't inject everything
// Inject: External services, databases, configuration
// Don't inject: Utility functions, time.Now(), etc.

// 7. Use context for request-scoped values
func (s *Service) DoThing(ctx context.Context) error

// 8. Test with mocks, not real implementations
repo := &MockRepository{}
service := NewService(repo)
```
