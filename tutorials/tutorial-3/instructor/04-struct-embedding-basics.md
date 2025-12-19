# Instructor Notes: Struct Embedding Basics

## Teaching Techniques
- Show the "magic" of field promotion
- Demonstrate method promotion
- Contrast with explicit composition
- Show the syntax: anonymous field (no name)

## Demo Flow
1. Show User struct with methods
2. Show Admin embedding User (anonymous field)
3. Demonstrate field promotion: `admin.Username` (not `admin.User.Username`)
4. Demonstrate method promotion: `admin.GetDisplayName()`
5. Show you can still access explicitly: `admin.User.Email`

## Key Emphasis
- **Anonymous field**: No field name = embedding
- **Field promotion**: Embedded fields accessible directly
- **Method promotion**: Embedded methods accessible directly
- **Still accessible explicitly**: Can use `admin.User.Username` too

## Common Questions
- "Why would I want promotion?" - Convenience, code reuse
- "Can I still access the embedded struct?" - Yes! `admin.User`
- "What's the difference from inheritance?" - No virtual methods, explicit

## Engagement
- "Notice how `admin.Username` works - that's promotion!"
- "What happens if Admin also has a Username field?" (preview conflict)
- "This is Go's way of code reuse without inheritance"

## Gotchas
- Anonymous field syntax is just no field name
- Promotion only works for exported fields/methods
- Can still access embedded type explicitly

## Transition
- "What if we embed multiple structs? Let's see..."
