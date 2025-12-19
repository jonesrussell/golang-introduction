# Instructor Notes: Best Practices

## Teaching Techniques
- Show side-by-side: good vs bad
- Explain WHY each practice matters
- Show real consequences of bad patterns

## Key Patterns to Emphasize

### Constructor Validation
- **Why**: Prevents invalid state
- **Show**: What happens without validation (negative price, empty name)

### Consistent Receiver Types
- **Why**: Predictable behavior
- **Show**: Mixed receivers cause confusion

### Named Field Initialization
- **Why**: Readable, maintainable
- **Show**: Positional breaks when fields reordered

### Export Control
- **Why**: Encapsulation, API design
- **Show**: Public vs private fields

## Common Mistakes to Highlight
1. **Value receiver when modifying** - Show it doesn't work!
2. **Nil pointer dereference** - Show the panic
3. **Comparing non-comparable structs** - Show compilation error
4. **Large struct by value** - Explain performance impact

## Engagement
- "What happens if we forget the pointer receiver?"
- "Why can't we compare these structs?"
- "When would you use a value receiver?"

## Real-World Context
- "In production code, you'll see these patterns everywhere"
- "These practices prevent bugs before they happen"
- "Code reviews will catch these issues"

## Transition
- "Now let's see when to use structs vs other types"
