# Instructor Notes: Practical Example

## Teaching Techniques
- Build incrementally, don't paste all at once
- Explain each pattern as you add it
- Show error handling with goroutines
- Demonstrate proper cleanup

## Build Order
1. Define Result struct
2. Define Scraper struct
3. Create worker function (with context)
4. Create Scrape method (worker pool)
5. Show rate limiting
6. Show proper channel closing
7. Show WaitGroup usage
8. Show context cancellation
9. Demonstrate it all working

## Live Commentary
- "First, let's set up our worker pool..."
- "Notice how we use context for cancellation..."
- "Rate limiting prevents overwhelming the server..."
- "WaitGroup ensures all workers finish..."

## Things to Emphasize
- Worker pool pattern
- Context for cancellation
- Rate limiting with time.Sleep
- Proper channel closing
- Error handling in concurrent code

## Engagement
- "What happens if we don't close the channel?"
- "How does context cancellation work here?"
- Challenge: "Add retry logic for failed requests"

## Variations to Mention
- Could add more sophisticated rate limiting
- Could add result caching
- Could add progress reporting
- Could add different worker types

## Common Mistakes to Watch For
- Forgetting to close channels
- Not checking context.Done()
- Goroutine leaks
- Race conditions
