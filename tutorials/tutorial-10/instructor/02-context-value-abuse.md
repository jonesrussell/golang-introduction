# Instructor Notes: Context Value Abuse

## Teaching Techniques
- Show bad example (using context for data)
- Explain why it's wrong
- Show the fix (explicit parameters)
- Emphasize: "Context is for cancellation, not data"

## Key Emphasis
- **Context purpose**: Cancellation, timeouts, request-scoped values
- **Not for data**: Don't use context as parameter bag
- **Explicit is better**: Pass data as parameters
- **When to use**: Only for request-scoped values (user ID, trace ID)

## Common Questions
- "When can I use context values?" - Request-scoped values only
- "Why not use context for everything?" - Makes dependencies hidden
- "What's the harm?" - Makes code harder to understand and test

## Engagement
- "Notice how context values hide dependencies"
- "This makes code harder to test"
- "Explicit parameters are better"

## Real-World Context
- Context value abuse is common
- Makes code harder to understand
- Explicit parameters are preferred

## Transition
- "Now let's see global state problems..."
