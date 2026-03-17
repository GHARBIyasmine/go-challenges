# Quest 24 — Files

## What is File I/O in Go?

The `os` package handles file creation, opening, reading, and writing at a low level. `bufio` adds buffering for efficient line-by-line reading. `io` and `io/ioutil` (now `os` and `io` in modern Go) provide helpers. Always close files with `defer f.Close()`. `encoding/json` and `encoding/csv` handle structured formats.

```go
f, err := os.Create("out.txt")
defer f.Close()
fmt.Fprintln(f, "hello")
```

## The Challenge

1. **Write**: create a file `tasks.txt` and write 5 to-do items, one per line
2. **Read**: read the file back line-by-line using `bufio.Scanner` and print each line with a line number
3. **Append**: open the file in append mode and add 2 more tasks
4. **CSV**: write a `products.csv` with columns `name,price,quantity` for 4 products using `encoding/csv`. Read it back and print a formatted table.
5. **JSON**: marshal a slice of structs to a `data.json` file with indentation. Read and unmarshal it back, then print the result.

## What you will practice

- `os.Create`, `os.Open`, `os.OpenFile` with flags
- `defer f.Close()`
- `bufio.Scanner` for line-by-line reading
- `encoding/csv` Writer and Reader
- `encoding/json` Marshal/Unmarshal with file I/O

## Hints

- Append mode: `os.OpenFile("tasks.txt", os.O_APPEND|os.O_WRONLY, 0644)`
- `bufio.NewScanner(f)` + `scanner.Scan()` loop
- JSON with indent: `json.MarshalIndent(data, "", "  ")`
- Always check errors from `os.Create` and `f.Close()`
