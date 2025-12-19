package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Product struct {
	ID      int
	Name    string
	Price   float64
	InStock bool
}

func main() {
	person := Person{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       28,
	}

	product := Product{
		ID:      1,
		Name:    "Laptop",
		Price:   999.99,
		InStock: true,
	}

	fmt.Println("Person:", person)
	fmt.Println("Product:", product)
}
