# Quest 23 — Regex

## What is Regex in Go?

The `regexp` package provides RE2-syntax regular expressions. Always **compile once** with `regexp.MustCompile` (panics on bad pattern) or `regexp.Compile` (returns error). Compiled regexes are safe for concurrent use. Key methods: `MatchString`, `FindString`, `FindAllString`, `ReplaceAllString`, `FindStringSubmatch` (capture groups).

```go
re := regexp.MustCompile(`\d+`)
re.FindAllString("a1b2c3", -1)  // ["1", "2", "3"]
```

## The Challenge

Build an **input validator and extractor**:

1. Validate an email address format using regex
2. Extract all URLs from a block of text
3. Extract all dates in `YYYY-MM-DD` format using **capture groups** for year, month, day separately
4. Redact all credit card numbers (format: `1234-5678-9012-3456`) by replacing digits with `*` while keeping the dashes
5. Check if a password meets rules: min 8 chars, at least one uppercase, one digit, one special char

## What you will practice

- `regexp.MustCompile`
- `MatchString`, `FindAllString`, `ReplaceAllString`
- Named capture groups: `(?P<year>\d{4})`
- `FindStringSubmatch` for extracting groups
- Compiling patterns once and reusing them

## Hints

- Email pattern: `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
- URL pattern: `https?://[^\s]+`
- Date with groups: `(\d{4})-(\d{2})-(\d{2})`
- Credit card redact: `re.ReplaceAllStringFunc` to process each match
