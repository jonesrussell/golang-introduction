# Instructor Notes: Common Pitfalls

## Teaching Techniques
- Show each pitfall with bad example
- Explain why it's wrong
- Show the fix
- Emphasize: "These are common mistakes - learn them!"

## Pitfalls to Cover

### 1. Race Conditions
- **Problem**: Concurrent access without protection
- **Fix**: Mutex or atomic operations
- **Show**: Race detector catches this

### 2. Deadlocks
- **Problem**: Circular wait (goroutines waiting for each other)
- **Fix**: Design channel flow carefully
- **Show**: Deadlock detection

### 3. Goroutine Leaks
- **Problem**: Goroutine blocked forever
- **Fix**: Use buffered channels or ensure receiver
- **Show**: Context for cancellation

### 4. Wrong Channel Close
- **Problem**: Closing from receiver, sending to closed
- **Fix**: Only sender closes, never send after close
- **Show**: Proper closing pattern

### 5. WaitGroup Copy
- **Problem**: Copying WaitGroup (value vs pointer)
- **Fix**: Pass pointer or use closure
- **Show**: The bug and fix

## Key Emphasis
- **Race detector**: Always use `go run -race`
- **Deadlock detection**: Go runtime detects some deadlocks
- **Goroutine leaks**: Use context for cancellation
- **Channel rules**: Sender closes, never send after close

## Common Questions
- "How do I find race conditions?" - Use race detector
- "What causes deadlocks?" - Circular waits
- "How do I prevent leaks?" - Always have a way to cancel/exit

## Engagement
- "Watch this - it looks fine but has a race condition"
- "The race detector will catch this"
- "This is why we use context for cancellation"

## Real-World Context
- These pitfalls are common in production
- Race detector is essential tool
- Proper design prevents most issues

## Transition
- "Let's summarize best practices..."
