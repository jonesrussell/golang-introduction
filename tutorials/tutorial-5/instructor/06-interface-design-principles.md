# Instructor Notes: Interface Design Principles

## Teaching Techniques
- Show good vs bad examples
- Explain the "why" behind each principle
- Emphasize: "These principles make code better"

## Key Principles to Cover

### 1. Keep Interfaces Small
- **Why**: Easier to implement, more implementations possible
- **Show**: Large interface vs small interfaces
- **Example**: Repository broken into Creator, Finder, Updater

### 2. Accept Interfaces, Return Structs
- **Why**: Flexibility for callers, concrete types for returns
- **Show**: Function parameters use interfaces
- **Show**: Functions return concrete types

### 3. Define Interfaces at Point of Use
- **Why**: Consumer defines what they need
- **Show**: Handler defines UserGetter interface
- **Show**: Repository satisfies it without knowing about Handler

### 4. Interface Segregation
- **Why**: Don't force unused methods
- **Show**: Bad (Animal with Walk/Swim/Fly)
- **Show**: Good (separate Walker, Swimmer, Flyer)

## Key Emphasis
- **Small = powerful**: More implementations possible
- **Consumer defines**: Interface at point of use
- **Producer returns**: Concrete types
- **No unused methods**: Segregate interfaces

## Common Questions
- "How small is small?" - 1-3 methods is ideal
- "Where do I define interfaces?" - Where they're used
- "Why return structs?" - More flexibility for callers

## Engagement
- "Notice how small interfaces enable more implementations"
- "The consumer defines what they need - powerful pattern"
- "This is idiomatic Go interface design"

## Real-World Context
- Standard library follows these principles
- Production code uses these patterns
- Code reviews check for these principles

## Transition
- "Let's put it all together in a practical example..."
