package main

import "fmt"

// Exercise 6: Constructor Pattern
//
// Your task:
// 1. Define a BankAccount struct with Owner (string) and Balance (float64)
// 2. Create a NewBankAccount(owner string, initial float64) constructor
//    - Return *BankAccount
//    - Set initial balance (minimum 0)
// 3. Add Deposit(amount float64) method
// 4. Add Withdraw(amount float64) bool method (return false if insufficient)
// 5. Add String() method for display
//
// Expected output:
//   Created: Alice's account: $100.00
//   After deposit: Alice's account: $150.00
//   Withdraw $30: true
//   After withdraw: Alice's account: $120.00
//   Withdraw $200: false
//
// Run with: go run main.go

// TODO: Define BankAccount struct

// TODO: Create NewBankAccount constructor
// func NewBankAccount(owner string, initial float64) *BankAccount {
//     ...
// }

// TODO: Add Deposit method (pointer receiver)

// TODO: Add Withdraw method (pointer receiver, returns bool)

// TODO: Add String method

func main() {
	// TODO: Create account using constructor
	// account := NewBankAccount("Alice", 100)

	// Uncomment when ready:
	// fmt.Println("Created:", account)
	//
	// account.Deposit(50)
	// fmt.Println("After deposit:", account)
	//
	// success := account.Withdraw(30)
	// fmt.Println("Withdraw $30:", success)
	// fmt.Println("After withdraw:", account)
	//
	// success = account.Withdraw(200)
	// fmt.Println("Withdraw $200:", success)

	_ = fmt.Println
}
