# Instructor Notes: Nil Pointer Paranoia

## Teaching Techniques
- Show bad example (excessive nil checks)
- Explain when nil checks are needed
- Show the fix (trust the type system)
- Emphasize: "Not everything needs nil checks"

## Key Emphasis
- **Nil checks**: Only when nil is valid value
- **Trust types**: Go's type system helps
- **Zero values**: Use zero values effectively
- **When to check**: Only when nil is meaningful

## Common Questions
- "When do I check for nil?" - When nil is a valid value
- "What about pointers?" - Check if nil is meaningful
- "Is it safe to skip checks?" - Yes, if nil isn't valid

## Engagement
- "Notice how excessive nil checks clutter code"
- "Go's type system helps here"
- "Only check when nil is meaningful"

## Real-World Context
- Excessive nil checks are common
- Makes code harder to read
- Trust the type system

## Transition
- "Let's see error string matching..."
