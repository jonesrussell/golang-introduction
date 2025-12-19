# Instructor Notes: Methods on Structs

## Teaching Techniques
- Start with value receivers (simpler)
- Then show pointer receivers (for modification)
- Emphasize the decision: value vs pointer
- Show Go's automatic conversion (it's magic!)

## Demo Flow
1. Value receiver method (Area) - doesn't modify
2. Pointer receiver method (Scale) - does modify
3. Show automatic conversion (value can call pointer method)
4. BankAccount example - real-world pattern
5. Explain when to use each

## Key Emphasis
- **Value receiver**: `func (r Rectangle)` - gets copy, can't modify
- **Pointer receiver**: `func (r *Rectangle)` - gets pointer, can modify
- **Automatic conversion**: Go handles `(&rect).Method()` and `(*ptr).Method()` automatically
- **Consistency**: Pick one style per type (all pointer or all value)

## Decision Guide
- **Use pointer receiver when:**
  - Method modifies the receiver
  - Struct is large (avoid copying)
  - Consistency (if one method uses pointer, all should)
- **Use value receiver when:**
  - Method only reads data
  - Struct is small
  - You want immutability

## Common Questions
- "Why would I use a value receiver?" - Immutability, small structs
- "Can I mix them?" - Yes, but not recommended (be consistent)
- "What's the performance difference?" - Minimal for small structs

## Engagement
- "Notice how Go automatically converts between values and pointers"
- "What happens if we use a value receiver for a method that modifies?"
- Challenge: "Add a method to check if balance is negative"

## Gotchas
- Value receiver methods can't modify the original
- Pointer receiver methods can modify even if called on a value
- Be consistent within a type!

## Transition
- "Let's put it all together in a practical example"
