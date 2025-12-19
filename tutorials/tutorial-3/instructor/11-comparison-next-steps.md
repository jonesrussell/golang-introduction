# Instructor Notes: Comparison & Next Steps

## Teaching Techniques
- Compare to inheritance-based languages
- Show Go's advantages
- Recap what was covered
- Preview next topics

## Comparison to Other Languages

### Go vs Python/Java
- **Go**: No inheritance keyword, no `extends`, no `super`
- **Go**: Composition is explicit and visible
- **Go**: No virtual methods - simpler resolution
- **Go**: Can't accidentally break parent class
- **Go**: Simpler mental model

## Key Advantages to Emphasize
- **Explicit**: You can see what's embedded
- **Safe**: Compile-time conflict detection
- **Simple**: No complex inheritance hierarchies
- **Flexible**: Mix and match as needed

## Recap What Was Covered
- Basic composition (explicit fields)
- Struct embedding (anonymous fields)
- Field and method promotion
- Multiple embedding and conflicts
- Embedding with interfaces
- Practical patterns (mixins, decorators)
- When to use embedding vs composition
- Common pitfalls

## Preview Next Topics
- Interfaces in depth (tutorial 5)
- Polymorphism in Go
- Type assertions and switches
- Error handling patterns

## Practice Recommendations
- **Easy**: Shape hierarchy using embedding
- **Medium**: Notification system with different notifiers
- **Challenge**: Plugin system with common functionality

## Engagement
- "What surprised you about Go's approach?"
- "How does this compare to languages you know?"
- "Try building something with embedding"

## Closing
- "You've learned Go's composition model"
- "Embedding is powerful but use it wisely"
- "Next: Interfaces - Go's polymorphism mechanism"
