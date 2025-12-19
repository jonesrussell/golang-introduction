# Instructor Notes: Best Practices

## Teaching Techniques
- Provide clear guidelines
- Show good vs bad examples
- Emphasize: "These practices make code maintainable"

## Key Practices
- Accept interfaces, return structs
- Small, focused interfaces
- Constructor injection for dependencies
- Define interfaces where you use them
- Use functional options for complex configuration

## Key Emphasis
- **Accept interfaces**: Function parameters should be interfaces
- **Return structs**: Return concrete types
- **Small interfaces**: One method is often enough
- **Where to define**: Where you use them

## Common Questions
- "Where do I define interfaces?" - Where you use them
- "How big should interfaces be?" - Small, focused
- "When do I use functional options?" - Complex configuration

## Engagement
- "Notice how these practices make code cleaner"
- "This is the Go way"
- "These patterns are everywhere in production"

## Real-World Context
- Standard library follows these practices
- Used extensively in production
- Makes code more maintainable

## Transition
- "Let's wrap up..."
