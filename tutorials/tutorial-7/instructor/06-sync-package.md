# Instructor Notes: sync Package

## Teaching Techniques
- Show each sync primitive clearly
- Demonstrate when to use each
- Show common patterns
- Emphasize: "Use channels when possible, sync when needed"

## Primitives to Cover

### WaitGroup
- **Purpose**: Wait for multiple goroutines
- **Show**: Add(1), Done(), Wait()
- **Pattern**: Defer wg.Done() in goroutine

### Mutex
- **Purpose**: Protect shared state
- **Show**: Lock/Unlock with defer
- **Pattern**: Always use defer for safety

### RWMutex
- **Purpose**: Multiple readers, single writer
- **Show**: RLock for reads, Lock for writes
- **When**: Reads >> writes

### sync.Once
- **Purpose**: One-time initialization
- **Show**: Singleton pattern
- **When**: Expensive initialization, only once

## Key Emphasis
- **WaitGroup**: For waiting on goroutines
- **Mutex**: For protecting shared state
- **RWMutex**: When reads are common
- **Once**: For initialization
- **Always defer**: Lock/Unlock, Done()

## Common Questions
- "When do I use mutex vs channels?" - Mutex for shared state, channels for communication
- "What's the difference from RWMutex?" - RWMutex allows concurrent reads
- "Why sync.Once?" - Thread-safe one-time initialization

## Engagement
- "Notice how defer ensures we always unlock"
- "RWMutex is perfect for caches"
- "sync.Once is the singleton pattern in Go"

## Real-World Context
- sync package used extensively
- Mutex for shared state protection
- WaitGroup for coordinating goroutines

## Transition
- "Now let's see context for cancellation..."
