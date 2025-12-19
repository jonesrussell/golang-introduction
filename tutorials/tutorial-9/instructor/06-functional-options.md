# Instructor Notes: Functional Options

## Teaching Techniques
- Show functional options pattern
- Demonstrate how it works
- Show when to use it
- Emphasize: "Elegant way to handle optional parameters"

## Key Emphasis
- **Functional options**: Functions that configure struct
- **Flexible**: Easy to add new options
- **Backward compatible**: Adding options doesn't break code
- **Clean API**: No need for many constructors

## Common Questions
- "When do I use functional options?" - When you have many optional parameters
- "How do I implement it?" - Options are functions that modify struct
- "Is it worth it?" - Yes, for complex configuration

## Engagement
- "Notice how clean the API is"
- "Adding new options is easy"
- "This is a common Go pattern"

## Real-World Context
- Used extensively in Go libraries
- Standard pattern for configuration
- Makes APIs more flexible

## Transition
- "Let's put it all together in a practical example..."
