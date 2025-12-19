# Instructor Notes: Error Interface

## Teaching Techniques
- Show error is just an interface (connect to tutorial 5!)
- Demonstrate creating errors
- Show nil means "no error"
- Emphasize: "It's just an interface with one method"

## Demo Flow
1. Show error interface definition (just Error() string)
2. Create error with errors.New()
3. Create formatted error with fmt.Errorf()
4. Show nil check pattern
5. Show function returning error

## Key Emphasis
- **error is an interface**: Just like any other interface
- **errors.New()**: Simple error creation
- **fmt.Errorf()**: Formatted error creation
- **nil means success**: Zero value of error is nil
- **Any type with Error()**: Can be an error

## Common Questions
- "Is error special?" - No! Just an interface
- "Why nil for no error?" - Zero value pattern
- "Can I create my own error types?" - Yes! (Coming in section 5)

## Engagement
- "Notice error is just an interface - you learned this in tutorial 5!"
- "nil means no error - this is Go's pattern"
- "Every function that can fail returns an error"

## Real-World Context
- Every Go function that can fail follows this pattern
- Standard library uses this everywhere
- This is idiomatic Go

## Transition
- "Now let's see how to handle these errors..."
