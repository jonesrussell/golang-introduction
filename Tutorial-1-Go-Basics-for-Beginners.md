## **Video Tutorial Plan: Go Basics for Beginners**

### **Video Metadata**
- **Title:** Go Basics: Variables, Types, and Control Flow for Beginners
- **Duration Target:** 25-35 minutes
- **Difficulty:** Beginner (no prior Go experience needed)
- **Prerequisites:** Basic programming concepts helpful but not required

---

## **Video Structure**

### **1. Introduction (2-3 min)**
- Welcome and what viewers will learn
- Why Go? Brief mention of simplicity, performance, and use cases
- Show the final code we'll build (simple program that demonstrates all concepts)
- Setup check: Go installed, editor ready (VS Code/Cursor recommended)

---

### **2. Hello World & Package Basics (3-4 min)**

**Topics to cover:**
- Creating `main.go`
- `package main` declaration
- `import "fmt"`
- `func main()` as entry point
- `fmt.Println()` for output
- Running with `go run main.go`

**Code Example:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

**Key teaching points:**
- Every Go file starts with a package declaration
- `main` package is special - it's executable
- Import standard library packages
- `main()` function is where execution begins

---

### **3. Variables & Declaration (6-7 min)**

**Topics to cover:**
- Variable declaration with `var`
- Type inference
- Short declaration (`:=`)
- Zero values
- Constants with `const`
- Multiple variable declaration

**Code Examples:**
```go
// Explicit type declaration
var name string = "Russell"
var age int = 30

// Type inference
var city = "Toronto"

// Short declaration (most common)
country := "Canada"

// Zero values
var count int        // 0
var isActive bool    // false
var message string   // ""

// Constants
const MaxRetries = 3
const Pi = 3.14159

// Multiple declaration
var (
    firstName string = "John"
    lastName  string = "Doe"
    score     int    = 95
)
```

**Key teaching points:**
- `:=` can only be used inside functions
- Go is statically typed but has type inference
- Unused variables are compilation errors (good for code quality!)
- Zero values prevent uninitialized variable bugs
- Constants must be compile-time values

---

### **4. Basic Types (5-6 min)**

**Topics to cover:**
- Numeric types: `int`, `int64`, `float64`
- String type
- Boolean type
- Type conversion (explicit only)
- String concatenation

**Code Examples:**
```go
// Numeric types
var count int = 42
var price float64 = 19.99
var distance int64 = 1000000

// Strings
message := "Learning Go"
multiLine := `This is a
multi-line string
using backticks`

// Booleans
isComplete := true
hasError := false

// Type conversion (explicit)
var x int = 10
var y float64 = float64(x)  // Must convert explicitly
// var z float64 = x  // This would be an error!

// String operations
firstName := "Jane"
lastName := "Smith"
fullName := firstName + " " + lastName
fmt.Printf("Name: %s, Length: %d\n", fullName, len(fullName))
```

**Key teaching points:**
- No implicit type conversion (prevents bugs)
- `int` vs `int64` - platform-dependent vs explicit size
- String concatenation with `+`
- `fmt.Printf` for formatted output
- Backticks for raw/multi-line strings

---

### **5. Control Flow: If Statements (4-5 min)**

**Topics to cover:**
- Basic if/else
- If with initialization statement
- No parentheses needed (Go style)
- Comparison operators

**Code Examples:**
```go
// Basic if/else
age := 20
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}

// If with initialization
if score := calculateScore(); score > 90 {
    fmt.Println("Excellent!")
} else if score > 70 {
    fmt.Println("Good job!")
} else {
    fmt.Println("Keep practicing!")
}
// score is only available inside if/else block

// Comparison operators
x, y := 10, 20
if x < y {
    fmt.Println("x is less than y")
}
if x != y {
    fmt.Println("x and y are different")
}
```

**Key teaching points:**
- No parentheses around condition (Go enforces clean style)
- Braces are mandatory (prevents bugs)
- Init statement scope is limited to if/else block
- Standard comparison operators: `==`, `!=`, `<`, `>`, `<=`, `>=`

---

### **6. Control Flow: Loops (5-6 min)**

**Topics to cover:**
- For loop (the only loop in Go!)
- While-style loop
- Infinite loop
- Range over collections
- Break and continue

