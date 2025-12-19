# Methods on Structs

**Duration:** 8-10 minutes

## Topics to cover:
- Method definition syntax
- Value receivers vs pointer receivers
- When to use each receiver type
- Method chaining
- Methods vs functions

## Code Examples

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver (does NOT modify original)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with value receiver
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Method with pointer receiver (CAN modify original)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Method with pointer receiver for consistency
func (r *Rectangle) String() string {
    return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

// Usage
rect := Rectangle{Width: 10, Height: 5}

// Value receiver methods
area := rect.Area()           // 50
perimeter := rect.Perimeter() // 30

// Pointer receiver method - modifies original
rect.Scale(2)
fmt.Println(rect.Width)  // 20
fmt.Println(rect.Height) // 10

// Go allows calling pointer receiver methods on values (and vice versa)
rect2 := Rectangle{Width: 5, Height: 3}
rect2.Scale(3)  // Go automatically does (&rect2).Scale(3)

pRect := &Rectangle{Width: 7, Height: 4}
area2 := pRect.Area()  // Go automatically does (*pRect).Area()
```

## BankAccount Example

```go
type BankAccount struct {
    Owner   string
    Balance float64
}

// Constructor
func NewBankAccount(owner string, initialBalance float64) *BankAccount {
    return &BankAccount{
        Owner:   owner,
        Balance: initialBalance,
    }
}

// Pointer receiver - modifies state
func (ba *BankAccount) Deposit(amount float64) {
    if amount > 0 {
        ba.Balance += amount
    }
}

// Pointer receiver - modifies state
func (ba *BankAccount) Withdraw(amount float64) bool {
    if amount > 0 && amount <= ba.Balance {
        ba.Balance -= amount
        return true
    }
    return false
}

// Value receiver - just reads data
func (ba BankAccount) GetBalance() float64 {
    return ba.Balance
}

// Value receiver - formats output
func (ba BankAccount) String() string {
    return fmt.Sprintf("%s's account: $%.2f", ba.Owner, ba.Balance)
}

// Usage
account := NewBankAccount("Alice", 1000)
account.Deposit(500)
account.Withdraw(200)
fmt.Println(account)  // Alice's account: $1300.00
```

## Key teaching points:
- Value receivers: `func (r Rectangle)` - gets a copy, cannot modify original
- Pointer receivers: `func (r *Rectangle)` - gets pointer, can modify original
- Go automatically converts between values and pointers for method calls
- **When to use pointer receivers:**
  - Method needs to modify the receiver
  - Struct is large (avoid copying)
  - Consistency: if any method uses pointer receiver, all should
- **When to use value receivers:**
  - Method only reads data
  - Struct is small (a few fields)
  - You want immutability guarantees
- Convention: be consistent within a type (all pointer or all value)
