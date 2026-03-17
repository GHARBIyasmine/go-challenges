# Quest 01 — Variables

## What are Variables?

Named storage that holds a typed value. Go is **statically typed** — every variable has a fixed type known at compile time. You can declare variables with `var` (explicit) or `:=` (short declaration, type inferred). Constants use `const` and cannot be changed after declaration.

```go
var name string = "Go"
age := 5
const Pi = 3.14159
```

## The Challenge

Build a **unit converter** that:

1. Declares a temperature value in Celsius using `var` with an explicit type
2. Converts it to Fahrenheit `(C * 9/5) + 32` and Kelvin `C + 273.15`
3. Declares a `const` for absolute zero in Celsius (`-273.15`)
4. Prints all three results with their units
5. Uses a `var` block to declare at least 3 variables at once

**Example output:**
```
Temperature: 100°C
Fahrenheit:  212.00°F
Kelvin:      373.15K
Absolute zero: -273.15°C
```

## What you will practice

- `var` declarations (single and block)
- `:=` short declarations
- `const` keyword
- Basic arithmetic with typed variables
- Explicit type conversion (`float64(x)`)

## Hints

- Celsius to Fahrenheit: `(c * 9 / 5) + 32`
- Celsius to Kelvin: `c + 273.15`
- Use `fmt.Printf("%.2f", value)` for 2 decimal places

