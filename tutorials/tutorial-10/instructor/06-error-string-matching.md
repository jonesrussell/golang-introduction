# Instructor Notes: Error String Matching

## Teaching Techniques
- Show bad example (matching error strings)
- Explain why it's wrong
- Show the fix (error types, errors.Is, errors.As)
- Emphasize: "Use error types, not strings"

## Key Emphasis
- **String matching**: Fragile, breaks easily
- **Error types**: Use sentinel errors or custom types
- **errors.Is**: Check for specific errors
- **errors.As**: Extract error details

## Common Questions
- "What's wrong with string matching?" - Fragile, breaks on changes
- "How do I check errors?" - Use errors.Is or errors.As
- "When do I use sentinel errors?" - For expected errors

## Engagement
- "Notice how string matching breaks easily"
- "Error types are more robust"
- "This is the Go way"

## Real-World Context
- Error string matching is common
- Makes code fragile
- Error types are preferred

## Transition
- "Now let's see goroutine leaks..."
