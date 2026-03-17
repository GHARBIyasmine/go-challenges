# Quest 20 — Rate Limiting

## What is Rate Limiting?

Controlling how frequently an operation is performed — e.g., max 5 API requests per second. In Go, this is idiomatically implemented using a ticker as a token source. The `golang.org/x/time/rate` package provides a production-grade token bucket limiter.

```go
limiter := time.NewTicker(200 * time.Millisecond)
// Wait for a token before each request
<-limiter.C
```

## The Challenge

1. Implement **basic rate limiting**: process 8 "requests" but allow only 1 every 150ms using a ticker. Print the timestamp of each request.
2. Implement **bursty rate limiting**: allow a burst of 3 requests immediately, then 1 per 200ms for the remaining requests (use a buffered channel pre-filled with tokens).
3. Create a `RateLimiter` struct that wraps the logic, with a `Wait()` method that blocks until a token is available and a `TryAcquire() bool` method that is non-blocking.

## What you will practice

- Ticker-based rate limiting
- Buffered channel as a token bucket
- Non-blocking channel receive with `select` + `default`
- Measuring elapsed time per request

## Hints

- Basic: `for req := range requests { <-ticker.C; handle(req) }`
- Burst: pre-fill a buffered channel `burstyLimiter := make(chan time.Time, 3)` with 3 values upfront
- `TryAcquire`: use `select { case <-tokens: return true; default: return false }`
