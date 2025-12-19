package main

import "fmt"

// Exercise 5: Practical Patterns
//
// Your task:
// Build a simple logging system using embedding:
//
// 1. Define a BaseLogger struct with Level (string)
// 2. Add Log(message string) method to BaseLogger
// 3. Define a FileLogger that embeds BaseLogger
//    - Add Filename field
//    - Add a WriteToFile(msg string) method
// 4. Define a ConsoleLogger that embeds BaseLogger
//    - Override Log() to add "[Console]" prefix
//
// Expected output:
//   [INFO] Starting application
//   [DEBUG] Writing to app.log
//   [Console][WARN] User not found
//
// Run with: go run main.go

// TODO: Define BaseLogger with Level field

// TODO: Add Log(message) method to BaseLogger
// func (l BaseLogger) Log(message string) {
//     fmt.Printf("[%s] %s\n", l.Level, message)
// }

// TODO: Define FileLogger embedding BaseLogger

// TODO: Add WriteToFile method to FileLogger

// TODO: Define ConsoleLogger embedding BaseLogger

// TODO: Override Log() in ConsoleLogger to add prefix
// func (c ConsoleLogger) Log(message string) {
//     fmt.Printf("[Console][%s] %s\n", c.Level, message)
// }

func main() {
	// TODO: Create FileLogger and ConsoleLogger
	// fileLogger := FileLogger{
	//     BaseLogger: BaseLogger{Level: "DEBUG"},
	//     Filename:   "app.log",
	// }
	//
	// consoleLogger := ConsoleLogger{
	//     BaseLogger: BaseLogger{Level: "WARN"},
	// }
	//
	// infoLogger := BaseLogger{Level: "INFO"}

	// Uncomment when ready:
	// infoLogger.Log("Starting application")
	// fileLogger.Log("Writing to " + fileLogger.Filename)
	// consoleLogger.Log("User not found")

	_ = fmt.Println
}
