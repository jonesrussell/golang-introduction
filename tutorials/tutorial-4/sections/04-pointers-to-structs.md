# Pointers to Structs

**Duration:** 6-7 minutes

## Topics to cover:
- Creating [pointers to structs](https://go.dev/ref/spec#Pointer_types)
- [Automatic dereferencing](https://go.dev/ref/spec#Selectors) with `.` notation
- Go's simplified syntax
- The [`new()` function](https://go.dev/ref/spec#Allocation)
- When to use `&Struct{}` vs `new()`

## Code Examples

```go
type User struct {
    ID       int
    Username string
    Email    string
}

func main() {
    // Method 1: Create struct, then get pointer
    u1 := User{ID: 1, Username: "alice", Email: "alice@example.com"}
    ptr1 := &u1

    // Method 2: Create pointer directly (most common)
    ptr2 := &User{
        ID:       2,
        Username: "bob",
        Email:    "bob@example.com",
    }

    // Method 3: Using new() - all fields are zero values
    ptr3 := new(User)
    ptr3.ID = 3
    ptr3.Username = "charlie"

    // Go automatically dereferences struct pointers!
    // These are equivalent:
    fmt.Println(ptr2.Username)      // bob
    fmt.Println((*ptr2).Username)   // bob (explicit dereference)

    // This is why Go doesn't need -> operator
    // In C: ptr->Username
    // In Go: ptr.Username (automatic dereference)

    // Modifying through pointer
    updateUser(ptr2)
    fmt.Println(ptr2.Email)  // updated@example.com
}

func updateUser(u *User) {
    u.Email = "updated@example.com"  // Modifies original
}
```

## Comparing Initialization

```go
// &Type{} - preferred when you have initial values
user := &User{
    ID:       1,
    Username: "alice",
}

// new(Type) - returns pointer with zero values
user := new(User)
// Equivalent to:
user := &User{}

// Both return *User
```

## Key teaching points:
- [`&Struct{}`](https://go.dev/ref/spec#Address_operators) is idiomatic for struct pointers with values
- [`new()`](https://go.dev/ref/spec#Allocation) returns pointer with zero values
- Go [auto-dereferences](https://go.dev/ref/spec#Selectors) struct pointers (no `->` needed)
- Pointer fields accessed same as value fields
- Constructor functions typically return `*Type`
