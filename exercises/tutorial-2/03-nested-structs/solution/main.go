package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Employee struct {
	Name    string
	Email   string
	Address Address
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s %s", a.Street, a.City, a.ZipCode)
}

func main() {
	emp := Employee{
		Name:  "John Smith",
		Email: "john@company.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "10001",
		},
	}

	fmt.Println("Employee:", emp.Name)
	fmt.Println("City:", emp.Address.City)
	fmt.Println("Full Address:", emp.Address.FullAddress())
}
