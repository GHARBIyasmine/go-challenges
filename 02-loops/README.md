# Quest 02 — Loops

## What are Loops?

Go has exactly **one** loop keyword: `for`. It covers all looping patterns — there is no `while` or `do-while`. The three forms are: classic 3-part (`for i := 0; i < n; i++`), condition-only (acts like `while`), and `range`-based (iterates over slices, maps, strings, channels).

```go
for i := 0; i < 5; i++ { }          // classic
for x > 0 { }                        // while-style
for i, v := range slice { }          // range
```

## The Challenge

1. Print the **multiplication table** from 1 to 10 using nested `for` loops
2. Using a `range`-based loop, sum a hard-coded slice of 6 numbers and print the total
3. Using a condition-only loop, find the first power of 2 greater than 1000

**Example output:**
```
1  2  3  4  5  6  7  8  9  10
2  4  6  8  10 ...
...

Sum: 97
First power of 2 > 1000: 1024
```

## What you will practice

- 3-part `for` loop
- Nested loops
- `range` over a slice
- Condition-only loop (while-style)
- `break` when a condition is met

## Hints

- Use `fmt.Printf("%4d", v)` for aligned table columns
- `range` gives you both index and value: `for i, v := range nums`
- Start with `n := 1` and double it each iteration: `n *= 2`
