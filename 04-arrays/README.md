# Quest 04 — Arrays

## What are Arrays?

Fixed-length, typed sequences stored contiguously in memory. The **length is part of the type** — `[3]int` and `[5]int` are different, incompatible types. Arrays are value types in Go — assigning one array to another copies all elements.

```go
var a [5]int
b := [3]string{"go", "is", "cool"}
c := [...]int{1, 2, 3, 4}  // length inferred
```

## The Challenge

Declare a `[7]int` array representing daily step counts for a week (pick any values, min 2000, max 15000). Using only index-based `for` loops (no slices):

1. Print each day with its step count (Day 1: 8500 steps)
2. Calculate and print the **total** steps
3. Calculate and print the **average** steps per day
4. Find and print the **highest** and **lowest** step days

**Example output:**
```
Day 1: 8500 steps
Day 2: 7200 steps
...
Total:   54300 steps
Average: 7757 steps/day
Best day:  Day 3 (12000 steps)
Worst day: Day 6 (2800 steps)
```

## What you will practice

- Array declaration and initialization
- Index-based loop over an array
- `len()` on arrays
- Tracking min/max values

## Hints

- Initialize max to `array[0]` and min to `array[0]` before looping
- Arrays are 0-indexed but your days should display as 1-indexed
