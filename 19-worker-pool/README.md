# Quest 19 — Worker Pool

## What is a Worker Pool?

A concurrency pattern where a fixed number of goroutines (workers) pull tasks from a shared jobs channel and send results to a results channel. It limits concurrency to a set number of goroutines regardless of how many tasks there are — efficient for CPU or I/O bound work.

```
jobs ──▶ [worker 1]
      ──▶ [worker 2]  ──▶ results
      ──▶ [worker 3]
```

## The Challenge

Build a **URL batch processor** simulation:

1. Create a `jobs` channel and a `results` channel
2. Launch 3 worker goroutines. Each worker reads a job (an int representing a "URL ID"), simulates work with `time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)`, and sends a result string like `"processed job 7 by worker 2"`
3. Send 10 jobs into the jobs channel, then close it
4. Collect and print all 10 results
5. Print total elapsed time

## What you will practice

- Worker pool pattern
- Closing a jobs channel to signal workers to stop
- Collecting results from multiple goroutines safely
- `sync.WaitGroup` with goroutines that close the results channel when done

## Hints

- Launch workers before sending jobs to avoid deadlock
- Close `results` after all workers are done: use a WaitGroup in a separate goroutine that calls `wg.Wait()` then `close(results)`
- Workers: `for j := range jobs { ... results <- result }`
