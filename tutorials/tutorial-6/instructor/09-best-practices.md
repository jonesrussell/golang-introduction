# Instructor Notes: Best Practices

## Teaching Techniques
- Provide clear guidelines
- Show good vs bad examples
- Emphasize consistency
- Connect to real-world code

## Key Practices to Cover

### 1. Always Check Errors
- **Show**: Every function that returns error
- **Bad**: Ignoring with `_`
- **Good**: Check and handle

### 2. Add Context When Wrapping
- **Show**: Bare return vs wrapped
- **Emphasize**: Each layer adds context
- **Pattern**: "operation: context: root error"

### 3. Use errors.Is for Sentinel Errors
- **Show**: `==` vs errors.Is()
- **Why**: Works with wrapped errors
- **Pattern**: Check entire chain

### 4. Use errors.As for Custom Types
- **Show**: Type assertion vs errors.As()
- **Why**: Works with wrapped errors
- **Pattern**: Extract typed error

### 5. Return Early
- **Show**: Nested ifs vs early returns
- **Why**: Cleaner, less nesting
- **Pattern**: Check error, return immediately

### 6. Document Errors
- **Show**: Function comments listing possible errors
- **Why**: Part of API contract
- **Pattern**: Document in godoc

## Key Emphasis
- **Check every error**: No exceptions
- **Add context**: Makes debugging easier
- **Use Is/As**: Type-safe error checking
- **Document errors**: Part of your API

## Common Questions
- "Do I really need to check every error?" - Yes!
- "How much context should I add?" - Enough to understand where it happened
- "Should I document all errors?" - Yes, especially sentinel errors

## Engagement
- "Notice how these practices make code robust"
- "Error documentation is part of your API"
- "These patterns are everywhere in Go code"

## Real-World Context
- Standard library follows these practices
- Code reviews check for proper error handling
- These patterns prevent production bugs

## Transition
- "Let's wrap up and see what's next..."
