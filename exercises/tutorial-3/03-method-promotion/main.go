package main

import "fmt"

// Exercise 3: Method Promotion
//
// Your task:
// 1. Define a Person struct with Name, Email
// 2. Add a Greet() method to Person
// 3. Add a Contact() method to Person
// 4. Define Employee that embeds Person, adds Department
// 5. Call Person methods directly on Employee (method promotion)
//
// Expected output:
//   Hello, I'm Alice
//   Contact me at: alice@company.com
//   Department: Engineering
//
// Run with: go run main.go

// TODO: Define Person struct

// TODO: Add Greet() method to Person
// func (p Person) Greet() string {
//     return fmt.Sprintf("Hello, I'm %s", p.Name)
// }

// TODO: Add Contact() method to Person

// TODO: Define Employee that EMBEDS Person

func main() {
	// TODO: Create an Employee
	// emp := Employee{
	//     Person: Person{
	//         Name:  "Alice",
	//         Email: "alice@company.com",
	//     },
	//     Department: "Engineering",
	// }

	// Uncomment when ready:
	// Method promotion - call Person methods directly on Employee!
	// fmt.Println(emp.Greet())
	// fmt.Println(emp.Contact())
	// fmt.Println("Department:", emp.Department)

	_ = fmt.Println
}
