# Quest 18 — Tickers

## What are Tickers?

`time.Ticker` sends the current time on its channel at regular intervals — useful for recurring tasks. `time.Timer` fires once after a delay. Both should be stopped with `.Stop()` to release resources. Tickers are the Go equivalent of setInterval.

```go
ticker := time.NewTicker(500 * time.Millisecond)
defer ticker.Stop()
for t := range ticker.C {
    fmt.Println("tick at", t)
}
```

## The Challenge

1. Create a ticker that prints `"ping"` every 200ms. Stop it after 1 second (use `time.After` or a `time.Timer`). Count and print how many ticks occurred.
2. Build a simple **health-check monitor**: every 300ms it "checks" a service (simulate with a random bool), prints OK or FAIL, and stops after 5 checks.
3. Combine a fast ticker (100ms) and a slow ticker (500ms) using `select` — label each tick as "fast" or "slow". Run for 2 seconds.

## What you will practice

- `time.NewTicker` and `.Stop()`
- `time.NewTimer` and `.Stop()`
- Combining tickers with `select`
- Counting events from channel receives

## Hints

- `ticker.Stop()` does NOT close the channel — use a done channel or `time.After` for the stop signal
- `math/rand` for the random bool: `rand.Intn(2) == 0`
- For counting: declare `count := 0` and increment inside the tick case
