# Quest 32 — Signals & Exit

## What are Signals and Exit in Go?

OS signals (like `SIGINT` from Ctrl+C, `SIGTERM` from kill) notify a process of external events. `os/signal` lets Go programs intercept these signals for **graceful shutdown** — finishing in-flight work before exiting. `os.Exit(code)` terminates immediately. `defer` does NOT run after `os.Exit`.

```go
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit  // block until signal received
// do cleanup...
```

## The Challenge

Build a **graceful shutdown server** simulation:

1. Start a goroutine simulating ongoing background work (prints "working..." every 500ms)
2. Set up signal handling to catch `SIGINT` (Ctrl+C) and `SIGTERM`
3. When a signal is received:
   - Print "Shutdown signal received"
   - Signal the background worker to stop (use a done channel)
   - Wait up to 3 seconds for it to finish (use a context with timeout)
   - Print "Shutdown complete" and exit with code 0
4. If you press Ctrl+C twice, force-exit immediately
5. Demonstrate `os.Exit` vs returning from `main` (show that defers don't run with `os.Exit`)

**Run it and press Ctrl+C to trigger graceful shutdown.**

## What you will practice

- `signal.Notify` with a buffered channel
- `syscall.SIGINT`, `syscall.SIGTERM`
- Done channel pattern for stopping goroutines
- `context.WithTimeout` for shutdown deadline
- `os.Exit` behavior vs normal return
- The difference between `log.Fatal` (calls `os.Exit`) and `panic`

## Hints

- Signal channel must be buffered: `make(chan os.Signal, 1)`
- `signal.Stop(quit)` unregisters signal forwarding
- Done channel: `done := make(chan struct{})`, close it to broadcast stop to all listeners
- Double Ctrl+C: keep a counter or use a second `signal.Notify` channel
