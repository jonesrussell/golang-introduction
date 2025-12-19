# Instructor Notes: Map Patterns

## Teaching Techniques
- Show common patterns
- Demonstrate best practices
- Show initialization patterns
- Emphasize: "Patterns make code clearer"

## Patterns to Cover
- Initialize with make() or literal
- Use comma-ok for existence checks
- Use maps for sets (map[T]bool)
- Use maps for counting (map[T]int)

## Key Emphasis
- **Initialization**: Use make() or literal
- **Existence checks**: Comma-ok pattern
- **Sets**: map[T]bool or map[T]struct{}
- **Counting**: map[T]int

## Common Questions
- "How do I make a set?" - map[T]bool or map[T]struct{}
- "Why struct{} for sets?" - Zero memory overhead
- "How do I count occurrences?" - map[T]int

## Engagement
- "Maps are versatile - sets, counters, lookups"
- "Notice how struct{} saves memory"
- "These patterns are everywhere in Go code"

## Real-World Context
- These patterns used extensively
- Maps are versatile data structures
- Idiomatic Go uses these patterns

## Transition
- "Let's see performance considerations..."
