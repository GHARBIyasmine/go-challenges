# Quest 28 — HTTP Client

## What is an HTTP Client in Go?

`http.Get`, `http.Post`, and `http.NewRequest` + `http.Client.Do` are Go's tools for making HTTP requests. Always close `resp.Body` with `defer resp.Body.Close()`. Custom clients allow setting timeouts, headers, and redirect policies. `encoding/json` decodes JSON responses.

```go
resp, err := http.Get("https://api.example.com/data")
defer resp.Body.Close()
json.NewDecoder(resp.Body).Decode(&result)
```

## The Challenge

Use the free public API `https://jsonplaceholder.typicode.com`:

1. `GET /users` — fetch and print all user names and emails in a table
2. `GET /posts?userId=1` — fetch posts for user 1, print titles
3. `POST /posts` — create a new post (send JSON body), print the response
4. Create a custom `http.Client` with a 5-second timeout
5. Write a `fetchJSON[T any](url string, target *T) error` generic function that handles GET + decode in one call

## What you will practice

- `http.Get`, `http.NewRequest`, `http.Client`
- `resp.Body` reading and closing
- `json.NewDecoder` for response bodies
- Custom `http.Client` with timeout
- Query parameters: `url.Values` and `req.URL.RawQuery`
- Generic HTTP helper function

## Hints

- Custom client: `client := &http.Client{Timeout: 5 * time.Second}`
- Check status: `if resp.StatusCode != http.StatusOK { return error }`
- POST with JSON: `http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))`
- Set header: `req.Header.Set("Content-Type", "application/json")`
