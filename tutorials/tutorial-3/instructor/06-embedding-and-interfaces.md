# Instructor Notes: Embedding and Interfaces

## Teaching Techniques
- Show interface satisfaction through embedding
- Demonstrate method promotion satisfies interfaces
- Show interface embedding (composing interfaces)
- Connect to standard library (io.Reader, io.Writer)

## Demo Flow
1. Show Notifier interface
2. Show EmailNotifier implementing it
3. Show User embedding EmailNotifier
4. Demonstrate: User automatically satisfies Notifier!
5. Show method overriding
6. Show interface embedding (Reader + Writer = ReadWriter)

## Key Emphasis
- **Interface satisfaction**: Embedded methods count!
- **Method overriding**: Outer struct can override embedded methods
- **Calling original**: Can still call `a.User.Notify()`
- **Interface embedding**: Compose interfaces (like io.ReadWriteCloser)

## Common Questions
- "Does embedding make my type satisfy interfaces?" - Yes, if embedded type does
- "Can I override embedded methods?" - Yes! Outer struct method wins
- "What's interface embedding?" - Composing interfaces together

## Engagement
- "Notice how User satisfies Notifier without implementing it!"
- "This is how the standard library composes interfaces"
- "Embedding gives us both code reuse AND interface satisfaction"

## Real-World Context
- Standard library uses this extensively (io package)
- Common pattern: embed to satisfy interfaces
- Enables polymorphism without inheritance

## Transition
- "Let's see practical patterns using embedding..."
