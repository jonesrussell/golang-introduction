# Instructor Notes: Best Practices

## Teaching Techniques
- Provide clear guidelines
- Show good vs bad examples
- Emphasize consistency
- Connect to real-world code

## Guidelines to Cover

### Pointer Receivers
- Use when method modifies receiver
- Use when receiver is large
- Be consistent within a type

### Function Parameters
- Use pointer when function modifies argument
- Use pointer when nil is valid (optional)

### Return Values
- Return pointer when nil is meaningful
- Return pointer for large objects
- Return value for small, immutable types

### Initialization
- Prefer `&Type{}` with values
- `new(Type)` for zero values
- Constructor pattern for defaults

## Key Emphasis
- **Consistency**: Pick one style per type
- **Document nil**: When nil is a valid value
- **Idiomatic patterns**: Follow Go conventions
- **Clarity**: Make intent clear

## Common Questions
- "Should all methods use pointer receivers?" - If any do, all should (consistency)
- "When do I return a pointer?" - When nil is meaningful or object is large
- "What's the best initialization?" - `&Type{}` with values, or constructor

## Engagement
- "Notice the consistency - all methods use pointer receivers"
- "This is idiomatic Go code"
- "Following these patterns makes code predictable"

## Real-World Context
- Standard library follows these patterns
- Code reviews will check for consistency
- These patterns are everywhere in Go codebases

## Transition
- "Let's wrap up and see what's next..."
