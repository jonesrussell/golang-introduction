# Embedding and Interfaces

**Duration:** 7-8 minutes

## Topics to cover:
- How embedded structs [satisfy interfaces](https://go.dev/ref/spec#Interface_types)
- [Interface composition](https://go.dev/ref/spec#Interface_types) via embedding
- Wrapper pattern with embedding
- Overriding embedded methods

## Code Examples

```go
// Interface satisfaction through embedding

type Notifier interface {
    Notify(message string)
}

type EmailNotifier struct {
    Email string
}

func (e EmailNotifier) Notify(message string) {
    fmt.Printf("Email to %s: %s\n", e.Email, message)
}

// User embeds EmailNotifier and satisfies Notifier interface
type User struct {
    Name string
    EmailNotifier  // Embedded - promotes Notify method
}

// User automatically satisfies Notifier interface
func SendNotification(n Notifier, msg string) {
    n.Notify(msg)
}

user := User{
    Name: "Alice",
    EmailNotifier: EmailNotifier{
        Email: "alice@example.com",
    },
}

SendNotification(user, "Welcome!")  // Works! User satisfies Notifier
```

## Overriding Embedded Methods

```go
// Overriding embedded methods
type AdminUser struct {
    User
    Permissions []string
}

// Override the Notify method
func (a AdminUser) Notify(message string) {
    fmt.Println("[ADMIN NOTIFICATION]")
    a.User.Notify(message)  // Can still call embedded version
    fmt.Printf("Admin %s has been notified\n", a.Name)
}

admin := AdminUser{
    User: User{
        Name: "Bob",
        EmailNotifier: EmailNotifier{
            Email: "bob@example.com",
        },
    },
    Permissions: []string{"admin"},
}

SendNotification(admin, "System alert")  // Uses overridden method
```

## Interface Embedding

```go
// Interface embedding
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Compose interfaces by embedding
type ReadWriter interface {
    Reader  // Embedded interface
    Writer  // Embedded interface
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// This is how io.ReadWriteCloser is defined in standard library!

// Any type that has Read, Write, and Close methods satisfies ReadWriteCloser
type File struct {
    name string
    data []byte
}

func (f *File) Read(p []byte) (int, error) {
    // Implementation
    return 0, nil
}

func (f *File) Write(p []byte) (int, error) {
    // Implementation
    return 0, nil
}

func (f *File) Close() error {
    // Implementation
    return nil
}

var rwc ReadWriteCloser = &File{name: "test.txt"}
```

## Key teaching points:
- Embedded struct methods count toward [interface satisfaction](https://go.dev/ref/spec#Interface_types)
- Can override embedded methods by defining same method on outer struct
- Can still call original embedded method explicitly
- [Interface embedding](https://go.dev/ref/spec#Interface_types) creates composite interfaces
- This is how [standard library composes interfaces](https://pkg.go.dev/io#ReadWriteCloser) (io.Reader, io.Writer, etc.)
- Embedding provides delegation pattern
