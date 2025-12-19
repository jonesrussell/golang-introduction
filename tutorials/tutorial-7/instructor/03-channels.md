# Instructor Notes: Channels

## Teaching Techniques
- Show channels as communication mechanism
- Demonstrate blocking behavior (synchronization)
- Show buffered vs unbuffered
- Show closing channels and range
- Emphasize: "Channels are typed"

## Demo Flow
1. Create channel: `make(chan string)`
2. Send in goroutine: `ch <- "message"`
3. Receive in main: `msg := <-ch` (blocks!)
4. Show buffered channel (doesn't block immediately)
5. Show closing channel
6. Show range over channel
7. Show comma-ok pattern

## Key Emphasis
- **Channels are typed**: `chan int`, `chan string`, etc.
- **Unbuffered**: Synchronizes (send blocks until receive)
- **Buffered**: Can hold N values before blocking
- **Close from sender**: Only sender should close
- **Range until closed**: `for v := range ch`

## Common Questions
- "When do I use buffered vs unbuffered?" - Buffered for throughput, unbuffered for sync
- "Who closes the channel?" - The sender
- "What happens if I send to closed channel?" - Panic!

## Engagement
- "Notice how the receive blocks - that's synchronization!"
- "Buffered channels allow some asynchrony"
- "Range is perfect for consuming channels"

## Gotchas
- Sending to closed channel = panic
- Closing from receiver = wrong (only sender closes)
- Unbuffered channel needs both sender and receiver ready

## Real-World Context
- Channels are Go's primary communication mechanism
- Used extensively in concurrent code
- Enable safe data sharing

## Transition
- "Now let's see how to handle multiple channels..."
