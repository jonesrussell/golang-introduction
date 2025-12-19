# Instructor Notes: Error Wrapping

## Teaching Techniques
- Show why context matters (debugging!)
- Demonstrate `%w` verb (preserves chain)
- Show `errors.Is()` for checking wrapped errors
- Show `errors.As()` for extracting typed errors
- Contrast `%w` vs `%v`

## Demo Flow
1. Show error without context (hard to debug)
2. Show wrapping with `%w` (adds context, preserves chain)
3. Show errors.Is() checking wrapped error
4. Show errors.As() extracting typed error
5. Show full error chain

## Key Emphasis
- **`%w` verb**: Wraps error, preserves chain
- **`%v` verb**: Formats error, breaks chain (don't use for wrapping!)
- **errors.Is()**: Checks entire chain for match
- **errors.As()**: Extracts typed error from chain
- **Add context**: Each layer adds its context

## Common Questions
- "What's the difference between `%w` and `%v`?" - `%w` preserves chain, `%v` doesn't
- "Why wrap errors?" - Adds context for debugging
- "How do I check wrapped errors?" - Use errors.Is() or errors.As()

## Engagement
- "Notice how the error chain tells the full story"
- "This makes debugging so much easier"
- "errors.Is() checks the entire chain - powerful!"

## Real-World Context
- Error wrapping is standard practice in Go
- Makes debugging production issues much easier
- Error chains show the full call stack

## Transition
- "Now let's create custom error types..."
