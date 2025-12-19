# Instructor Notes: Interface Design

## Teaching Techniques
- Show good interface design
- Demonstrate small, focused interfaces
- Show interface segregation
- Emphasize: "Accept interfaces, return structs"

## Key Emphasis
- **Small interfaces**: One method is often enough
- **Focused**: Interface should have single responsibility
- **Accept interfaces**: Function parameters should be interfaces
- **Return structs**: Return concrete types

## Common Questions
- "How big should interfaces be?" - Small, focused
- "Where do I define interfaces?" - Where you use them
- "What's the Go philosophy?" - Small interfaces, implicit satisfaction

## Engagement
- "Notice how small these interfaces are"
- "This is the Go way - small, focused interfaces"
- "Interfaces are implicit in Go"

## Real-World Context
- Go standard library uses small interfaces
- io.Reader, io.Writer are examples
- Small interfaces are more flexible

## Transition
- "Let's see functional options pattern..."
