# Instructor Notes: Tight Coupling

## Teaching Techniques
- Show bad example (tight coupling)
- Demonstrate why it's hard to test
- Show the problems it causes
- Emphasize: "This is what we want to avoid"

## Key Emphasis
- **Tight coupling**: Direct dependencies
- **Hard to test**: Can't mock dependencies
- **Hard to change**: Changes ripple through code
- **Problem**: Dependencies are hardcoded

## Common Questions
- "What's wrong with this code?" - Can't test it easily
- "Why is it hard to test?" - Dependencies are hardcoded
- "How do I fix it?" - Use dependency injection

## Engagement
- "Notice how hard it is to test this"
- "What if we want to use a different database?"
- "This is the problem we're solving"

## Real-World Context
- Tight coupling is common in legacy code
- Makes testing difficult
- Reduces flexibility

## Transition
- "Now let's see how constructor injection solves this..."
