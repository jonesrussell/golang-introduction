# Instructor Notes: Common Pitfalls

## Teaching Techniques
- Show each pitfall clearly
- Explain why it's wrong
- Show the fix
- Emphasize: "These are common mistakes"

## Pitfalls to Cover
- Mutating slices (sharing underlying array)
- Writing to nil maps
- Assuming map iteration order
- Not pre-allocating when size known
- Slicing beyond capacity

## Key Emphasis
- **Slice mutations**: Slicing shares underlying array
- **Nil maps**: Cannot write to nil map
- **Iteration order**: Random, don't assume
- **Pre-allocation**: Use when size known

## Common Questions
- "Why did my slice change?" - Slicing shares underlying array
- "Why did I get a panic?" - Writing to nil map
- "Why is my map order different?" - Iteration is random

## Engagement
- "Watch this - it looks fine but has a bug"
- "This is why understanding internals matters"
- "These pitfalls are common in production code"

## Real-World Context
- These pitfalls are common
- Understanding prevents bugs
- Code reviews catch these

## Transition
- "Let's wrap up with best practices..."
