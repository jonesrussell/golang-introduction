# Instructor Notes: Performance

## Teaching Techniques
- Show performance considerations
- Demonstrate when to pre-allocate
- Show benchmark examples
- Emphasize: "Understand when optimization matters"

## Key Topics
- Pre-allocate slices when size known
- Pre-allocate maps when size known
- Avoid unnecessary allocations
- Use capacity hints

## Key Emphasis
- **Pre-allocation**: Use make() with size/capacity
- **Capacity hints**: Help avoid reallocations
- **Measure**: Don't optimize blindly
- **Trade-offs**: Memory vs performance

## Common Questions
- "When should I pre-allocate?" - When you know the size
- "Does it matter?" - For large datasets, yes
- "How do I measure?" - Use benchmarks

## Engagement
- "Performance matters, but measure first"
- "Pre-allocation can make a big difference"
- "Don't optimize prematurely"

## Real-World Context
- Performance matters in production
- Pre-allocation is common pattern
- Measure before optimizing

## Transition
- "Let's put it all together in a practical example..."
