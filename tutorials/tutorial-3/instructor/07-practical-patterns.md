# Instructor Notes: Practical Patterns

## Teaching Techniques
- Show real-world patterns (mixins, decorators)
- Explain when each pattern is useful
- Show code reuse in action
- Connect to production code

## Key Patterns to Cover

### Mixin Pattern
- **Auditable**: Adds audit fields to any struct
- **SoftDeletable**: Adds soft delete functionality
- **Show**: How easy it is to add to any type

### Decorator Pattern
- **Show**: Adding behavior without modifying original
- **Example**: Logging wrapper around Writer

### Base Functionality
- **Show**: Common fields/methods in base struct
- **Example**: BaseEntity with ID and timestamps

## Key Emphasis
- **Mixins**: Reusable functionality you can add
- **Decorators**: Wrap existing functionality
- **Base types**: Common functionality shared
- **When to use**: Each pattern has its place

## Common Questions
- "When do I use a mixin?" - When you want to add functionality to multiple types
- "What's the difference from inheritance?" - No hierarchy, just composition
- "Can I combine multiple mixins?" - Yes! That's the power

## Engagement
- "Notice how Auditable can be added to any struct"
- "This is how you achieve code reuse in Go"
- "Mixins are like traits in other languages"

## Real-World Context
- Many Go libraries use mixins
- Common in web frameworks (Gin, Echo)
- Database models often use audit mixins

## Transition
- "Now let's see when to use embedding vs composition..."