**Code Examples:**
```go
// Classic for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-style loop
count := 0
for count < 3 {
    fmt.Println("Count:", count)
    count++
}

// Infinite loop (with break)
counter := 0
for {
    counter++
    if counter > 5 {
        break
    }
    if counter == 3 {
        continue  // Skip to next iteration
    }
    fmt.Println(counter)
}

// Range over string
name := "Go"
for index, char := range name {
    fmt.Printf("Index %d: %c\n", index, char)
}

// Ignore index with _
for _, char := range name {
    fmt.Printf("%c ", char)
}
```

**Key teaching points:**
- Go only has `for` - no `while` or `do-while`
- Different forms of `for` cover all loop needs
- `range` is idiomatic for iterating
- Use `_` to ignore values you don't need
- `break` exits loop, `continue` skips to next iteration

---

### **7. Control Flow: Switch (3-4 min)**

**Topics to cover:**
- Basic switch
- Multiple values in case
- No fallthrough by default
- Switch without expression (replaces if/else chains)

**Code Examples:**
```go
// Basic switch
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("Almost weekend!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Midweek day")
}

// Switch with initialization
switch hour := 14; {
case hour < 12:
    fmt.Println("Good morning")
case hour < 17:
    fmt.Println("Good afternoon")
default:
    fmt.Println("Good evening")
}

// Type switch (preview for later)
// We'll cover this more when we get to interfaces
```

**Key teaching points:**
- No `break` needed (doesn't fall through by default)
- Can have multiple values per case
- Switch without expression acts like if/else chain
- Cleaner than long if/else chains

---

### **8. Practical Example: Building a Simple Program (5-6 min)**

**Build together:** A number guessing game or grade calculator

**Example: Grade Calculator**
```go
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

**Walk through:**
- Declare variables for student data
- Use for loop with range to calculate total
- Type conversion for float division
- Switch statement for grade assignment
- If statement for pass/fail
- Formatted output with Printf

---

### **9. Common Beginner Mistakes (3-4 min)**

**Cover these pitfalls:**

```go
// 1. Unused variables (compilation error)
func badExample() {
    x := 10  // Declared but never used - won't compile!
}

// 2. Shadowing variables
count := 5
if true {
    count := 10  // This creates a NEW variable!
    fmt.Println(count)  // Prints 10
}
fmt.Println(count)  // Still prints 5

// 3. Wrong scope with :=
var err error
if data, err := getData(); err != nil {  // Creates NEW err!
    return err
}
// Original err is still nil here

// Better:
var data string
data, err = getData()  // Use existing err
if err != nil {
    return err
}

// 4. Implicit type conversion
var x int = 10
var y float64 = x  // ERROR! Must use float64(x)
```

---

### **10. Next Steps & Wrap-up (2-3 min)**

**Recap what was covered:**
- Variables and type system
- If statements and conditions
- For loops (all forms)
- Switch statements

**Preview next topics:**
- Functions and parameters
- Structs and methods
- Slices and maps
- Error handling

**Homework/Practice suggestions:**
- Build a temperature converter (Celsius ↔ Fahrenheit)
- Create a simple calculator
- Write FizzBuzz
- Build a multiplication table generator

**Resources:**
- Go Tour: tour.golang.org
- Go Playground: play.golang.org
- Your GitHub repo with example code

---

## **Production Notes**

### **Screen Setup:**
- Code editor on left (80% screen)
- Terminal on right for output
- Use large font (18-20pt minimum)
- Dark theme for reduced eye strain

### **Teaching Techniques:**
- Type code live (don't copy/paste) - shows thinking process
- Run code frequently to show output
- Intentionally make small mistakes and fix them
- Pause after each concept for viewers to try themselves

### **Code Quality Reminders:**
- Use meaningful variable names
- Add comments explaining "why" not "what"
- Show `go fmt` to auto-format
- Mention `go vet` for catching issues

### **Engagement:**
- Ask rhetorical questions: "What do you think this will print?"
- "Pause here and try this yourself"
- "Notice how Go forces us to..." (highlight safety features)

---

## **Supplementary Materials to Provide**

1. **GitHub Repository:**
   - All code examples from video
   - Practice exercises with solutions
   - README with setup instructions

2. **Cheat Sheet (PDF/Gist):**
   - Variable declaration syntax
   - All loop forms
   - Switch patterns
   - Common operators

3. **Practice Exercises:**
   - Easy: Print numbers 1-100
   - Medium: FizzBuzz
   - Challenge: Simple text-based menu system

---

This structure keeps beginners engaged while building a solid foundation. The practical example ties everything together, and highlighting common mistakes prevents frustration. 
