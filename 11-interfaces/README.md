# Quest 11 — Interfaces

## What are Interfaces?

Named collections of method signatures. A type **implicitly satisfies** an interface by implementing all its methods — no `implements` keyword needed. This is Go's mechanism for polymorphism and decoupled design. The empty interface `any` (or `interface{}`) accepts any type.

```go
type Speaker interface {
    Speak() string
}
// Any type with Speak() satisfies Speaker automatically
```

## The Challenge

1. Define a `Shape` interface with two methods: `Area() float64` and `Perimeter() float64`
2. Implement `Shape` for three types: `Circle` (radius), `Rectangle` (width, height), `Triangle` (a, b, c sides — use Heron's formula for area)
3. Write a `printShapeInfo(s Shape)` function that prints the type name, area, and perimeter
4. Store all three shapes in a `[]Shape` slice and range over it calling `printShapeInfo`
5. Use a **type switch** to print shape-specific details (e.g., radius for Circle)

## What you will practice

- Interface definition and implicit satisfaction
- Polymorphism through interfaces
- `[]Shape` — slice of interface type
- Type assertion: `s.(Circle)`
- Type switch: `switch v := s.(type)`

## Hints

- Heron's formula: `s = (a+b+c)/2`, `area = sqrt(s*(s-a)*(s-b)*(s-c))`
- Import `"math"` for `math.Pi`, `math.Sqrt`
- Type switch: `switch v := shape.(type) { case Circle: ... }`
