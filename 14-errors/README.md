# ✅ Quest 14 — Errors

## What are Errors?

In Go, **errors are values** — functions return them explicitly rather than throwing exceptions. The built-in `error` interface has one method: `Error() string`. The `errors` and `fmt` packages let you create, wrap, and inspect errors. Wrapping with `%w` preserves the error chain for `errors.Is` and `errors.As`.

```go
var ErrNotFound = errors.New("not found")
fmt.Errorf("load failed: %w", ErrNotFound)
errors.Is(err, ErrNotFound)  // true even if wrapped
```

## The Challenge

Simulate a **user lookup service**:

1. Define a custom `NotFoundError` struct with a `Username string` field and an `Error() string` method
2. Define a custom `ValidationError` struct with a `Field string` and `Message string`
3. Write `FindUser(username string) (string, error)` that:
   - Returns `ValidationError` if username is empty
   - Returns `NotFoundError` if username is not `"alice"` or `"bob"`
   - Returns the user's display name if found
4. In main, call `FindUser` with 4 inputs: `""`, `"charlie"`, `"alice"`, `"bob"`
5. Use `errors.As` to detect and print specific error details for each failure

## What you will practice

- Custom error types with `Error() string`
- `fmt.Errorf` with `%w` wrapping
- `errors.Is` (sentinel errors) vs `errors.As` (typed errors)
- Error handling patterns in Go

## Hints

- `errors.As(err, &target)` fills `target` if the error chain contains that type
- Wrap your custom errors: `return fmt.Errorf("findUser: %w", &NotFoundError{Username: username})`
- The zero value check: `if username == "" { return "", &ValidationError{...} }`
