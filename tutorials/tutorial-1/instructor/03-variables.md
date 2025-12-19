# Instructor Notes: Variables & Declaration

## Teaching Techniques
- Show all three ways to declare, then explain when to use each
- Emphasize `:=` as the most common (80% of the time)
- Show zero values in playground

## Demo Flow
1. Start with explicit `var name string = "value"`
2. Show type inference `var name = "value"`
3. Introduce `:=` shorthand
4. Show zero values by printing uninitialized vars
5. Demo constants
6. Show grouped declaration with `var ()`

## Key Emphasis
- `:=` only works inside functions (demo the error!)
- Unused variables = compilation error (show this!)
- Zero values are a safety feature

## Engagement
- Ask: "Which style do you prefer so far?"
- Challenge: "What's the zero value of bool?"
- "Pause here and try declaring your own variables"

## Code Quality Reminders
- Use meaningful variable names
- Add comments explaining "why" not "what"
- Show `go fmt` to auto-format

## Common Mistakes to Show
- Try using `:=` at package level (error)
- Declare but don't use a variable (error)
