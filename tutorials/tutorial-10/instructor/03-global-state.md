# Instructor Notes: Global State

## Teaching Techniques
- Show bad example (global variables)
- Explain why it's wrong
- Show the fix (dependency injection)
- Emphasize: "Global state makes testing hard"

## Key Emphasis
- **Global state**: Hard to test
- **Hidden dependencies**: Dependencies are not explicit
- **Race conditions**: Global state is not thread-safe
- **Fix**: Use dependency injection

## Common Questions
- "What's wrong with globals?" - Hard to test, hidden dependencies
- "When are globals OK?" - Rarely, maybe for constants
- "How do I fix it?" - Use dependency injection

## Engagement
- "Notice how hard it is to test this"
- "Global state creates hidden dependencies"
- "Dependency injection solves this"

## Real-World Context
- Global state is common in legacy code
- Makes testing difficult
- Dependency injection is the solution

## Transition
- "Let's see interface pollution..."
