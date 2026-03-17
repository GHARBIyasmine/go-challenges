# Quest 21 — Mutexes

## What are Mutexes?

A Mutex (mutual exclusion lock) prevents race conditions when multiple goroutines access shared data. `sync.Mutex` has `Lock()` and `Unlock()`. `sync.RWMutex` allows multiple concurrent readers but only one writer. Always defer `Unlock()` to prevent deadlocks.

```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
sharedData++
```

## The Challenge

1. First, write a **deliberately racy counter**: launch 100 goroutines each incrementing a shared `int` 1000 times without any synchronization. Run with `go run -race main.go` to see the race condition.
2. Fix it using `sync.Mutex` — confirm the final count is always 100,000.
3. Build a **thread-safe cache** struct using `sync.RWMutex` with `Get(key string) (string, bool)` and `Set(key, value string)` methods.
4. Test the cache by launching 10 reader goroutines and 3 writer goroutines simultaneously.

## What you will practice

- Identifying and creating a race condition
- `sync.Mutex` Lock/Unlock
- `defer mu.Unlock()` pattern
- `sync.RWMutex` for read-heavy workloads
- Go's race detector (`-race` flag)

## Hints

- Race detector: `go run -race main.go`
- For the cache: `mu.RLock() / mu.RUnlock()` in `Get`, `mu.Lock() / mu.Unlock()` in `Set`
- Use `sync.WaitGroup` to wait for all goroutines before printing the final count
