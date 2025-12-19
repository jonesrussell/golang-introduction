# Instructor Notes: Testing with Interfaces

## Teaching Techniques
- Show how interfaces enable testing
- Demonstrate mock implementation
- Show table-driven tests
- Emphasize: "This is why interfaces matter"

## Demo Flow
1. Show production code with interface dependency
2. Show mock implementation (simple struct)
3. Show test using mock
4. Show error scenario testing
5. Show table-driven test pattern

## Key Emphasis
- **Mock implementations**: Easy to create (just implement interface)
- **Test behavior**: Not implementation details
- **Error scenarios**: Mock can simulate failures
- **No frameworks needed**: Go's interfaces are enough
- **Dependency injection**: Interfaces enable it

## Common Questions
- "Do I need a mocking framework?" - No! Just implement the interface
- "How do I test error cases?" - Mock can return errors
- "Is this dependency injection?" - Yes! Via interfaces

## Engagement
- "Notice how easy it is to create a mock"
- "We can test error scenarios without a real database"
- "This is why 'accept interfaces' matters"

## Real-World Context
- This pattern is used everywhere in Go testing
- Makes tests fast and isolated
- Enables testing without external dependencies

## Best Practices
- Keep mocks simple
- Test behavior, not implementation
- Use table-driven tests for multiple scenarios
- Mock only what you need

## Transition
- "Let's watch out for common mistakes..."
