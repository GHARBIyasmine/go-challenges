# Quest 26 — Environment Variables

## What are Environment Variables?

Key-value pairs provided by the operating system or shell to a running process. In Go, `os.Getenv` reads them, `os.Setenv` sets them for the current process, and `os.LookupEnv` distinguishes between "not set" and "set to empty string". The `godotenv` library (or manual parsing) loads `.env` files.

```go
port := os.Getenv("PORT")           // "" if not set
val, ok := os.LookupEnv("API_KEY")  // ok=false if not set
```

## The Challenge

Build a **config loader** for a fictional web service:

1. Read these env vars: `APP_PORT` (default `8080`), `APP_ENV` (default `development`), `DB_URL` (required — exit with error if missing), `MAX_CONNECTIONS` (int, default `10`)
2. Use `os.LookupEnv` to distinguish missing from empty
3. Create a `Config` struct and a `LoadConfig() (Config, error)` function that reads all vars and returns a populated struct or error
4. Print the loaded config in a formatted way
5. Manually parse a `.env` file (format: `KEY=VALUE` per line, skip `#` comments) without using any external library

**Run it like:**
```bash
APP_PORT=9000 APP_ENV=production DB_URL=postgres://localhost/mydb go run main.go
```

## What you will practice

- `os.Getenv`, `os.LookupEnv`, `os.Setenv`
- `strconv.Atoi` for int env vars
- `os.Exit(1)` for fatal errors
- File parsing for `.env` format
- Struct-based config pattern

## Hints

- `os.LookupEnv` returns `(value string, found bool)`
- For the `.env` parser: `strings.SplitN(line, "=", 2)` splits on first `=` only
- Skip lines starting with `#` or empty lines
