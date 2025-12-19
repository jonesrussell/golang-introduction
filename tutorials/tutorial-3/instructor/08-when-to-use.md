# Instructor Notes: When to Use Embedding vs Composition

## Teaching Techniques
- Provide clear decision framework
- Show good vs bad examples
- Explain the "why" behind each choice
- Use decision flowchart

## Decision Framework

### Use Embedding When:
- Want promoted fields/methods (convenience)
- Implementing mixin behavior
- Embedded type is "base" functionality
- Want interface satisfaction

### Use Composition When:
- Clear "has-a" relationship
- Want explicit access (clarity)
- Might have multiple of same type
- Field has semantic meaning

## Key Emphasis
- **Embedding**: For convenience and code reuse
- **Composition**: For clarity and explicit relationships
- **Don't embed just to save typing**: Clarity matters more
- **Think about the relationship**: "Has-a" vs "Is-a-kind-of"

## Common Questions
- "How do I decide?" - Use the decision framework
- "Can I mix both?" - Yes! Use embedding for mixins, composition for relationships
- "What if I'm not sure?" - Prefer composition (clearer)

## Engagement
- "Let's analyze: should Car embed Engine or have Engine as field?"
- "What's the relationship: has-a or is-a-kind-of?"
- "When in doubt, choose clarity over convenience"

## Anti-Patterns to Highlight
- Embedding just to avoid typing
- Embedding unrelated types
- Over-embedding (too many embedded types)

## Transition
- "Let's put it all together in a complete example..."
