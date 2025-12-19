package main

import "fmt"

type Engine struct {
	Horsepower int
	Type       string
}

func (e Engine) Describe() string {
	return fmt.Sprintf("%d HP %s", e.Horsepower, e.Type)
}

type Car struct {
	Brand  string
	Model  string
	Engine Engine // Explicit composition - named field
}

func main() {
	car := Car{
		Brand: "Toyota",
		Model: "Camry",
		Engine: Engine{
			Horsepower: 203,
			Type:       "V6",
		},
	}

	// Must access Engine through field name
	fmt.Println("Car:", car.Brand, car.Model)
	fmt.Println("Engine:", car.Engine.Describe())
	fmt.Println("Horsepower:", car.Engine.Horsepower)
}
