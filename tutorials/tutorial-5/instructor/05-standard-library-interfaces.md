# Instructor Notes: Standard Library Interfaces

## Teaching Techniques
- Show how small interfaces are powerful
- Demonstrate io.Reader and io.Writer
- Show fmt.Stringer in action
- Show error is just an interface
- Demonstrate interface composition

## Demo Flow
1. Show io.Reader (one method - very small!)
2. Show custom type implementing it
3. Show it works with io.ReadAll
4. Show fmt.Stringer - custom string representation
5. Show error interface (just one method!)
6. Show interface composition (ReadWriter)

## Key Emphasis
- **Small interfaces**: One method can be powerful (Reader, Writer)
- **Composition**: Build larger interfaces from smaller ones
- **Stringer**: Customize how types print
- **error**: Just an interface with Error() method
- **Ecosystem**: Standard library interfaces enable interop

## Common Questions
- "Why is Reader so small?" - Small = many implementations possible
- "How do I customize fmt.Println output?" - Implement Stringer
- "Is error special?" - No! Just an interface

## Engagement
- "Notice how small Reader is - just one method!"
- "Our custom type works with standard library functions"
- "This is how Go achieves ecosystem compatibility"

## Real-World Context
- Standard library built on small interfaces
- Your types can work with standard library functions
- This pattern is everywhere in Go

## Transition
- "Let's learn how to design good interfaces..."
