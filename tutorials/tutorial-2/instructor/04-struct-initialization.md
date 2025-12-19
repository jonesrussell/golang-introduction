# Instructor Notes: Struct Initialization

## Teaching Techniques
- Show zero values first (surprises many!)
- Emphasize named fields (most readable)
- Warn against positional initialization
- Show constructor pattern (idiomatic Go)

## Demo Flow
1. Zero value initialization - show all fields get zero values
2. Named field initialization (RECOMMENDED)
3. Partial initialization (remaining get zero values)
4. Show positional (but warn against it)
5. Pointer initialization with `&`
6. Constructor function pattern

## Key Emphasis
- **Zero values**: Every field gets initialized (no nil pointers for primitives!)
- **Named fields**: Most readable, order doesn't matter
- **Trailing comma**: Required on multi-line (Go convention)
- **Constructors**: Idiomatic pattern for validation and defaults

## Common Questions
- "Why do I need the trailing comma?" - Go convention, makes diffs cleaner
- "What's the zero value for my struct?" - Each field gets its type's zero value
- "When do I use `&`?" - When you need a pointer (large structs, mutating)

## Engagement
- "Notice how zero values make structs safe to use immediately"
- "Positional initialization works, but what if we add a field?"
- Challenge: "Create a constructor that validates the email"

## Gotchas
- Positional initialization breaks if fields are reordered
- Zero value for string is `""`, not `nil`
- `new()` returns a pointer with zero values

## Transition
- "Now let's see how to access and modify these fields"
