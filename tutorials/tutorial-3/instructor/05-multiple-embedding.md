# Instructor Notes: Multiple Embedding

## Teaching Techniques
- Show power of multiple embedding
- Demonstrate name conflicts (compile-time errors!)
- Show resolution: explicit access
- Emphasize: "Explicit is better than implicit"

## Demo Flow
1. Show Timestamps and Metadata structs
2. Show Article embedding both
3. Demonstrate all fields/methods promoted
4. Show conflict example (A and B both have Name)
5. Show compilation error - "ambiguous selector"
6. Show resolution: explicit access `c.A.Name`

## Key Emphasis
- **Multiple embedding**: Can embed as many as needed
- **Name conflicts**: Compile-time error (safe!)
- **Explicit resolution**: Must specify which embedded type
- **Outer wins**: Outer struct fields shadow embedded ones

## Common Questions
- "What if two embedded types have the same method?" - Compile error, must be explicit
- "Can the outer struct override?" - Yes! Outer struct fields/methods win
- "How do I know which one to use?" - Be explicit when ambiguous

## Engagement
- "Notice the compile error - Go catches this at compile time!"
- "What happens if Article also has a Name field?"
- "This is safer than inheritance - conflicts are explicit"

## Gotchas
- Conflicts cause compile errors (good!)
- Must be explicit when ambiguous
- Outer struct fields shadow embedded ones
- Can still access all embedded types explicitly

## Transition
- "Now let's see how embedding works with interfaces..."
