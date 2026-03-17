# Quest 16 — Channels

## What are Channels?

Typed conduits for communication between goroutines. Channels synchronize goroutines safely without shared memory — "do not communicate by sharing memory; share memory by communicating." A channel can be **unbuffered** (blocks until both sides are ready) or **buffered** (has a queue of capacity N).

```go
ch := make(chan int)       // unbuffered
ch := make(chan int, 10)   // buffered, capacity 10
ch <- 42                   // send
v := <-ch                  // receive
close(ch)                  // signal no more values
```

## The Challenge

Build a **pipeline** with three stages connected by channels:

1. **Generator**: a goroutine that sends integers 1–10 into a channel, then closes it
2. **Squarer**: a goroutine that reads from the generator channel, squares each number, sends to a second channel
3. **Printer**: the main goroutine reads from the squarer channel and prints each result

Then build a second example using a **buffered channel** as a semaphore to limit concurrency to 3 simultaneous "workers" out of 10.

## What you will practice

- Unbuffered and buffered channels
- `close(ch)` and `range ch` (stops when closed)
- Pipeline pattern
- Using a buffered channel as a semaphore

## Hints

- Generator: `func generate(ch chan<- int)` — send-only channel parameter
- Range over a channel: `for v := range ch { ... }` — stops when channel is closed
- Buffered semaphore: `sem := make(chan struct{}, 3)`, send before work, receive after
