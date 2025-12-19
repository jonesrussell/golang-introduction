# Instructor Notes: What Are Structs?

## Teaching Techniques
- Start with the "messy" example (separate variables)
- Show the problem: "How do we pass all this data around?"
- Then reveal the struct solution
- Emphasize type safety

## Demo Flow
1. Show the messy code with separate variables
2. Ask: "What if we need to pass this to a function?"
3. Introduce the struct syntax
4. Show how it groups related data
5. Explain exported vs unexported fields

## Key Emphasis
- **Type safety**: Can't accidentally mix up fields
- **Organization**: Related data stays together
- **Exported fields**: Capital letter = public (can be accessed from other packages)
- **Unexported fields**: Lowercase = private (package-only)

## Common Questions
- "Why capital letters?" - Exported vs unexported
- "Can I have methods?" - Yes! (Coming in section 6)
- "Are they like classes?" - Similar, but simpler (no inheritance)

## Engagement
- "Notice how the struct makes the relationship clear"
- "What happens if we try to access a lowercase field from another package?"

## Transition
- "Now let's learn how to define structs properly"
