# Instructor Notes: Goroutine Leaks

## Teaching Techniques
- Show bad example (leaking goroutines)
- Explain why it's wrong
- Show the fix (context, channels)
- Emphasize: "Always have a way to stop goroutines"

## Key Emphasis
- **Goroutine leaks**: Goroutines that never exit
- **Context**: Use context for cancellation
- **Channels**: Ensure receivers exist
- **Always cleanup**: Have a way to stop goroutines

## Common Questions
- "What causes leaks?" - Blocked goroutines with no way to exit
- "How do I prevent leaks?" - Use context for cancellation
- "How do I detect leaks?" - Monitor goroutine count

## Engagement
- "Notice how this goroutine never exits"
- "Context provides cancellation"
- "Always have a way to stop goroutines"

## Real-World Context
- Goroutine leaks are common
- Can cause resource exhaustion
- Context is the solution

## Transition
- "Let's see premature optimization..."
