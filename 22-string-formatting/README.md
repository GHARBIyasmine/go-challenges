# Quest 22 — String Formatting

## What is String Formatting?

Go's `fmt` package uses **verbs** to format values: `%v` (default), `%+v` (with field names), `%#v` (Go syntax), `%T` (type), `%d` (int), `%f` (float), `%s` (string), `%q` (quoted), `%x` (hex), `%b` (binary). Width and precision: `%8.2f`, `%-10s` (left-align). Implementing `fmt.Stringer` (`String() string`) customizes how a type prints.

## The Challenge

1. Create a `Product` struct with `Name string`, `Price float64`, `Stock int`
2. Print it using `%v`, `%+v`, and `%#v` — observe the differences
3. Implement `String() string` on `Product` with a clean custom format
4. Print a formatted **table** of 5 products with right-aligned columns: Name (left, 15 chars), Price (right, 8 chars, 2 decimals), Stock (right, 6 chars)
5. Format the price in hex, binary, and with thousand separators using `fmt.Sprintf` and `strconv`

**Example table:**
```
Name               Price  Stock
---------------  -------  -----
Apple MacBook   1299.99     42
USB-C Cable        9.99    500
```

## What you will practice

- All major fmt verbs
- Width and precision formatting: `%-15s`, `%8.2f`
- `fmt.Sprintf` vs `fmt.Printf` vs `fmt.Fprintf`
- `fmt.Stringer` interface
- `strings.Repeat` for table borders

## Hints

- Left-align: `%-15s`, right-align: `%15s`
- `%08.2f` — zero-padded float
- `fmt.Fprintf(os.Stdout, ...)` writes to a writer — same verbs as Printf
