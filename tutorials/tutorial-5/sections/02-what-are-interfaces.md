# What Are Interfaces?

**Duration:** 5-6 minutes

## Topics to cover:
- [Interfaces](https://go.dev/ref/spec#Interface_types) define behavior, not data
- Method signatures only
- [Implicit implementation](https://go.dev/ref/spec#Interface_types)
- Any type that has the methods satisfies the interface

## Code Examples

```go runnable
package main

import "fmt"

// Interface definition - just method signatures
type Speaker interface {
    Speak() string
}

// Dog implements Speaker (implicitly - no "implements" keyword!)
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

// Cat also implements Speaker
type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

// Robot implements Speaker too
type Robot struct {
    Model string
}

func (r Robot) Speak() string {
    return "Beep boop!"
}

// Function accepts any Speaker
func MakeSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    robot := Robot{Model: "R2D2"}

    // All satisfy Speaker interface
    MakeSpeak(dog)    // Woof!
    MakeSpeak(cat)    // Meow!
    MakeSpeak(robot)  // Beep boop!

    // Can store in slice of interface type
    speakers := []Speaker{dog, cat, robot}
    for _, s := range speakers {
        fmt.Println(s.Speak())
    }
}
```

## Key teaching points:
- Interface = set of [method signatures](https://go.dev/ref/spec#Interface_types)
- No `implements` keyword - [implementation is implicit](https://go.dev/ref/spec#Interface_types)
- If a type has the methods, it [satisfies the interface](https://go.dev/ref/spec#Interface_types)
- Enables [polymorphism](https://go.dev/doc/faq#polymorphism) without inheritance
- Types don't need to know about interfaces they implement
