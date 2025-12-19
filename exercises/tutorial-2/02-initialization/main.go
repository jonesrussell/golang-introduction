package main

import "fmt"

// Exercise 2: Struct Initialization
//
// Your task:
// 1. Create a User struct with: Username, Email, Age, IsActive
// 2. Create instances using different initialization methods:
//    - Zero value (var declaration)
//    - Literal with all fields
//    - Literal with named fields (partial)
// 3. Print each to see the differences
//
// Expected output:
//   Zero value: { 0 false}
//   Full init: {alice alice@example.com 25 true}
//   Partial init: {bob  30 false}
//
// Run with: go run main.go

// TODO: Define User struct with Username, Email, Age, IsActive

func main() {
	// TODO: Create zero-value User using var
	// var zeroUser User

	// TODO: Create fully initialized User
	// fullUser := User{...}

	// TODO: Create partially initialized User (only Username and Age)
	// partialUser := User{...}

	// Uncomment when ready:
	// fmt.Println("Zero value:", zeroUser)
	// fmt.Println("Full init:", fullUser)
	// fmt.Println("Partial init:", partialUser)

	_ = fmt.Println
}
