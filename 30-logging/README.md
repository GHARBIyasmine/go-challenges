# Quest 30 — Logging

## What is Logging in Go?

The standard `log` package provides basic logging with timestamps. Go 1.21 introduced `log/slog` — a structured logging package with levels (Debug, Info, Warn, Error) and key-value attributes. Structured logs (JSON or text) are machine-readable and searchable in log aggregation tools.

```go
slog.Info("server started", "port", 8080)
slog.Error("request failed", "err", err, "path", r.URL.Path)
```

## The Challenge

1. Use the basic `log` package: set a custom prefix and flags (`log.Ldate|log.Ltime|log.Lshortfile`). Log a series of events. Use `log.Fatal` to exit on a simulated critical error.
2. Create a structured logger with `slog` using JSON output. Log Info, Warn, and Error events with key-value attributes.
3. Create a **text handler** for human-readable output and a **JSON handler** for machine-readable output. Write a function that logs the same event to both.
4. Create a child logger with extra context: `logger.With("requestID", "abc123", "userID", 42)` and use it throughout a simulated request handler.

## What you will practice

- `log.SetPrefix`, `log.SetFlags`
- `log.Fatal`, `log.Panic` vs `log.Print`
- `slog.New`, `slog.NewJSONHandler`, `slog.NewTextHandler`
- `slog.Info`, `slog.Warn`, `slog.Error` with attributes
- `logger.With(...)` for contextual loggers
- Writing to `os.Stderr` vs `os.Stdout`

## Hints

- JSON handler: `slog.New(slog.NewJSONHandler(os.Stdout, nil))`
- Set as default: `slog.SetDefault(logger)`
- Log levels: `slog.LevelDebug`, `slog.LevelInfo`, `slog.LevelWarn`, `slog.LevelError`
