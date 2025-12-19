# Control Flow: Switch

**Duration:** 3-4 minutes

## Topics to cover:
- Basic [switch](https://go.dev/ref/spec#Switch_statements)
- [Multiple values in case](https://go.dev/tour/flowcontrol/10)
- [No fallthrough by default](https://go.dev/ref/spec#Switch_statements)
- [Switch without expression](https://go.dev/tour/flowcontrol/11) (replaces if/else chains)

## Code Examples

```go snippet
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

## Key teaching points:
- [No `break` needed](https://go.dev/ref/spec#Switch_statements) (doesn't fall through by default)
- Can have [multiple values per case](https://go.dev/tour/flowcontrol/10)
- [Switch without expression](https://go.dev/tour/flowcontrol/11) acts like if/else chain
- Cleaner than long if/else chains
