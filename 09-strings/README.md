# ✅ Quest 09 — Strings

## What are Strings?

Immutable sequences of bytes in Go, encoded as **UTF-8**. A `string` and a `[]byte` are easily convertible. The `strings` package covers searching, splitting, replacing. The `strconv` package handles conversions to/from other types. `rune` is an alias for `int32` and represents a Unicode code point — important when iterating over non-ASCII text.

```go
s := "héllo"
len(s)           // byte count (may differ from char count)
[]rune(s)        // rune slice — safe for Unicode
strings.ToUpper(s)
```

## The Challenge

Build a **string inspector** program. Given the input string `"racecar 🚀 Go"`:

1. Print the number of **bytes** (`len`)
2. Print the number of **runes** (`utf8.RuneCountInString`)
3. Print the number of **words** (`strings.Fields`)
4. **Reverse** the string safely (handle multi-byte runes correctly)
5. Check if the word `"racecar"` is a **palindrome**
6. Replace `"Go"` with `"Golang"` and print

## What you will practice

- `len` vs `utf8.RuneCountInString`
- Converting string ↔ `[]rune` for safe iteration
- `strings.Fields`, `strings.Replace`, `strings.ToLower`
- Building a palindrome checker

## Hints

- Reverse via rune slice: `r := []rune(s); for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 { r[i], r[j] = r[j], r[i] }`
- Palindrome: reverse the word and compare to original
- Import `"unicode/utf8"` for `RuneCountInString`
