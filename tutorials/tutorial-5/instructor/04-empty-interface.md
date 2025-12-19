# Instructor Notes: Empty Interface

## Teaching Techniques
- Show `interface{}` accepts anything
- Demonstrate type assertions (safe and unsafe)
- Show type switch pattern
- Warn: "Use sparingly - prefer specific types"

## Demo Flow
1. Show `interface{}` accepts any type
2. Show type assertion: `i.(string)`
3. Show safe assertion: `s, ok := i.(string)`
4. Show type switch pattern
5. Show when to use (and not use) empty interface

## Key Emphasis
- **`interface{}` / `any`**: Accepts any type (zero methods)
- **Type assertion**: Extract concrete type from interface
- **Comma-ok pattern**: Safe assertion (prevents panic)
- **Type switch**: Handle multiple types elegantly
- **Use sparingly**: Prefer specific interfaces

## Common Questions
- "When do I use `interface{}`?" - When type is truly unknown
- "What's the difference from `any`?" - `any` is alias (Go 1.18+)
- "How do I get the type back?" - Type assertion or type switch

## Engagement
- "Notice how `interface{}` accepts literally anything"
- "Type assertion is how we get the concrete type back"
- "Type switch is cleaner than multiple if-else"

## Real-World Context
- Used in JSON unmarshaling (unknown structure)
- Used in generic containers (before Go 1.18 generics)
- Used in variadic functions (fmt.Printf)

## Transition
- "Let's see interfaces in the standard library..."
