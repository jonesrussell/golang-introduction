# Instructor Notes: Wrap Up

## Recap What Was Covered
- Context value abuse
- Global state
- Interface pollution
- Nil pointer paranoia
- Error string matching
- Goroutine leaks
- Premature optimization
- Mutex misuse
- Init abuse

## Key Takeaways to Emphasize
- Learn from mistakes
- Go has specific pitfalls
- Best practices prevent bugs
- Code reviews catch these

## Preview Next Tutorial
- "Next: Structured Logging - production-ready logging"
- "You'll learn how to log effectively"
- "We'll build on everything you learned today"

## Practice Recommendations
- **Easy**: Find anti-patterns in code
- **Medium**: Refactor code to fix anti-patterns
- **Challenge**: Review code for all anti-patterns

## Cheat Sheet Highlights
- Context: For cancellation, not data
- Globals: Use dependency injection
- Interfaces: Small, focused
- Errors: Use types, not strings
- Goroutines: Always have cancellation

## Engagement
- "What anti-patterns have you seen?"
- "Try refactoring some code"
- "Questions? Let's address them now"

## Closing
- "You've learned common Go anti-patterns"
- "Avoiding these makes code better"
- "Next tutorial: Structured Logging - see you there!"
