# Instructor Notes: Reference Types

## Teaching Techniques
- Clarify the confusion: "reference types" vs pointers
- Show what actually happens (header copy)
- Demonstrate when you still need pointers
- Explain the return-value pattern

## Key Concept
- **"Reference types"**: Slices, maps, channels
- **What it means**: Header is copied, data is shared
- **Not true references**: Still pass by value (the header)

## Demo Flow
1. Show slice modification: `s[0] = 999` works
2. Show append: `append(s, 100)` doesn't affect caller
3. Explain: Header copied, backing array shared
4. Show pointer to slice: `*s = append(*s, 100)` works
5. Show return-value pattern: Often cleaner

## Key Emphasis
- **Modifying contents**: No pointer needed (slices, maps)
- **Replacing entire thing**: Pointer needed (or return value)
- **Return value pattern**: Often cleaner than `*[]Type`
- **Know the difference**: Modifying vs replacing

## Common Questions
- "Do I need a pointer for slices?" - Only if replacing the slice itself
- "Why does append need reassignment?" - Creates new slice header
- "What about maps?" - Similar - modify contents OK, replace needs pointer

## Engagement
- "Notice how modifying slice elements works without pointer"
- "But appending doesn't affect the caller - why?"
- "The return-value pattern is idiomatic Go"

## Real-World Context
- Most Go code uses return values for slices
- `append()` returns new slice - reassign it
- Maps are similar - modify contents, pointer to replace

## Transition
- "Let's summarize best practices..."
