# Quest 13 — Generics

## What are Generics?

Introduced in Go 1.18, **type parameters** let you write functions and types that work across multiple types while keeping full type safety. A **constraint** (an interface) limits what operations are valid on the type parameter.

```go
func Map[T, U any](s []T, f func(T) U) []U { ... }
// Call: Map([]int{1,2,3}, func(x int) string { return strconv.Itoa(x) })
```

## The Challenge

Write three generic utility functions for slices:

1. `Map[T, U any](slice []T, fn func(T) U) []U` — transforms each element
2. `Filter[T any](slice []T, fn func(T) bool) []T` — keeps elements where fn is true
3. `Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U` — folds into a single value

Then use them to:
- Double a `[]int`
- Filter only even numbers from a `[]int`
- Sum a `[]int` using Reduce
- Join a `[]string` with commas using Reduce
- Chain Map + Filter: convert `[]string` of numbers to `[]int`, then filter those > 3

## What you will practice

- Type parameter syntax `[T any]`
- Constraints (`any`, `comparable`, custom interfaces)
- Calling generic functions with explicit and inferred type params
- `golang.org/x/exp/constraints` or defining your own `Number` constraint

## Hints

- A `Number` constraint: `type Number interface { int | float64 }`
- Inferred types: `Map([]int{1,2}, double)` — Go infers `T=int, U=int`
- `strconv.Atoi` converts string to int (returns error — handle it)
