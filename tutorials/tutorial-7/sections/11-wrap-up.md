# Next Steps & Wrap-up

**Duration:** 2-3 minutes

## Recap What Was Covered:
- [Goroutines](https://go.dev/ref/spec#Go_statements) and the [`go` keyword](https://go.dev/ref/spec#Go_statements)
- [Channels](https://go.dev/ref/spec#Channel_types) for communication
- [Select](https://go.dev/ref/spec#Select_statements) for multiple channels
- Concurrency patterns
- [sync package](https://pkg.go.dev/sync) (WaitGroup, Mutex, Once)
- [Context](https://pkg.go.dev/context) for cancellation
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
- [Go Blog: "Go Concurrency Patterns"](https://go.dev/blog/pipelines): go.dev/blog/pipelines
- [Go Blog: "Pipelines and cancellation"](https://go.dev/blog/pipelines): go.dev/blog/pipelines
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency): go.dev/doc/effective_go#concurrency

## Key teaching points:
- Go makes concurrency accessible
- [Channels](https://go.dev/ref/spec#Channel_types) provide safe communication
- [Context](https://pkg.go.dev/context) enables proper cancellation
- Always test with the [race detector](https://go.dev/doc/articles/race_detector)
