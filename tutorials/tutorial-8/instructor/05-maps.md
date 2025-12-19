# Instructor Notes: Maps

## Teaching Techniques
- Show map creation and initialization
- Demonstrate zero value (nil map)
- Show comma-ok pattern
- Show iteration
- Emphasize: "Maps are reference types"

## Key Emphasis
- **Zero value**: nil map (cannot write to)
- **Comma-ok**: Check if key exists
- **Iteration**: Order is random
- **Reference type**: Maps are references

## Common Questions
- "What's the zero value of a map?" - nil (cannot write to)
- "How do I check if key exists?" - Comma-ok pattern
- "Is iteration order guaranteed?" - No, random order

## Engagement
- "Notice how nil maps behave differently"
- "Comma-ok is the idiomatic way to check existence"
- "Map iteration order is intentionally random"

## Gotchas
- Writing to nil map = panic
- Iteration order is random
- Maps are not safe for concurrent writes

## Real-World Context
- Maps used extensively in Go
- Essential for many data structures
- Understanding behavior prevents bugs

## Transition
- "Let's see common map patterns..."
