# ✅ Quest 06 — Maps

## What are Maps?

Built-in key-value hash tables. Keys must be a comparable type (strings, ints, etc.). Maps are **unordered** — iteration order is randomized intentionally. The zero value of a map is `nil` — always initialize with `make` or a map literal before use.

```go
m := make(map[string]int)
m["score"] = 42
val, ok := m["score"]  // ok is false if key missing
```

## The Challenge

Build a **word frequency counter**:

1. Take this hard-coded sentence: `"the quick brown fox jumps over the lazy dog the fox"`
2. Split it into words using `strings.Fields`
3. Count each word's occurrences in a `map[string]int`
4. Print all words and counts **sorted alphabetically** (use `sort.Strings` on the keys)
5. Find and print the most frequent word(s)
6. Delete the word `"the"` from the map and reprint

## What you will practice

- `make(map[...]...)` initialization
- Map literal syntax
- The "comma ok" idiom: `val, ok := m[key]`
- Iterating over maps
- Extracting and sorting keys manually

## Hints

- `strings.Fields(s)` splits by whitespace
- To get sorted keys: collect them into a `[]string`, then `sort.Strings(keys)`
- Check existence: `if count, ok := freq[word]; ok { ... }`
- `delete(m, key)` removes a key
