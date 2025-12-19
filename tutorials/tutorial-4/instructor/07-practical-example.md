# Instructor Notes: Practical Example

## Teaching Techniques
- Build incrementally, don't paste all at once
- Explain each pointer use as you add it
- Show nil checks in action
- Run code frequently to demonstrate

## Build Order
1. Define Node struct with `*Node` field (pointer to next)
2. Define LinkedList struct with `*Node` head
3. Create constructor (returns `*LinkedList`)
4. Add Append method (pointer receiver, modifies list)
5. Add Prepend method
6. Add Find method (returns `*Node` or nil)
7. Add Delete method
8. Add Display method
9. Show it all working together

## Live Commentary
- "Notice the `*Node` field - that's a pointer to the next node"
- "Append uses a pointer receiver because it modifies the list"
- "Find returns `*Node` - nil means not found"
- "Always check for nil before traversing"

## Things to Emphasize
- Pointer fields for linked structures
- Pointer receivers for methods that modify
- Nil checks before dereferencing
- Returning pointers (nil for "not found")

## Engagement
- "What happens if we try to append to a nil list?"
- "Let's trace through the Find method"
- Challenge: "Add a method to reverse the list"

## Variations to Mention
- Could make it doubly linked (prev pointer too)
- Could add indexing
- Could add sorting

## Common Mistakes to Watch For
- Forgetting nil checks
- Not using pointer receiver for modification
- Returning nil without documenting it
