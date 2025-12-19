# Practical Example: Concurrent Web Scraper

**Duration:** 10-12 minutes

## Build Together

A concurrent web scraper with rate limiting and proper cleanup.

```go runnable
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Result holds scraping results
type Result struct {
    URL      string
    Size     int
    Duration time.Duration
    Error    error
}

// Scraper handles concurrent URL fetching
type Scraper struct {
    maxWorkers int
    rateLimit  time.Duration
}

func NewScraper(maxWorkers int, rateLimit time.Duration) *Scraper {
    return &Scraper{
        maxWorkers: maxWorkers,
        rateLimit:  rateLimit,
    }
}

func (s *Scraper) Scrape(ctx context.Context, urls []string) <-chan Result {
    results := make(chan Result)
    urlChan := make(chan string)

    // Create worker pool
    var wg sync.WaitGroup
    for i := 0; i < s.maxWorkers; i++ {
        wg.Add(1)
        go s.worker(ctx, i, urlChan, results, &wg)
    }

    // Send URLs with rate limiting
    go func() {
        ticker := time.NewTicker(s.rateLimit)
        defer ticker.Stop()

        for _, url := range urls {
            select {
            case <-ctx.Done():
                break
            case <-ticker.C:
                urlChan <- url
            }
        }
        close(urlChan)
    }()

    // Close results when all workers done
    go func() {
        wg.Wait()
        close(results)
    }()

    return results
}

func (s *Scraper) worker(ctx context.Context, id int, urls <-chan string, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()

    for url := range urls {
        select {
        case <-ctx.Done():
            return
        default:
            result := s.fetch(url)
            results <- result
        }
    }
}

func (s *Scraper) fetch(url string) Result {
    start := time.Now()
    
    // Simulate fetch with varying times
    time.Sleep(100 * time.Millisecond)
    
    return Result{
        URL:      url,
        Size:     len(url) * 100,  // Simulated size
        Duration: time.Since(start),
    }
}

// Progress tracks scraping progress
type Progress struct {
    mu        sync.Mutex
    total     int
    completed int
}

func (p *Progress) Update() {
    p.mu.Lock()
    defer p.mu.Unlock()
    p.completed++
}

func (p *Progress) Print() {
    p.mu.Lock()
    defer p.mu.Unlock()
    fmt.Printf("\rProgress: %d/%d", p.completed, p.total)
}

func main() {
    fmt.Println("=== Concurrent Web Scraper ===\n")

    urls := []string{
        "https://example.com/page1",
        "https://example.com/page2",
        "https://example.com/page3",
        "https://example.com/page4",
        "https://example.com/page5",
        "https://example.com/page6",
        "https://example.com/page7",
        "https://example.com/page8",
    }

    // Create scraper with 3 workers and 50ms rate limit
    scraper := NewScraper(3, 50*time.Millisecond)

    // Create cancellable context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Track progress
    progress := &Progress{total: len(urls)}

    fmt.Printf("Scraping %d URLs with %d workers...\n\n", len(urls), 3)
    
    // Start scraping
    results := scraper.Scrape(ctx, urls)

    // Process results as they arrive
    var allResults []Result
    for result := range results {
        allResults = append(allResults, result)
        progress.Update()
        progress.Print()
    }

    // Print summary
    fmt.Println("\n\n=== Results ===")
    for _, r := range allResults {
        if r.Error != nil {
            fmt.Printf("FAIL  %s: %v\n", r.URL, r.Error)
        } else {
            fmt.Printf("OK    %s (%d bytes, %v)\n", r.URL, r.Size, r.Duration.Round(time.Millisecond))
        }
    }

    fmt.Printf("\n=== Summary ===\n")
    fmt.Printf("Total: %d URLs processed\n", len(allResults))
}
```

## Walk Through:
- Worker pool for concurrent fetching
- Rate limiting with ticker
- Context for cancellation
- WaitGroup for synchronization
- Mutex for progress tracking
- Channels for result collection
