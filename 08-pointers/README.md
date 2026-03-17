# Quest 08 — Pointers

## What are Pointers?

Variables that store a **memory address** rather than a value. `&x` gives the address of `x`. `*p` dereferences a pointer to get/set the value at that address. Go has no pointer arithmetic — this makes it safer than C. Pointer receivers on methods allow mutating struct state.

```go
x := 42
p := &x   // p is *int
*p = 100  // x is now 100
```

## The Challenge

1. Write a `swap(a, b *int)` function that swaps two integers in place using pointers. Verify by printing before and after.
2. Create a `Counter` struct with a single `count int` field. Add:
   - `Increment()` method with a **pointer receiver** that adds 1
   - `Reset()` method with a pointer receiver that sets count to 0
   - `Value()` method with a **value receiver** that returns the count
3. In main, create a counter, increment it 5 times, print the value, reset it, print again.

## What you will practice

- `&` (address-of) and `*` (dereference) operators
- Passing pointers to functions
- Pointer receivers vs value receivers
- When mutation requires a pointer

## Hints

- `swap` signature: `func swap(a, b *int)`
- Inside swap: `*a, *b = *b, *a`
- Pointer receiver: `func (c *Counter) Increment()`
- Value receiver: `func (c Counter) Value() int`
