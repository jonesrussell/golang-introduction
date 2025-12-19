# Common Patterns

**Duration:** 4-5 minutes

## 1. Option Pattern for Configuration

```go
type ServerOption func(*Server)

func WithPort(port int) ServerOption {
    return func(s *Server) { s.port = port }
}

func NewServer(opts ...ServerOption) *Server {
    s := &Server{port: 8080}
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

## 2. Registry Pattern

```go
package plugins

var registry = make(map[string]Plugin)

func Register(name string, p Plugin) {
    registry[name] = p
}

func Get(name string) Plugin {
    return registry[name]
}
```

## 3. Package-Level Errors

```go
package user

var (
    ErrNotFound    = errors.New("user not found")
    ErrInvalidData = errors.New("invalid user data")
)
```

## 4. Package Initialization Order

```go
// a.go: var A = initA()   // Called first (alphabetically)
// b.go: var B = initB()   // Called second
// Both init() functions run after var initialization
```
