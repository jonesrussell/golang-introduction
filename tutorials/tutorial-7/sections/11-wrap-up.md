# Next Steps & Wrap-up

**Duration:** 2-3 minutes

## Recap What Was Covered:
- Goroutines and the `go` keyword
- Channels for communication
- Select for multiple channels
- Concurrency patterns
- sync package (WaitGroup, Mutex, Once)
- Context for cancellation
- Web scraper example
- Common pitfalls

## Preview Next Topics:
- Slices and maps internals
- Testing concurrent code
- Production monitoring

## Practice Suggestions:
1. **Easy:** Concurrent file downloader
2. **Medium:** Chat server with multiple clients
3. **Challenge:** Rate-limited API client
4. **Advanced:** Distributed task queue

## Cheat Sheet

```
Goroutine:           go func()
Channel:             ch := make(chan T)
Buffered:            ch := make(chan T, n)
Send:                ch <- value
Receive:             value := <-ch
Close:               close(ch)
Range:               for v := range ch {}
Select:              select { case ... }

WaitGroup:           wg.Add(1), wg.Done(), wg.Wait()
Mutex:               mu.Lock(), mu.Unlock()
Context:             ctx, cancel := context.WithCancel(ctx)
```

## Resources:
- Go Blog: "Go Concurrency Patterns"
- Go Blog: "Pipelines and cancellation"

## Key teaching points:
- Go makes concurrency accessible
- Channels provide safe communication
- Context enables proper cancellation
- Always test with the race detector
