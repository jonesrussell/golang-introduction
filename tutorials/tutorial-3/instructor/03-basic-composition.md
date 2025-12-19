# Instructor Notes: Basic Composition

## Teaching Techniques
- Start with explicit composition (named fields)
- Show "has-a" relationships clearly
- Contrast with embedding (coming next)
- Use real-world examples (Person has Address, Car has Engine)

## Demo Flow
1. Show Address struct
2. Show Person with Address field (explicit)
3. Demonstrate field access: `person.Address.Street`
4. Show method access: `person.Address.FullAddress()`
5. Explain: "This is composition - clear and explicit"

## Key Emphasis
- **Explicit composition**: Named fields show clear relationships
- **"Has-a" relationship**: Person HAS an Address
- **Field access**: Must go through field name
- **Clarity**: No ambiguity about what belongs to what

## Common Questions
- "Why not just put fields directly in Person?" - Separation of concerns
- "When do I use this vs embedding?" - When relationship is "has-a"
- "Can I have multiple Addresses?" - Yes! (slice of Addresses)

## Engagement
- "Notice how explicit this is - no magic"
- "What if a Person has multiple addresses?"
- "This is the foundation - embedding builds on this"

## Transition
- "Now let's see embedding - where fields are promoted automatically"
