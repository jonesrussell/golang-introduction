# Instructor Notes: Concurrency Patterns

## Teaching Techniques
- Show real-world patterns
- Explain when to use each
- Build examples incrementally
- Connect to production code

## Patterns to Cover

### Worker Pool
- **When**: Fixed number of workers processing jobs
- **Show**: Jobs channel, results channel, worker goroutines
- **Use case**: Rate limiting, resource management

### Pipeline
- **When**: Chain of processing stages
- **Show**: Each stage is a goroutine with channels
- **Use case**: Data transformation pipelines

### Fan-out/Fan-in
- **When**: Distribute work, collect results
- **Show**: Multiple workers, single collector
- **Use case**: Parallel processing

### Semaphore
- **When**: Limit concurrent operations
- **Show**: Buffered channel as semaphore
- **Use case**: Rate limiting, resource limits

## Key Emphasis
- **Worker pool**: Fixed workers, job queue
- **Pipeline**: Chain stages with channels
- **Semaphore**: Limit concurrency with buffered channel
- **Choose pattern**: Based on problem structure

## Common Questions
- "When do I use worker pool?" - When you need to limit workers
- "What's the difference from pipeline?" - Worker pool has fixed workers, pipeline has stages
- "How does semaphore work?" - Buffered channel limits concurrent operations

## Engagement
- "Notice how patterns compose together"
- "These patterns are everywhere in production code"
- "Worker pool is one of the most common patterns"

## Real-World Context
- These patterns used extensively in production
- Standard library uses them
- Foundation for building concurrent systems

## Transition
- "Now let's see the sync package for synchronization..."
