# Instructor Notes: Init Abuse

## Teaching Techniques
- Show bad example (using init for everything)
- Explain why it's wrong
- Show the fix (explicit initialization)
- Emphasize: "Init should be minimal"

## Key Emphasis
- **Init abuse**: Using init for too much
- **Hard to test**: Init runs automatically
- **Hidden dependencies**: Dependencies are not explicit
- **Fix**: Use explicit initialization

## Common Questions
- "When should I use init?" - Rarely, for package-level setup
- "What's wrong with init?" - Hard to test, hidden dependencies
- "How do I fix it?" - Use explicit initialization functions

## Engagement
- "Notice how init hides dependencies"
- "This makes testing hard"
- "Explicit initialization is better"

## Real-World Context
- Init abuse is common
- Makes testing difficult
- Explicit initialization is preferred

## Transition
- "Let's wrap up with a summary..."
