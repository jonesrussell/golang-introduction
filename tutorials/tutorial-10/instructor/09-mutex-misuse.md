# Instructor Notes: Mutex Misuse

## Teaching Techniques
- Show bad example (wrong mutex usage)
- Explain why it's wrong
- Show the fix (proper locking)
- Emphasize: "Use channels when possible, mutex when needed"

## Key Emphasis
- **Mutex misuse**: Wrong scope, deadlocks, not needed
- **Channels first**: Prefer channels for communication
- **Mutex when needed**: For shared state protection
- **Proper locking**: Lock scope, defer unlock

## Common Questions
- "When do I use mutex vs channels?" - Channels for communication, mutex for shared state
- "What's wrong with this mutex?" - Wrong scope, deadlock risk
- "How do I avoid deadlocks?" - Lock in consistent order, use defer

## Engagement
- "Notice how this mutex is misused"
- "Channels are often better"
- "Use mutex only when needed"

## Real-World Context
- Mutex misuse is common
- Channels are often better
- Proper locking is important

## Transition
- "Let's see init function abuse..."
