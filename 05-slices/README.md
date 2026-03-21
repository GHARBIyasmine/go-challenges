# ✅ Quest 05 — Slices

## What are Slices?

Dynamic, flexible views into an underlying array. Unlike arrays, slices can grow and shrink. They are the backbone of Go collections. A slice has a length and a capacity. `append` adds elements (may allocate a new array underneath), `copy` clones data safely.

```go
s := []int{1, 2, 3}
s = append(s, 4)
sub := s[1:3]  // [2, 3]
```

## The Challenge

Build a simple **in-memory to-do list** CLI simulation (no real input needed — hardcode the operations):

1. Start with an empty `[]string` slice
2. Add 5 tasks using `append`
3. Print all tasks with their index
4. Remove the task at index 2 (without leaving a gap)
5. Print the updated list
6. Create a copy of the list using `copy` and modify the copy — prove the original is unchanged

**Example output:**
```
--- To-Do List ---
0: Buy groceries
1: Write tests
2: Read a book
3: Fix bug #42
4: Call dentist

Removed: "Read a book"
--- Updated List ---
0: Buy groceries
1: Write tests
2: Fix bug #42
3: Call dentist

Original unchanged: true
```

## What you will practice

- Slice literals and `append`
- Removing an element by index: `append(s[:i], s[i+1:]...)`
- `copy` and understanding slice independence
- `len` and `cap`

## Hints

- To remove index `i`: `s = append(s[:i], s[i+1:]...)`
- `copy` returns the number of elements copied
- After copying, modifying one slice should not affect the other
