# Instructor Notes: Best Practices

## Teaching Techniques
- Provide clear guidelines
- Show good vs bad examples
- Emphasize: "These practices make logging effective"

## Key Practices
- Use appropriate log levels
- Include context (request ID, user ID)
- Don't log sensitive data
- Use structured fields
- Log errors with context

## Key Emphasis
- **Log levels**: Use appropriately
- **Context**: Include request ID, user ID, etc.
- **Security**: Don't log passwords, tokens
- **Structured**: Use fields, not string formatting

## Common Questions
- "What should I log?" - Errors, important events, context
- "What shouldn't I log?" - Sensitive data, too verbose
- "How much context?" - Enough to debug, not too much

## Engagement
- "Notice how context helps debugging"
- "Structured fields make searching easy"
- "These practices make logging effective"

## Real-World Context
- These practices are standard
- Used in production systems
- Essential for debugging

## Transition
- "Let's put it all together in a practical example..."
