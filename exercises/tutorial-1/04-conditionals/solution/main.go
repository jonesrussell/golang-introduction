package main

import "fmt"

func main() {
	age := 25
	score := 85

	// Check if adult
	if age >= 18 {
		fmt.Printf("Age %d: Adult\n", age)
	} else {
		fmt.Printf("Age %d: Minor\n", age)
	}

	// Determine grade
	var grade string
	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else {
		grade = "F"
	}
	fmt.Printf("Score %d: Grade %s\n", score, grade)

	// If with initialization
	if num := 42; num > 0 {
		fmt.Printf("Number %d is positive\n", num)
	} else {
		fmt.Printf("Number %d is not positive\n", num)
	}
}
