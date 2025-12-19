# Instructor Notes: Custom Error Types

## Teaching Techniques
- Show when custom types are useful
- Demonstrate implementing Error() method
- Show adding fields for context
- Show implementing Unwrap() for chain support

## Demo Flow
1. Show simple custom error (NotFoundError)
2. Show using errors.As() to extract it
3. Show custom error with wrapped cause
4. Show implementing Unwrap()
5. Show HTTPError example (real-world pattern)

## Key Emphasis
- **Custom types**: Add structured error data
- **Implement Error()**: Required to satisfy interface
- **Implement Unwrap()**: Enables errors.Is/As to work
- **Pointer receiver**: Use for error methods
- **Programmatic handling**: Access fields for logic

## Common Questions
- "When do I use custom errors?" - When you need structured data
- "Do I need Unwrap()?" - Only if you wrap other errors
- "Why pointer receiver?" - Consistency, can modify if needed

## Engagement
- "Notice how we can access the fields programmatically"
- "This enables smart error handling"
- "HTTPError shows a real-world pattern"

## Real-World Context
- Custom errors used extensively in production
- Enable programmatic error handling
- Better than string matching

## Transition
- "Let's see another pattern - sentinel errors..."
