# Instructor Notes: Interface Satisfaction Rules

## Teaching Techniques
- Show the receiver type matters!
- Demonstrate value vs pointer receiver implications
- Show compile-time checks
- Emphasize: "This trips up many developers"

## Demo Flow
1. Show Buffer with value receiver - both Type and *Type work
2. Show File with pointer receiver - only *Type works
3. Show the compile error for File{}
4. Explain the rule clearly
5. Show compile-time check pattern

## Key Emphasis
- **Value receiver**: Both `Type` and `*Type` satisfy interface
- **Pointer receiver**: Only `*Type` satisfies interface
- **Exact match required**: Method signatures must match exactly
- **Compile-time check**: Use `var _ Interface = Type{}` to verify

## Common Questions
- "Why does pointer receiver matter?" - Method set rules
- "How do I know if my type satisfies?" - Compile-time check
- "Can I mix value and pointer receivers?" - Yes, but be careful

## Engagement
- "Watch this - the receiver type matters!"
- "This is a common gotcha - let's understand it"
- "The compile-time check catches this early"

## Gotchas
- Pointer receiver means only pointers satisfy
- Value receiver means both values and pointers satisfy
- Method signature must match exactly (parameters and returns)

## Transition
- "Now let's see a special interface - the empty interface..."
