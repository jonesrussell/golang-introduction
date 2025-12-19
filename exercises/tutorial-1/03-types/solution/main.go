package main

import "fmt"

func main() {
	quantity := 42
	price := 19.99

	// Must convert int to float64 for multiplication
	total := float64(quantity) * price

	firstName := "John"
	lastName := "Smith"
	fullName := firstName + " " + lastName

	receipt := `========
Thank you
for shopping!
========`

	fmt.Println("Quantity:", quantity)
	fmt.Println("Price:", price)
	fmt.Println("Total:", total)
	fmt.Println("Customer:", fullName)
	fmt.Println("Receipt:")
	fmt.Println(receipt)
}
