# Instructor Notes: Control Flow - Switch

## Teaching Techniques
- Compare to C/Java switch (no fallthrough!)
- Show switch as if/else replacement
- Demo multiple values in case

## Demo Flow
1. Basic switch with string
2. Multiple values in one case
3. Switch with init statement
4. Switch without expression (like if/else)

## Key Emphasis
- NO break needed - doesn't fall through!
- Multiple values: `case "a", "b", "c":`
- Switch without expression is powerful

## Live Coding
```go
// Show this is cleaner than if/else chains
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
// ...
}
```

## Engagement
- Ask: "Is this cleaner than if/else?"
- Mention type switch exists (preview for interfaces)
- "When would you use switch vs if?"

## When to Use Switch
- Multiple discrete values to check
- Replacing long if/else chains
- Type assertions (later topic)
