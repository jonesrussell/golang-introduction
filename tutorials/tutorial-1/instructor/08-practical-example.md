# Instructor Notes: Practical Example

## Teaching Techniques
- Build incrementally, don't paste all at once
- Explain each step as you type
- Run after each addition to show progress

## Build Order
1. Declare student name and scores slice
2. Write the loop to calculate total
3. Calculate average (show type conversion!)
4. Add switch for grade determination
5. Add if for pass/fail
6. Add Printf statements

## Live Commentary
- "First, let's set up our data..."
- "Now we need to loop through scores..."
- "Notice we need float64 for division..."
- "Switch is perfect for grade ranges..."

## Things to Emphasize
- Type conversion: `float64(total) / float64(len(scores))`
- Switch without expression for ranges
- Printf formatting: `%.2f` for two decimals

## Engagement
- "What grade would you expect for these scores?"
- "Let's change the scores and see what happens"
- Challenge: "Add a score and predict the new average"

## Variations to Mention
- Could use a map for grades
- Could read from user input
- Could handle multiple students
