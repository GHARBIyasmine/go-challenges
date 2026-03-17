# Quest 29 — Context

## What is Context?

`context.Context` carries deadlines, cancellation signals, and request-scoped values across API boundaries and goroutines. It's the standard way to cancel long-running operations. `context.WithCancel` lets you cancel manually, `context.WithTimeout` adds a deadline, `context.WithValue` attaches key-value data.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
doSomething(ctx)
```

## The Challenge

1. Write `fetchWithContext(ctx context.Context, delay time.Duration) (string, error)` — simulates a slow operation. Cancel it if the context expires.
2. Call it three ways: with a 2s timeout (fast enough), with a 500ms timeout (too slow — times out), and with manual `cancel()` called after 100ms.
3. Build a **request pipeline**: a chain of 3 functions that each do work and check `ctx.Err()` before proceeding. Cancel mid-pipeline and observe where it stops.
4. Use `context.WithValue` to pass a request ID through the pipeline and log it at each stage.

## What you will practice

- `context.Background()` and `context.TODO()`
- `context.WithCancel`, `context.WithTimeout`, `context.WithDeadline`
- `ctx.Done()` channel and `ctx.Err()`
- `context.WithValue` and `ctx.Value(key)`
- Passing context as the first argument (Go convention)

## Hints

- In a goroutine: `select { case <-ctx.Done(): return ctx.Err(); case result := <-work: ... }`
- Always `defer cancel()` immediately after `WithTimeout` or `WithCancel`
- Context values: use unexported key types to avoid collisions: `type ctxKey string`
