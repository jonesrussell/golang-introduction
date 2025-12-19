# Instructor Notes: When to Use Pointers

## Teaching Techniques
- Provide clear decision framework
- Show performance implications
- Explain each use case with examples
- Use the decision flowchart

## Decision Framework

### Use Pointers When:
1. **Need to modify**: Function/method must change the original
2. **Large structs**: Avoid copying expensive data
3. **Nil is meaningful**: Optional values, "not found" cases
4. **Consistency**: If any method uses pointer, all should

### Use Values When:
1. **Small, immutable**: Point, Color, simple types
2. **Want immutability**: Guarantee the original won't change
3. **Slices/maps/channels**: Already "reference-like" (but see gotchas)

## Key Emphasis
- **Mutation**: Pointers for modification
- **Performance**: Pointers for large structs (avoid copying)
- **Nil semantics**: Pointers when nil is a valid state
- **Consistency**: Pick one style per type

## Common Questions
- "How large is 'large'?" - 3-4+ fields, or contains large arrays
- "What about small structs?" - Values are fine, copying is cheap
- "Can I mix value and pointer receivers?" - Yes, but not recommended (be consistent)

## Engagement
- "Let's analyze: should this use a pointer?"
- "What if the struct is huge? Copying would be expensive"
- "Notice how nil can mean 'use default'"

## Real-World Context
- Most production code uses pointer receivers
- Constructors return pointers
- API design: nil often means "optional" or "not set"

## Transition
- "Let's watch out for common gotchas..."
