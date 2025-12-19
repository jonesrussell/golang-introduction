package main

import "fmt"

// Exercise 1: Basic Composition
//
// Your task:
// 1. Define an Engine struct with: Horsepower (int), Type (string)
// 2. Define a Car struct with: Brand, Model (strings), Engine (nested)
// 3. Add a Describe() method to Engine
// 4. Create a Car and access its engine properties
//
// Note: This uses explicit composition (named fields)
//
// Expected output:
//   Car: Toyota Camry
//   Engine: 203 HP V6
//   Horsepower: 203
//
// Run with: go run main.go

// TODO: Define Engine struct

// TODO: Add Describe() method to Engine
// func (e Engine) Describe() string {
//     return fmt.Sprintf("%d HP %s", e.Horsepower, e.Type)
// }

// TODO: Define Car struct with Engine field (explicit composition)

func main() {
	// TODO: Create a Car with nested Engine
	// car := Car{
	//     Brand: "Toyota",
	//     Model: "Camry",
	//     Engine: Engine{...},
	// }

	// Note: Must access Engine through field name
	// Uncomment when ready:
	// fmt.Println("Car:", car.Brand, car.Model)
	// fmt.Println("Engine:", car.Engine.Describe())
	// fmt.Println("Horsepower:", car.Engine.Horsepower)

	_ = fmt.Println
}
