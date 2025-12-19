# Instructor Notes: Constructor Injection

## Teaching Techniques
- Show constructor injection pattern
- Demonstrate how it enables testing
- Show interface usage
- Emphasize: "This is the Go way"

## Key Emphasis
- **Constructor injection**: Dependencies passed in constructor
- **Interfaces**: Define dependencies as interfaces
- **Testable**: Easy to inject mocks
- **Flexible**: Can swap implementations

## Common Questions
- "How do I inject dependencies?" - Pass them in constructor
- "What interfaces do I need?" - Define what you need, not what you have
- "Is this the only way?" - No, but it's the most common

## Engagement
- "Notice how easy it is to test now"
- "We can swap implementations easily"
- "This is the Go way of doing DI"

## Real-World Context
- Constructor injection is standard in Go
- Used extensively in production
- Enables testable code

## Transition
- "Now let's see how to test with mocks..."
