# Quest 27 — HTTP Server

## What is an HTTP Server in Go?

Go's `net/http` package makes building HTTP servers straightforward. `http.HandleFunc` registers route handlers. `http.ListenAndServe` starts the server. Handlers receive an `http.ResponseWriter` (to write responses) and an `*http.Request` (to read the request). Go's server is production-grade — used by major companies as-is.

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
})
http.ListenAndServe(":8080", nil)
```

## The Challenge

Build a **mini REST API** for the task manager from Quest 25 (in-memory is fine):

1. `GET /tasks` — returns all tasks as JSON
2. `POST /tasks` — creates a new task (JSON body: `{"name":"..."}`)
3. `GET /tasks/{id}` — returns a single task
4. `DELETE /tasks/{id}` — deletes a task
5. Add a middleware that logs every request: method, path, and response time
6. Return proper HTTP status codes (200, 201, 404, 400)

**Test with curl:**
```bash
curl -X POST http://localhost:8080/tasks -d '{"name":"Buy milk"}'
curl http://localhost:8080/tasks
```

## What you will practice

- `http.HandleFunc` and `http.ServeMux`
- `json.NewDecoder(r.Body)` and `json.NewEncoder(w)`
- `w.WriteHeader(http.StatusCreated)`
- `r.Method` for routing by HTTP method
- Middleware pattern (wrapping `http.Handler`)
- Parsing URL path segments

## Hints

- Parse ID from path: `strings.TrimPrefix(r.URL.Path, "/tasks/")`
- Middleware: `func logging(next http.Handler) http.Handler`
- Set content type: `w.Header().Set("Content-Type", "application/json")`
- Use a `sync.Mutex` to protect your in-memory task map
