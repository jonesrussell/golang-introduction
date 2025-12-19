# Common Concurrency Patterns

**Duration:** 8-10 minutes

## Topics to cover:
- Worker pool
- Fan-out/fan-in
- Pipeline
- Semaphore

## Worker Pool Pattern

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second)  // Simulate work
        results <- job * 2
    }
}

func main() {
    numJobs := 10
    numWorkers := 3

    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Start workers
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for r := 1; r <= numJobs; r++ {
        fmt.Println("Result:", <-results)
    }
}
```

## Pipeline Pattern

```go
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func double(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * 2
        }
        close(out)
    }()
    return out
}

func main() {
    // Pipeline: generate -> square -> double -> print
    nums := generate(1, 2, 3, 4, 5)
    squared := square(nums)
    doubled := double(squared)

    for result := range doubled {
        fmt.Println(result)  // 2, 8, 18, 32, 50
    }
}
```

## Semaphore Pattern

```go
func main() {
    maxConcurrent := 3
    sem := make(chan struct{}, maxConcurrent)

    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            sem <- struct{}{}        // Acquire
            defer func() { <-sem }() // Release

            fmt.Printf("Task %d running\n", id)
            time.Sleep(time.Second)
        }(i)
    }

    wg.Wait()
}
```

## Key teaching points:
- Worker pool distributes work across goroutines
- Pipelines chain processing stages
- Semaphores limit concurrent operations
- Choose pattern based on problem structure
