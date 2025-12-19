# Instructor Notes: Premature Optimization

## Teaching Techniques
- Show bad example (optimizing too early)
- Explain why it's wrong
- Show the fix (measure first)
- Emphasize: "Measure, then optimize"

## Key Emphasis
- **Premature optimization**: Optimizing before measuring
- **Measure first**: Use benchmarks
- **Optimize hot paths**: Only optimize what matters
- **Readability**: Don't sacrifice readability

## Common Questions
- "When should I optimize?" - After measuring and finding bottlenecks
- "How do I measure?" - Use benchmarks, profiling
- "What if it's slow?" - Measure first, then optimize

## Engagement
- "Notice how this optimization makes code harder to read"
- "Measure first, then optimize"
- "Readability matters"

## Real-World Context
- Premature optimization is common
- Makes code harder to maintain
- Measure first, optimize later

## Transition
- "Now let's see mutex misuse..."
