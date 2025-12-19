package main

import "fmt"

type User struct {
	Username string
	Email    string
	Age      int
	IsActive bool
}

func main() {
	// Zero value - all fields get their zero values
	var zeroUser User

	// Full initialization with all fields
	fullUser := User{
		Username: "alice",
		Email:    "alice@example.com",
		Age:      25,
		IsActive: true,
	}

	// Partial initialization - unspecified fields get zero values
	partialUser := User{
		Username: "bob",
		Age:      30,
	}

	fmt.Println("Zero value:", zeroUser)
	fmt.Println("Full init:", fullUser)
	fmt.Println("Partial init:", partialUser)
}
