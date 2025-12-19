# Instructor Notes: Sentinel Errors

## Teaching Techniques
- Show package-level error variables
- Demonstrate errors.Is() for comparison
- Show wrapping sentinel errors
- Contrast with custom error types

## Demo Flow
1. Show package-level sentinel errors (ErrXxx naming)
2. Show returning sentinel errors
3. Show checking with errors.Is()
4. Show wrapping sentinel errors
5. Show standard library examples (os.ErrNotExist, io.EOF)

## Key Emphasis
- **Sentinel errors**: Package-level variables
- **Naming convention**: ErrXxx
- **Use errors.Is()**: Not `==` (works with wrapped errors)
- **Document as API**: Part of your package's contract
- **When to use**: No dynamic data needed

## Common Questions
- "Why not use string comparison?" - Fragile, breaks with wrapping
- "When do I use sentinel vs custom type?" - Sentinel for simple, custom for structured
- "Can I wrap sentinel errors?" - Yes! Use errors.Is() to check

## Engagement
- "Notice the ErrXxx naming convention"
- "errors.Is() works even if error is wrapped"
- "Standard library uses this pattern extensively"

## Real-World Context
- Standard library has many sentinel errors
- Part of API design
- Enables robust error checking

## Transition
- "Now let's talk about panic and recover..."
