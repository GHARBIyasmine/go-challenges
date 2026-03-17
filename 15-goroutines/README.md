# Quest 15 — Goroutines

## What are Goroutines?

Lightweight, concurrently executing functions managed by the Go runtime — not OS threads. You launch one with the `go` keyword. Goroutines are cheap (a few KB of stack) and can number in the thousands. A `sync.WaitGroup` is used to wait for a set of goroutines to finish.

```go
go func() {
    fmt.Println("running concurrently")
}()
```

## The Challenge

Simulate a **news aggregator** that fetches from multiple sources concurrently:

1. Create a `fetchFeed(source string, delay time.Duration)` function that prints `"Fetched: <source>"` after sleeping for `delay`
2. Launch 5 goroutines, each simulating a different feed with different delays (50ms–300ms)
3. Use a `sync.WaitGroup` to wait for all goroutines to complete
4. Print total elapsed time at the end
5. As a bonus: run the same 5 fetches **sequentially** and compare total time

**Expected observation:** Concurrent version takes ~300ms. Sequential takes ~sum of all delays.

## What you will practice

- `go` keyword to launch goroutines
- `sync.WaitGroup` (`Add`, `Done`, `Wait`)
- `time.Sleep` and `time.Since`
- Seeing concurrency speed-up in practice

## Hints

- Call `wg.Add(1)` before launching the goroutine
- Call `defer wg.Done()` as the first line inside the goroutine
- Pass `source` and `delay` as arguments to avoid closure capture bugs: `go func(s string, d time.Duration) {...}(source, delay)`
