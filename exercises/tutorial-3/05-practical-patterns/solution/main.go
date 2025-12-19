package main

import "fmt"

type BaseLogger struct {
	Level string
}

func (l BaseLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.Level, message)
}

// FileLogger embeds BaseLogger
type FileLogger struct {
	BaseLogger
	Filename string
}

func (f FileLogger) WriteToFile(message string) {
	fmt.Printf("Writing to %s: %s\n", f.Filename, message)
}

// ConsoleLogger embeds BaseLogger but overrides Log()
type ConsoleLogger struct {
	BaseLogger
}

// Override Log() - this shadows the embedded method
func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[Console][%s] %s\n", c.Level, message)
}

func main() {
	fileLogger := FileLogger{
		BaseLogger: BaseLogger{Level: "DEBUG"},
		Filename:   "app.log",
	}

	consoleLogger := ConsoleLogger{
		BaseLogger: BaseLogger{Level: "WARN"},
	}

	infoLogger := BaseLogger{Level: "INFO"}

	// BaseLogger uses its own Log()
	infoLogger.Log("Starting application")

	// FileLogger uses promoted Log() from BaseLogger
	fileLogger.Log("Writing to " + fileLogger.Filename)

	// ConsoleLogger uses its own Log() (overrides BaseLogger)
	consoleLogger.Log("User not found")
}
