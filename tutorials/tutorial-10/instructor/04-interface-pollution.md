# Instructor Notes: Interface Pollution

## Teaching Techniques
- Show bad example (too many interfaces)
- Explain why it's wrong
- Show the fix (small interfaces)
- Emphasize: "Small interfaces are better"

## Key Emphasis
- **Interface pollution**: Too many, too large interfaces
- **Small interfaces**: One method is often enough
- **Where to define**: Where you use them, not where you implement
- **Go philosophy**: Small, focused interfaces

## Common Questions
- "How many methods should an interface have?" - As few as possible, often one
- "Where do I define interfaces?" - Where you use them
- "What's wrong with large interfaces?" - Less flexible, harder to implement

## Engagement
- "Notice how small these interfaces are"
- "This is the Go way - small, focused"
- "Large interfaces are less flexible"

## Real-World Context
- Interface pollution is common
- Go standard library uses small interfaces
- Small interfaces are more flexible

## Transition
- "Now let's see nil pointer paranoia..."
