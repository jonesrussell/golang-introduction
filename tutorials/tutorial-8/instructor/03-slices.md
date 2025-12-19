# Instructor Notes: Slices

## Teaching Techniques
- Show slice header structure (pointer, len, cap)
- Demonstrate slicing operations
- Show append and growth behavior
- Show copy function
- Emphasize: "Slices are references to underlying arrays"

## Key Emphasis
- **Slice header**: Pointer, length, capacity
- **Slicing**: Creates new slice sharing underlying array
- **Append**: May allocate new array if capacity exceeded
- **Copy**: Use for independent copies

## Common Questions
- "What happens when I slice?" - New slice shares underlying array
- "When does append allocate?" - When capacity is exceeded
- "How do I copy a slice?" - Use copy() function

## Engagement
- "Notice how slicing shares the underlying array"
- "This is why understanding internals matters"
- "Append is smart - it grows when needed"

## Gotchas
- Slicing shares underlying array (mutations visible)
- Append may allocate new array
- Zero-length slice vs nil slice

## Real-World Context
- Slices used everywhere in Go
- Understanding internals prevents bugs
- Essential for efficient code

## Transition
- "Let's see common slice patterns..."
