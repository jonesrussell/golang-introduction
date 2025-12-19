# Instructor Notes: The Problem

## Teaching Techniques
- Show the "bad" way first (duplication)
- Make the problem obvious
- Then promise: "Embedding solves this elegantly"

## Demo Flow
1. Show Person struct
2. Show Employee struct (duplicates Person fields)
3. Show Manager struct (duplicates even more)
4. Point out: "Look at all this duplication!"
5. Ask: "How do we share code without inheritance?"

## Key Emphasis
- **Code duplication**: Same fields repeated
- **Maintenance nightmare**: Change Person, update everywhere
- **No inheritance**: Can't extend classes
- **Solution coming**: Embedding solves this

## Engagement
- "Notice how much code we're repeating"
- "What if we need to change Person?"
- "Go doesn't have inheritance - so what's the solution?"

## Transition
- "Let's see how composition works first"
