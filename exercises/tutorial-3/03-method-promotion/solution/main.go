package main

import "fmt"

type Person struct {
	Name  string
	Email string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

func (p Person) Contact() string {
	return fmt.Sprintf("Contact me at: %s", p.Email)
}

// Employee embeds Person - methods are promoted
type Employee struct {
	Person
	Department string
}

func main() {
	emp := Employee{
		Person: Person{
			Name:  "Alice",
			Email: "alice@company.com",
		},
		Department: "Engineering",
	}

	// Method promotion - call Person methods directly on Employee!
	fmt.Println(emp.Greet())
	fmt.Println(emp.Contact())
	fmt.Println("Department:", emp.Department)
}
