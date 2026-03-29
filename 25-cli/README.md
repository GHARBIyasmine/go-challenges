# ✅ Quest 25 — CLI

## What is CLI in Go?

Go's `flag` package provides built-in command-line argument parsing. For more complex CLIs with subcommands, the popular `cobra` library is standard in the ecosystem. `os.Args` gives raw arguments. Flags can be strings, ints, bools, and support default values and usage descriptions.

```go
name := flag.String("name", "World", "who to greet")
flag.Parse()
fmt.Printf("Hello, %s!\n", *name)
// go run main.go -name=Alice
```

## The Challenge

Build a **task manager CLI** using only the standard `flag` package:

1. Support these commands via flags:
   - `-add "task name"` — adds a task to `tasks.json`
   - `-list` — lists all tasks with IDs
   - `-done <id>` — marks a task as complete
   - `-delete <id>` — removes a task
2. Persist tasks to a `tasks.json` file in the current directory
3. On `-list`, show: ID, status ([ ] or [x]), and task name
4. If no flags are given, print usage instructions

**Example:**
```
$ go run main.go -add "Buy milk"
Added task 1: Buy milk

$ go run main.go -list
1. [ ] Buy milk

$ go run main.go -done 1
Task 1 marked as done

$ go run main.go -list
1. [x] Buy milk
```

## What you will practice

- `flag.String`, `flag.Bool`, `flag.Int`
- `flag.Parse()` and reading flag values
- Combining CLI with file I/O and JSON
- Program structure for command dispatch

## Hints

- Check if a flag was set: `flag.Visit(func(f *flag.Flag) { ... })`
- Or check default values: if `*add != ""` then add was provided
- `flag.Usage` can be overridden with a custom function
