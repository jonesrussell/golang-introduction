# Instructor Notes: Goroutines

## Teaching Techniques
- Show the `go` keyword in action
- Demonstrate concurrent execution (interleaved output)
- Show the loop variable gotcha (important!)
- Emphasize: "Main exiting kills all goroutines"

## Demo Flow
1. Show regular function call (blocks)
2. Show goroutine call (non-blocking)
3. Show multiple goroutines running
4. Show interleaved output
5. Show loop variable capture gotcha
6. Show the fix (shadow variable or pass as argument)

## Key Emphasis
- **`go` keyword**: Starts a goroutine
- **Lightweight**: ~2KB stack (can have millions)
- **Main goroutine**: When it exits, all goroutines die
- **Loop variables**: Must capture correctly (common gotcha!)

## Common Questions
- "How many goroutines can I have?" - Millions (limited by memory)
- "What happens when main exits?" - All goroutines are killed
- "Why do I see the loop variable problem?" - Closure captures loop variable

## Engagement
- "Notice how output is interleaved - that's concurrency!"
- "Watch this gotcha with loop variables..."
- "This is why Go makes concurrency easy"

## Gotchas
- Loop variable capture (must shadow or pass as argument)
- Main exiting kills goroutines (need synchronization)
- Don't use time.Sleep for synchronization (use channels!)

## Real-World Context
- Goroutines used everywhere in Go servers
- Much lighter than OS threads
- Enable high concurrency

## Transition
- "Now let's see how goroutines communicate..."
