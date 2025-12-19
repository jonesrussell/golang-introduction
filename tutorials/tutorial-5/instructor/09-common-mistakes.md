# Instructor Notes: Common Mistakes

## Teaching Techniques
- Show bad examples first
- Explain why they're bad
- Show the fix
- Emphasize: "Learn from these mistakes"

## Mistakes to Cover

### 1. Interface Pollution
- **Problem**: Creating interfaces for single implementations
- **Fix**: Only create interface when you need it
- **Show**: Just use the concrete type!

### 2. Premature Abstraction
- **Problem**: Creating interface before you need it
- **Fix**: Wait until you have multiple implementations or need to mock
- **Show**: YAGNI principle (You Aren't Gonna Need It)

### 3. Returning Interface
- **Problem**: Caller doesn't know concrete type
- **Fix**: Return concrete type, accept interface
- **Show**: More flexibility for callers

### 4. Accepting Concrete When Interface Would Work
- **Problem**: Only works with specific type
- **Fix**: Accept interface (like io.Reader)
- **Show**: More flexible and testable

### 5. interface{} Abuse
- **Problem**: Using interface{} when type is known
- **Fix**: Use actual type or specific interface
- **Show**: Type safety matters

### 6. Large Interfaces
- **Problem**: Hard to implement, forces unused methods
- **Fix**: Small, focused interfaces
- **Show**: Compose when needed

## Key Emphasis
- **Don't over-abstract**: Wait until you need it
- **Return concrete**: More flexibility
- **Accept interfaces**: More flexible and testable
- **Keep small**: Easier to implement

## Engagement
- "What's wrong with creating an interface for every type?"
- "Why is returning an interface bad?"
- "Notice how small interfaces are easier to work with"

## Real-World Context
- These mistakes are common in Go codebases
- Code reviews catch these issues
- Following principles makes code better

## Transition
- "Let's wrap up and see what's next..."
