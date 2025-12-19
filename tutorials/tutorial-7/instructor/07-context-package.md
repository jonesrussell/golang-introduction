# Instructor Notes: Context Package

## Teaching Techniques
- Show context as cancellation mechanism
- Demonstrate propagation
- Show timeout pattern
- Show context values (but warn: use sparingly)
- Emphasize: "Context is for cancellation, not data"

## Demo Flow
1. Show WithCancel (cancellation)
2. Show WithTimeout (timeout)
3. Show ctx.Done() check
4. Show context propagation
5. Show context values (but emphasize: rarely needed)
6. Show HTTP handler example

## Key Emphasis
- **Context for cancellation**: Primary use case
- **Propagation**: Pass context through call chain
- **First parameter**: Always pass context as first param
- **ctx.Done()**: Check in long operations
- **WithValue**: Use sparingly (not for general data)

## Common Questions
- "When do I use context?" - Cancellation, timeouts, request-scoped values
- "Do I need to pass context everywhere?" - Through call chains that need cancellation
- "What about context values?" - Use sparingly, prefer explicit parameters

## Engagement
- "Notice how cancellation propagates through the chain"
- "Context enables graceful shutdown"
- "This is how HTTP servers handle request cancellation"

## Real-World Context
- Context used extensively in HTTP servers
- Enables proper cleanup and cancellation
- Standard library uses context extensively

## Best Practices
- Always pass context as first parameter
- Check ctx.Done() in loops
- Use WithTimeout for operations that should time out
- Avoid context.WithValue for most data

## Transition
- "Let's put it all together in a practical example..."
