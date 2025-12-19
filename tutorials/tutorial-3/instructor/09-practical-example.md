# Instructor Notes: Practical Example

## Teaching Techniques
- Build incrementally, don't paste all at once
- Explain each mixin as you add it
- Show how they compose together
- Run code frequently to show it working

## Build Order
1. Create mixins (Timestamps, Identifiable)
2. Create base Person type using mixins
3. Create Employee embedding Person
4. Create Manager embedding Employee
5. Add methods to each level
6. Show how methods are promoted
7. Demonstrate the complete system

## Live Commentary
- "First, let's create our reusable mixins..."
- "Notice how Person uses both mixins..."
- "Employee embeds Person - gets all Person's functionality..."
- "Manager embeds Employee - gets everything from both levels..."

## Things to Emphasize
- Mixin pattern for reusable functionality
- Embedding creates type hierarchy without inheritance
- Method promotion at each level
- Can still access embedded types explicitly

## Engagement
- "What methods does Manager have access to?"
- "Let's add a method to Person - watch it appear in Manager!"
- Challenge: "Add a Department mixin and use it"

## Variations to Mention
- Could add more mixins (Auditable, SoftDeletable)
- Could add more employee types (Contractor, Intern)
- Could add database persistence layer

## Common Mistakes to Watch For
- Forgetting that embedding promotes methods
- Not understanding method resolution order
- Over-embedding (too many levels)
