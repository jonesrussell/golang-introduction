# Instructor Notes: Best Practices

## Teaching Techniques
- Provide clear guidelines
- Show good vs bad examples
- Emphasize: "These practices prevent bugs"
- Connect to production code

## Key Practices to Cover

### 1. Keep Goroutines Focused
- **Show**: One task per goroutine
- **Why**: Easier to reason about
- **Pattern**: Short-lived, focused

### 2. Use Context for Cancellation
- **Show**: ctx.Done() checks
- **Why**: Enables graceful shutdown
- **Pattern**: Always pass context

### 3. Prefer Channels for Communication
- **Show**: Channels vs mutex
- **Why**: "Don't communicate by sharing memory"
- **Pattern**: Use channels when possible

### 4. Close Channels from Sender
- **Show**: Sender closes, receiver ranges
- **Why**: Clear ownership
- **Pattern**: defer close(ch)

### 5. Use Buffered Channels to Prevent Blocking
- **Show**: Buffered vs unbuffered
- **Why**: Prevents deadlocks
- **Pattern**: Size based on throughput needs

### 6. Always Handle Cancellation
- **Show**: select with ctx.Done()
- **Why**: Prevents leaks
- **Pattern**: Check in loops

### 7. Use WaitGroup for Multiple Goroutines
- **Show**: WaitGroup pattern
- **Why**: Coordinate completion
- **Pattern**: Add before, Done in defer

### 8. Limit Concurrency
- **Show**: Semaphore pattern
- **Why**: Prevent resource exhaustion
- **Pattern**: Buffered channel as semaphore

### 9. Run Race Detector
- **Show**: `go test -race`
- **Why**: Catches data races
- **Pattern**: Always in CI/CD

### 10. Keep Mutex Scope Small
- **Show**: Lock, do work, unlock
- **Why**: Reduces contention
- **Pattern**: Use defer

## Key Emphasis
- **Design for cancellation**: From the start
- **Channels over mutex**: When possible
- **Always clean up**: Prevent leaks
- **Test with race detector**: Essential tool

## Common Questions
- "How do I know if I need a mutex?" - If shared state is modified
- "When do I use buffered channels?" - When you need throughput
- "Do I always need context?" - For cancellable operations, yes

## Engagement
- "Notice how these practices make code safer"
- "Race detector is your friend"
- "These patterns are everywhere in production"

## Real-World Context
- Standard library follows these practices
- Code reviews check for these patterns
- These practices prevent production bugs

## Transition
- "Let's wrap up and see what's next..."
