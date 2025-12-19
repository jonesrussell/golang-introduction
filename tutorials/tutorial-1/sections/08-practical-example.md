# Practical Example: Building a Simple Program

**Duration:** 5-6 minutes

## Build together:
A number guessing game or grade calculator

## Example: Grade Calculator

```go runnable
package main

import "fmt"

func main() {
    // Student information
    studentName := "Alex"
    scores := []int{85, 92, 78, 95, 88}
    
    // Calculate average
    total := 0
    for _, score := range scores {
        total += score
    }
    average := float64(total) / float64(len(scores))
    
    // Determine grade
    var grade string
    switch {
    case average >= 90:
        grade = "A"
    case average >= 80:
        grade = "B"
    case average >= 70:
        grade = "C"
    case average >= 60:
        grade = "D"
    default:
        grade = "F"
    }
    
    // Output results
    fmt.Printf("Student: %s\n", studentName)
    fmt.Printf("Average: %.2f\n", average)
    fmt.Printf("Grade: %s\n", grade)
    
    // Pass/Fail determination
    if average >= 60 {
        fmt.Println("Status: PASSED ✓")
    } else {
        fmt.Println("Status: FAILED ✗")
    }
}
```

## Walk through:
- Declare variables for student data
- Use for loop with range to calculate total
- Type conversion for float division
- Switch statement for grade assignment
- If statement for pass/fail
- Formatted output with Printf
