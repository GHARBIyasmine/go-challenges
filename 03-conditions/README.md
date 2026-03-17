# Quest 03 — Conditions

## What are Conditions?

Control flow based on boolean expressions. Go has `if/else if/else` and `switch`. Go's `switch` is more powerful than most languages — no automatic fallthrough, cases can be expressions or ranges, and you can switch on nothing (acts like if/else chains).

```go
if x > 10 { } else if x > 5 { } else { }

switch {
case x > 10: ...
case x > 5:  ...
}
```

## The Challenge

1. Write a **grade classifier**: given a score (0–100), use `if/else if/else` to print A (90+), B (80+), C (70+), D (60+), or F
2. **Rewrite the same logic** using a `switch` statement with expression cases
3. Add a bonus check: if the score is exactly 100, print "Perfect score!" using a switch with `fallthrough`

**Example output:**
```
Score: 85
if/else result:  B
switch result:   B

Score: 100
switch result:   A
Perfect score!
```

## What you will practice

- `if / else if / else`
- `switch` with expression cases
- `fallthrough` keyword
- Initializer statement in `if`: `if score := getScore(); score > 90 {}`

## Hints

- Switch cases can use `>=`: `case score >= 90:`
- `fallthrough` forces execution of the next case regardless of condition
