# Pointer Gotchas and Safety

**Duration:** 5-6 minutes

## Topics to cover:
- [Nil pointer dereference](https://go.dev/ref/spec#Address_operators)
- [Escaping to heap](https://go.dev/doc/faq#stack_or_heap)
- Returning pointers to local variables
- Pointer comparison

## Code Examples

```go
// GOTCHA 1: Nil pointer dereference
func getUser(id int) *User {
    if id <= 0 {
        return nil  // No user found
    }
    return &User{ID: id}
}

func main() {
    user := getUser(-1)
    // fmt.Println(user.Username)  // PANIC!

    // Safe pattern
    if user != nil {
        fmt.Println(user.Username)
    }

    // Or use comma-ok pattern (where applicable)
}

// GOTCHA 2: Returning pointer to loop variable
func getBadPointers() []*int {
    nums := []int{1, 2, 3}
    result := make([]*int, 0)

    for _, n := range nums {
        result = append(result, &n)  // BAD: all point to same address!
    }
    return result
}

func getGoodPointers() []*int {
    nums := []int{1, 2, 3}
    result := make([]*int, 0)

    for _, n := range nums {
        n := n  // Shadow with new variable
        result = append(result, &n)  // Each pointer is unique
    }
    return result
}

// GOTCHA 3: Uninitialized struct with pointer fields
type Container struct {
    Data *[]int
}

func main() {
    c := Container{}  // c.Data is nil
    // *c.Data = append(*c.Data, 1)  // PANIC!

    // Initialize first
    data := make([]int, 0)
    c.Data = &data
    *c.Data = append(*c.Data, 1)  // OK
}

// SAFE: Go allows returning pointer to local variable
func newUser() *User {
    u := User{ID: 1}  // Local variable
    return &u          // Go moves u to heap (escape analysis)
}
// This is SAFE in Go (unlike C!)
```

## Best Practices

```go
// GOOD: Check for nil before using
func processUser(u *User) error {
    if u == nil {
        return errors.New("user cannot be nil")
    }
    // Safe to use u
    return nil
}

// GOOD: Use constructor to ensure initialization
func NewUser(id int, name string) *User {
    return &User{
        ID:   id,
        Name: name,
    }
}

// GOOD: Document when nil is valid
// GetUser returns the user or nil if not found.
func GetUser(id int) *User {
    // ...
}
```

## Key teaching points:
- Always check for [nil](https://go.dev/ref/spec#The_zero_value) before dereferencing
- Loop variable address gotcha (shadowing fix)
- Go's [escape analysis](https://go.dev/doc/faq#stack_or_heap) makes returning local pointers safe
- Document when nil is a valid return value
- Initialize struct pointer fields
