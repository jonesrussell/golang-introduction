# Control Flow: Switch

**Duration:** 3-4 minutes

## Topics to cover:
- Basic switch
- Multiple values in case
- No fallthrough by default
- Switch without expression (replaces if/else chains)

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
- No `break` needed (doesn't fall through by default)
- Can have multiple values per case
- Switch without expression acts like if/else chain
- Cleaner than long if/else chains
