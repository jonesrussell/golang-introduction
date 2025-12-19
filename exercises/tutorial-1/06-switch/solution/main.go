package main

import "fmt"

func main() {
	dayNum := 3
	temperature := 22

	// Switch on dayNum
	fmt.Print("Day ", dayNum, ": ")
	switch dayNum {
	case 1:
		fmt.Println("Monday - Start of week")
	case 6, 7:
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek")
	}

	// Switch without expression for ranges
	fmt.Print("Temperature ", temperature, ": ")
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature <= 15:
		fmt.Println("Cold")
	case temperature <= 25:
		fmt.Println("Nice")
	default:
		fmt.Println("Hot")
	}
}
