# Instructor Notes: Panic and Recover

## Teaching Techniques
- Emphasize: "Panic is rare in Go!"
- Show when panic is appropriate
- Show recover pattern
- Contrast panic vs error

## Demo Flow
1. Show what panic does (stops execution)
2. Show common causes (nil pointer, index out of range)
3. Show when to panic (programmer errors, impossible states)
4. Show when NOT to panic (expected failures)
5. Show recover pattern
6. Show goroutine recovery

## Key Emphasis
- **Panic is rare**: Most code uses errors
- **When to panic**: Programmer errors, impossible states
- **When NOT to panic**: Expected failures (use errors!)
- **Recover**: Only works in deferred functions
- **Goroutines**: Each must recover its own panics

## Common Questions
- "When should I panic?" - Almost never! Use errors instead
- "What's the difference from exceptions?" - Panic is for programmer errors
- "Can I recover from any panic?" - Yes, but only in deferred functions

## Engagement
- "Panic is for 'this should never happen' situations"
- "If it can happen at runtime, use an error"
- "Recover is like a safety net"

## Real-World Context
- Most Go code never uses panic
- Recover used in HTTP servers, goroutine pools
- Convention: `Must*` functions panic on error

## Gotchas
- Recover only works in deferred functions
- Each goroutine must recover its own panics
- Panic across goroutines can't be caught by main

## Transition
- "Let's put it all together in a practical example..."
