# Instructor Notes: What Are Interfaces?

## Teaching Techniques
- Show the "magic" of implicit implementation
- Demonstrate polymorphism in action
- Show multiple types satisfying same interface
- Emphasize: "No `implements` keyword needed!"

## Demo Flow
1. Define Speaker interface (one method)
2. Show Dog implementing it (implicitly)
3. Show Cat implementing it
4. Show Robot implementing it
5. Demonstrate: All can be used as Speaker
6. Show slice of interface type

## Key Emphasis
- **Interface = method signatures**: Just the contract
- **Implicit implementation**: If you have the methods, you satisfy it
- **Polymorphism**: One function works with multiple types
- **No coupling**: Types don't need to know about interfaces

## Common Questions
- "How does Go know Dog implements Speaker?" - It checks at compile time
- "Do I need to declare it?" - No! It's implicit
- "What if I forget a method?" - Compile error (helpful!)

## Engagement
- "Notice - no `implements` keyword anywhere!"
- "All three types can be used as Speaker - that's polymorphism"
- "This is Go's way of achieving code reuse"

## Real-World Context
- Standard library uses this pattern extensively
- Enables writing generic functions
- Foundation for dependency injection

## Transition
- "But there are rules about how interfaces are satisfied..."
