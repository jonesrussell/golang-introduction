package main

import "fmt"

func main() {
	// Classic for loop
	fmt.Print("Counting up: ")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// While-style loop
	fmt.Print("Countdown: ")
	count := 3
	for count > 0 {
		fmt.Print(count, " ")
		count--
	}
	fmt.Println("Liftoff!")

	// Range over string
	for index, char := range "Go!" {
		fmt.Printf("Index %d: %c\n", index, char)
	}

	// Even numbers with continue
	fmt.Print("Even numbers: ")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
}
