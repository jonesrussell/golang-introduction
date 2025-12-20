package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

// Constructor - returns pointer to new BankAccount
func NewBankAccount(owner string, initial float64) *BankAccount {
	balance := initial
	if balance < 0 {
		balance = 0
	}
	return &BankAccount{
		Owner:   owner,
		Balance: balance,
	}
}

// Pointer receiver - modifies balance
func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.Balance += amount
	}
}

// Pointer receiver - modifies balance, returns success
func (ba *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && amount <= ba.Balance {
		ba.Balance -= amount
		return true
	}
	return false
}

// Pointer receiver - formats output
func (ba *BankAccount) String() string {
	return fmt.Sprintf("%s's account: $%.2f", ba.Owner, ba.Balance)
}

func main() {
	account := NewBankAccount("Alice", 100)

	fmt.Println("Created:", account)

	account.Deposit(50)
	fmt.Println("After deposit:", account)

	success := account.Withdraw(30)
	fmt.Println("Withdraw $30:", success)
	fmt.Println("After withdraw:", account)

	success = account.Withdraw(200)
	fmt.Println("Withdraw $200:", success)
}
