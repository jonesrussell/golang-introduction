# Instructor Notes: Structs vs Other Types

## Teaching Techniques
- Quick decision guide
- Show examples of each choice
- Emphasize: "It depends on your use case"

## Decision Framework
- **Struct**: Fixed schema, type safety, methods
- **Map**: Dynamic keys, uniform values
- **Interface**: Behavior contracts
- **Slice**: Ordered, same type

## Key Questions to Ask
- "Is the structure fixed?" → Struct
- "Do keys change at runtime?" → Map
- "Do I need behavior?" → Interface
- "Is it just a list?" → Slice

## Examples to Show
- Address → Struct (fixed fields)
- User scores → Map (dynamic keys)
- Writer → Interface (behavior)
- Shopping cart → Slice (ordered list)

## Engagement
- "When would you use a map instead of a struct?"
- "Can a struct implement an interface?" (preview for tutorial 5)
- "What if you need both?" (struct with map field)

## Real-World Context
- Most Go code uses structs for data
- Maps for configuration, caches
- Interfaces for polymorphism
- Slices for collections

## Transition
- "You now have the foundation - let's wrap up"
