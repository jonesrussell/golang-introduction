package main

import "fmt"

// Exercise 6: Switch Statements
//
// Your task:
// 1. Write a switch that prints the type of day based on 'dayNum' (1-7)
//    - 1: "Monday - Start of week"
//    - 6, 7: "Weekend!"
//    - default: "Midweek"
// 2. Write a switch without expression to categorize 'temperature':
//    - < 0: "Freezing"
//    - 0-15: "Cold"
//    - 16-25: "Nice"
//    - > 25: "Hot"
//
// Expected output:
//   Day 3: Midweek
//   Temperature 22: Nice
//
// Run with: go run main.go

func main() {
	dayNum := 3
	temperature := 22

	// TODO: Switch on dayNum
	fmt.Print("Day ", dayNum, ": ")
	// switch dayNum {
	// case 1:
	//     ...
	// case 6, 7:
	//     ...
	// default:
	//     ...
	// }

	// TODO: Switch without expression for temperature ranges
	fmt.Print("Temperature ", temperature, ": ")
	// switch {
	// case temperature < 0:
	//     ...
	// case temperature <= 15:
	//     ...
	// ...
	// }

	_ = fmt.Println // Remove when you add your code
}
