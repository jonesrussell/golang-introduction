# Instructor Notes: Error Handling Patterns

## Teaching Techniques
- Show the ubiquitous `if err != nil` pattern
- Demonstrate early returns (reduces nesting)
- Show different handling strategies
- Warn: "Never ignore errors!"

## Demo Flow
1. Show basic pattern: check, return
2. Show early returns (clean code)
3. Show handle and continue (non-critical errors)
4. Show multiple error checks
5. Show cleanup with defer
6. Show the BAD pattern (ignoring errors)

## Key Emphasis
- **Always check errors**: Every function that returns error
- **Early returns**: Reduces nesting, cleaner code
- **Never ignore**: `_` is almost always wrong
- **Defer cleanup**: Ensures cleanup even on error
- **Log vs return**: Non-critical vs critical errors

## Common Questions
- "Do I have to check every error?" - Yes! That's Go's philosophy
- "What if I want to ignore an error?" - You probably don't (show why)
- "When do I log vs return?" - Log non-critical, return critical

## Engagement
- "Notice how every error is checked - this is Go's way"
- "Early returns make code much cleaner"
- "What happens if we ignore this error?" (show the panic)

## Gotchas
- Ignoring errors with `_` is dangerous
- Defer ensures cleanup even if function returns early
- Log non-critical errors, return critical ones

## Real-World Context
- This pattern is everywhere in Go code
- Code reviews check for error handling
- Missing error checks are bugs waiting to happen

## Transition
- "Now let's see how to add context to errors..."
