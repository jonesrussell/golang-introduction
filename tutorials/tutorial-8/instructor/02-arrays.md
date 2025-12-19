# Instructor Notes: Arrays

## Teaching Techniques
- Show arrays are value types (copied)
- Demonstrate fixed size constraint
- Show when arrays are used (rarely!)
- Emphasize: "Arrays are the foundation, but slices are what you'll use"

## Key Emphasis
- **Fixed size**: Size is part of type
- **Value type**: Copied on assignment
- **Rarely used**: Slices are preferred
- **When to use**: Fixed-size buffers, specific use cases

## Common Questions
- "When do I use arrays?" - Rarely, mostly for fixed-size buffers
- "Why are arrays copied?" - They're value types, not references
- "What's the difference from slices?" - Arrays have fixed size, slices are dynamic

## Engagement
- "Notice how arrays are copied - that's different from slices"
- "Arrays are rarely used in practice"
- "Slices are built on arrays"

## Real-World Context
- Arrays used in low-level code (buffers)
- Most Go code uses slices
- Understanding arrays helps understand slices

## Transition
- "Now let's see slices - Go's workhorse..."
