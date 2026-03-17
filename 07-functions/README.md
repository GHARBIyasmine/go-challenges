# Quest 07 — Functions

## What are Functions?

First-class citizens in Go. Functions can return **multiple values** (commonly a result + error), accept **variadic arguments** (`...T`), be assigned to variables (closures), and use **defer** to schedule cleanup code that runs when the function returns — even if it panics.

```go
func divide(a, b float64) (float64, error) { ... }
func sum(nums ...int) int { ... }
defer fmt.Println("cleanup")
```

## The Challenge

Build a **mini calculator** with the following functions:

1. `add(a, b float64) float64`
2. `subtract(a, b float64) float64`
3. `multiply(a, b float64) float64`
4. `divide(a, b float64) (float64, error)` — return an error if b is 0
5. `sum(nums ...int) int` — variadic, sums any number of integers
6. In `main`, call each function and use `defer` to print `"--- calculation complete ---"` at the end of each call (wrap each in a helper that defers)

## What you will practice

- Multiple return values
- Named error returns
- Variadic functions (`...int`)
- `defer` execution order
- Function signatures and calling conventions

## Hints

- `errors.New("division by zero")` creates a simple error
- Variadic: call with `sum(1, 2, 3, 4)` or spread a slice `sum(nums...)`
- `defer` runs LIFO — if you defer twice, second deferred runs first
