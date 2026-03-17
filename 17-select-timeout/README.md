# Quest 17 — Select & Timeout

## What is Select?

`select` lets a goroutine wait on **multiple channel operations** simultaneously — it picks whichever case is ready first. If multiple are ready, it picks randomly. A `default` case makes it non-blocking. Combined with `time.After`, it's the standard way to implement timeouts in Go.

```go
select {
case msg := <-ch1:   // ...
case msg := <-ch2:   // ...
case <-time.After(1 * time.Second):  // timeout
}
```

## The Challenge

1. Create two channels: `fast` (responds after 100ms) and `slow` (responds after 500ms). Use `select` to print whichever arrives first.
2. Build a `fetchWithTimeout(url string, timeout time.Duration) (string, error)` function (simulate the fetch with a goroutine + sleep). Return a timeout error if it takes too long.
3. Create a **fan-in** function: merge two channels into one output channel using `select` in a loop.

## What you will practice

- `select` with multiple channel cases
- `time.After` for timeouts
- Non-blocking select with `default`
- Fan-in / merge pattern

## Hints

- `time.After(d)` returns a `<-chan time.Time` that fires after duration `d`
- For fan-in: use a goroutine with an infinite loop and `select` over both input channels, sending to the output channel
- For `fetchWithTimeout`: pass result through a `chan string` and race it against `time.After`
