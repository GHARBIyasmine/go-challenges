# Quest 12 — Enums (iota)

## What are Enums in Go?

Go has no built-in `enum` keyword. The idiomatic approach uses a **typed constant** with `iota` — a special counter that starts at 0 and increments automatically within a `const` block. Giving it a named type enables type safety and custom methods.

```go
type Direction int
const (
    North Direction = iota  // 0
    East                    // 1
    South                   // 2
    West                    // 3
)
```

## The Challenge

1. Create a `Weekday` type with constants `Monday` through `Sunday` using `iota` (start Monday at 1)
2. Create a `TrafficLight` type with `Red`, `Yellow`, `Green` states
3. Add a `String() string` method to both types so they print their names (not numbers)
4. Write a `Next() TrafficLight` method that returns the next state in the cycle: Red → Green → Yellow → Red
5. In main, loop through all 7 weekdays and print them. Then simulate 6 traffic light transitions starting from Red.

## What you will practice

- `iota` with a custom offset (`iota + 1`)
- Typed constants for type safety
- `fmt.Stringer` on enum types
- Methods on non-struct types
- `iota` with expressions (e.g., `1 << iota` for bit flags)

## Hints

- `Weekday` starting at 1: `Monday Weekday = iota + 1`
- `String()` can use a `switch` or a lookup slice: `var names = []string{"Red", "Yellow", "Green"}`
- For the cycle: Red(0) → Green(2) → Yellow(1) → Red(0)
