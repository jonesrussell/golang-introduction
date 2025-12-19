# Instructor Notes: Slice Patterns

## Teaching Techniques
- Show common patterns
- Demonstrate best practices
- Show performance considerations
- Emphasize: "Patterns make code clearer"

## Patterns to Cover
- Pre-allocate with make when size known
- Use nil slices for empty collections
- Avoid unnecessary allocations
- Use three-index slicing for capacity control

## Key Emphasis
- **Pre-allocation**: Use make() when size known
- **Nil slices**: Perfect for empty collections
- **Capacity**: Pre-allocate to avoid reallocations
- **Three-index slicing**: Control capacity

## Common Questions
- "When do I pre-allocate?" - When you know the size
- "What's the difference from nil slice?" - Nil slice is empty, allocated slice has capacity
- "Why three-index slicing?" - To control capacity

## Engagement
- "These patterns make code more efficient"
- "Notice how pre-allocation helps performance"
- "Nil slices are idiomatic Go"

## Real-World Context
- These patterns used in production code
- Performance matters for large datasets
- Idiomatic Go uses these patterns

## Transition
- "Now let's see maps - key-value stores..."
