package main

import "fmt"

func main() {
	// Explicit type declaration
	var greeting string = "Welcome to Go"

	// Short declaration
	year := 2024

	// Constant
	const maxScore = 100

	// Block declaration
	var (
		firstName string = "Jane"
		lastName  string = "Doe"
	)

	fmt.Println("Greeting:", greeting)
	fmt.Println("Year:", year)
	fmt.Println("Max Score:", maxScore)
	fmt.Println("Name:", firstName, lastName)
}
