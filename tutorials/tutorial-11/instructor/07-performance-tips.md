# Instructor Notes: Performance Tips

## Teaching Techniques
- Show performance considerations
- Demonstrate zero-allocation patterns
- Show when to use sugar logger
- Emphasize: "Zap is fast, but use it wisely"

## Key Emphasis
- **Zero-allocation**: Use structured fields
- **Sugar logger**: Convenient but slower
- **Level checks**: Check level before expensive operations
- **Sampling**: For high-volume logs

## Common Questions
- "Is Zap fast?" - Yes, very fast
- "When do I use sugar logger?" - Development, convenience
- "How do I handle high volume?" - Use sampling

## Engagement
- "Notice how structured logging is fast"
- "Zap is optimized for performance"
- "Use wisely for best performance"

## Real-World Context
- Performance matters in production
- Zap is optimized for speed
- Sampling for high-volume scenarios

## Transition
- "Let's wrap up..."
