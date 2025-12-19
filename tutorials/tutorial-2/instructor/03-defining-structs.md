# Instructor Notes: Defining Structs

## Teaching Techniques
- Build examples incrementally
- Show field tags early (they're important for JSON)
- Demonstrate nested structs with a real example
- Save anonymous structs for last (they're less common)

## Demo Flow
1. Basic struct (Person) - show simple fields
2. Multiple types (Product) - show variety
3. Field tags (User) - explain JSON tags (preview for later)
4. Nested structs (Employee with Address) - show composition
5. Anonymous struct (config) - show one-off use case

## Key Emphasis
- **Field tags**: Metadata for libraries (JSON, validation, etc.)
- **Nested structs**: Model "has-a" relationships
- **Anonymous structs**: Useful for one-off data (config, test data)
- **Convention**: One struct per logical concept

## Common Questions
- "What are those backticks?" - Field tags (metadata)
- "Can I nest multiple structs?" - Yes! (We'll see more in tutorial 3)
- "When do I use anonymous structs?" - Config, test data, temporary data

## Engagement
- "Notice how nested structs model real relationships"
- "Field tags are how Go libraries know how to serialize your data"
- Challenge: "Try adding a Phone field to the Employee struct"

## Gotchas
- Field tags use backticks, not quotes
- Anonymous structs can't be reused (that's the point!)
- Nested structs are value types (copied when assigned)

## Transition
- "Now let's see how to create instances of these structs"
