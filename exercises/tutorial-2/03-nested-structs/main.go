package main

import "fmt"

// Exercise 3: Nested Structs
//
// Your task:
// 1. Define an Address struct with: Street, City, ZipCode
// 2. Define an Employee struct with: Name, Email, Address (nested)
// 3. Create an employee and access nested fields
// 4. Add a method FullAddress() to Address that returns formatted string
//
// Expected output:
//   Employee: John Smith
//   City: New York
//   Full Address: 123 Main St, New York 10001
//
// Run with: go run main.go

// TODO: Define Address struct

// TODO: Define Employee struct with nested Address

// TODO: Add FullAddress() method to Address
// func (a Address) FullAddress() string {
//     return fmt.Sprintf("%s, %s %s", a.Street, a.City, a.ZipCode)
// }

func main() {
	// TODO: Create an Employee with nested Address
	// emp := Employee{
	//     Name: "John Smith",
	//     Email: "john@company.com",
	//     Address: Address{...},
	// }

	// Uncomment when ready:
	// fmt.Println("Employee:", emp.Name)
	// fmt.Println("City:", emp.Address.City)
	// fmt.Println("Full Address:", emp.Address.FullAddress())

	_ = fmt.Println
}
