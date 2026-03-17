# Quest 31 — Processes

## What are Processes in Go?

`os/exec` lets Go programs launch and interact with external processes. `exec.Command` creates a command, `.Run()` runs it synchronously, `.Start()` + `.Wait()` run it asynchronously. You can capture stdout/stderr, pipe input, and check exit codes.

```go
cmd := exec.Command("ls", "-la")
out, err := cmd.Output()  // captures stdout
```

## The Challenge

1. Run `echo "Hello from shell"` using `exec.Command` and print the output
2. Run `ls -la` (or `dir` on Windows) and capture both stdout and stderr separately
3. Build a `runCommand(name string, args ...string) (stdout, stderr string, err error)` helper function
4. Stream output in real-time: run a command with multiple lines of output (`ping -c 4 localhost` or `go version`) and print each line as it arrives using `cmd.StdoutPipe()` + `bufio.Scanner`
5. Check if a command exists on the system using `exec.LookPath`

## What you will practice

- `exec.Command`, `.Output()`, `.Run()`, `.CombinedOutput()`
- `cmd.Stdout`, `cmd.Stderr` assignment
- `cmd.StdoutPipe()` for streaming
- `exec.LookPath` to find executables
- Checking exit codes with `*exec.ExitError`

## Hints

- Separate capture: `cmd.Stdout = &outBuf; cmd.Stderr = &errBuf` (use `bytes.Buffer`)
- Streaming: `pipe, _ := cmd.StdoutPipe(); cmd.Start(); scanner := bufio.NewScanner(pipe)`
- Exit code: `if exitErr, ok := err.(*exec.ExitError); ok { exitCode = exitErr.ExitCode() }`
