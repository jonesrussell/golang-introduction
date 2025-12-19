# Instructor Notes: Pointer Gotchas

## Teaching Techniques
- Show each gotcha with bad example first
- Explain why it's wrong
- Show the fix
- Emphasize: "Learn from these mistakes"

## Gotchas to Cover

### 1. Nil Pointer Dereference
- **Problem**: Panic if you dereference nil
- **Fix**: Always check `if ptr != nil`
- **Show**: Safe pattern

### 2. Loop Variable Address
- **Problem**: All pointers point to same address
- **Fix**: Shadow variable: `n := n`
- **Show**: Bad vs good example

### 3. Uninitialized Pointer Fields
- **Problem**: Struct with `*Type` field is nil
- **Fix**: Initialize before use
- **Show**: Safe initialization

### 4. Returning Local Pointers
- **Surprise**: This is SAFE in Go!
- **Explain**: Escape analysis moves to heap
- **Contrast**: Unlike C, this works

## Key Emphasis
- **Always check nil**: Before dereferencing
- **Loop variable gotcha**: Common mistake, easy fix
- **Go is safe**: Can return local pointers (unlike C)
- **Initialize fields**: Pointer fields start as nil

## Common Questions
- "Why can I return a local pointer?" - Go's escape analysis
- "What's the loop variable problem?" - All iterations share same address
- "How do I avoid nil panics?" - Always check before use

## Engagement
- "Watch this - it looks wrong but it's safe in Go!"
- "This gotcha trips up many developers"
- "Go's garbage collector makes this safe"

## Best Practices
- Check for nil before using
- Use constructors to ensure initialization
- Document when nil is valid
- Shadow loop variables when taking addresses

## Transition
- "Let's put it all together in a practical example..."
