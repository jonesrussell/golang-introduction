# Instructor Notes: Printf Problems

## Teaching Techniques
- Show printf examples
- Explain why it's hard to parse
- Show search difficulties
- Emphasize: "Structured logging solves these problems"

## Key Emphasis
- **Hard to parse**: String formatting is inconsistent
- **Hard to search**: Can't filter by fields
- **No levels**: Can't filter by severity
- **No context**: Missing structured data

## Common Questions
- "What's wrong with printf?" - Hard to parse and search
- "Can't I just grep?" - Yes, but structured is better
- "Why does it matter?" - Production debugging needs structure

## Engagement
- "Notice how hard it is to search these logs"
- "Structured logging makes this easy"
- "This is why we need structured logging"

## Real-World Context
- Printf logging is common but problematic
- Structured logging is standard
- Essential for production systems

## Transition
- "Now let's see Zap basics..."
