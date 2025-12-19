# Instructor Notes: Practical Example

## Teaching Techniques
- Build incrementally, don't paste all at once
- Show error handling at each step
- Demonstrate wrapping with context
- Show different error types in action

## Build Order
1. Define sentinel errors
2. Define custom error types (ValidationError, FileError)
3. Implement Unwrap() for FileError
4. Create file reading function (wraps errors)
5. Create validation function (returns ValidationError)
6. Create processing function (handles multiple error types)
7. Show errors.Is() and errors.As() in action
8. Show complete error handling flow

## Live Commentary
- "First, let's define our error types..."
- "Notice how FileError wraps the underlying error..."
- "ValidationError gives us structured error data..."
- "Now we can handle different error types programmatically..."

## Things to Emphasize
- Custom error types for structured data
- Sentinel errors for simple cases
- Wrapping adds context at each layer
- errors.Is() and errors.As() for type-safe checks

## Engagement
- "What happens if the file doesn't exist?"
- "How do we handle validation errors differently?"
- Challenge: "Add a retry mechanism for transient errors"

## Variations to Mention
- Could add more error types
- Could add retry logic
- Could add error aggregation
- Could add structured logging

## Common Mistakes to Watch For
- Forgetting to wrap errors
- Using `==` instead of errors.Is()
- Not implementing Unwrap()
- Ignoring errors
