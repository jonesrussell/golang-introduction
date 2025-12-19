# Instructor Notes: Common Pitfalls

## Teaching Techniques
- Show bad examples first
- Explain why they're bad
- Show the fix
- Emphasize: "Learn from these mistakes"

## Pitfalls to Cover

### 1. Embedding to Avoid Typing
- **Problem**: Unclear semantic relationship
- **Fix**: Use explicit composition
- **Show**: `person.Street` vs `person.Address.Street`

### 2. Over-Embedding
- **Problem**: Too many embedded types = confusion
- **Fix**: Selective embedding + composition
- **Show**: Clear vs cluttered struct

### 3. Circular Embedding
- **Problem**: Infinite size, won't compile
- **Fix**: Redesign relationship
- **Show**: Compilation error

### 4. Name Conflicts
- **Problem**: Unintentional shadowing
- **Fix**: Different names or explicit composition
- **Show**: Confusion vs clarity

### 5. Exposing Internals
- **Problem**: Embedding sync.Mutex exposes Lock/Unlock
- **Fix**: Private field for internal use
- **Exception**: OK for types that ARE synchronization primitives

## Key Emphasis
- **Be intentional**: Know why you're embedding
- **Prefer clarity**: Composition is often clearer
- **Watch conflicts**: Name conflicts cause compile errors (good!)
- **Don't expose internals**: Use private fields when needed

## Engagement
- "What's wrong with embedding Address here?"
- "Why is this circular embedding bad?"
- "When is it OK to embed sync.Mutex?"

## Best Practices Summary
- Use embedding for mixins
- Use embedding for base functionality
- Use composition for "has-a" relationships
- Keep embedding shallow (1-2 levels)
- Document why you're embedding

## Transition
- "Let's compare Go's approach to other languages..."
