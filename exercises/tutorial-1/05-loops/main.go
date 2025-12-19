package main

import "fmt"

// Exercise 5: Loops
//
// Your task:
// 1. Print numbers 1 to 5 using a classic for loop
// 2. Use a while-style loop to count down from 3 to 1
// 3. Use range to print each character of "Go!" with its index
// 4. Print only even numbers from 1-10 (use continue for odd numbers)
//
// Expected output:
//   Counting up: 1 2 3 4 5
//   Countdown: 3 2 1 Liftoff!
//   Index 0: G
//   Index 1: o
//   Index 2: !
//   Even numbers: 2 4 6 8 10
//
// Run with: go run main.go

func main() {
	// TODO: Classic for loop - print 1 to 5
	fmt.Print("Counting up: ")
	// for i := 1; i <= 5; i++ { ... }
	fmt.Println()

	// TODO: While-style loop - countdown from 3
	fmt.Print("Countdown: ")
	// count := 3
	// for count > 0 { ... }
	fmt.Println("Liftoff!")

	// TODO: Range over string "Go!"
	// for index, char := range "Go!" { ... }

	// TODO: Print only even numbers 1-10, skip odd with continue
	fmt.Print("Even numbers: ")
	// for i := 1; i <= 10; i++ {
	//     if i % 2 != 0 {
	//         continue
	//     }
	//     ...
	// }
	fmt.Println()

	_ = fmt.Println // Remove when you add your code
}
