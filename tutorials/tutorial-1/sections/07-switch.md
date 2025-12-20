# Control Flow: Switch

**Duration:** 3-4 minutes

## Topics to cover:
- Basic [switch](https://go.dev/ref/spec#Switch_statements)
- [Multiple values in case](https://go.dev/tour/flowcontrol/10)
- [No fallthrough by default](https://go.dev/ref/spec#Switch_statements)
- [Switch without expression](https://go.dev/tour/flowcontrol/11) (replaces if/else chains)

## Basic [switch](https://go.dev/ref/spec#Switch_statements)

The switch statement provides a clean way to handle multiple conditions. Unlike some languages, Go doesn't require `break` statements - cases don't fall through by default.

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
```

## [Multiple values in case](https://go.dev/tour/flowcontrol/10)

You can list multiple values in a single case statement, separated by commas.

```go snippet
// Multiple values in case (already shown above)
day := "Saturday"
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Weekday")
}
```

## [Switch without expression](https://go.dev/tour/flowcontrol/11)

A switch without an expression acts like an if/else chain, making it cleaner than long if/else statements.

```go snippet
// Switch without expression (acts like if/else chain)
hour := 14
switch {
case hour < 12:
    fmt.Println("Good morning")
case hour < 17:
    fmt.Println("Good afternoon")
default:
    fmt.Println("Good evening")
}
```

## Key teaching points:
- [No `break` needed](https://go.dev/ref/spec#Switch_statements) (doesn't fall through by default)
- Can have [multiple values per case](https://go.dev/tour/flowcontrol/10)
- [Switch without expression](https://go.dev/tour/flowcontrol/11) acts like if/else chain
- Cleaner than long if/else chains
