# Testing CLI Tools

**Duration:** 3-4 minutes

## Key teaching points:
- Test CLI commands with [Cobra](https://pkg.go.dev/github.com/spf13/cobra) by setting output and args
- Use [`t.TempDir()`](https://pkg.go.dev/testing#T.TempDir) for temporary test directories
- Capture output with [`bytes.Buffer`](https://pkg.go.dev/bytes#Buffer) for assertions

```go
func TestListCommand(t *testing.T) {
    // Create temp directory
    dir := t.TempDir()
    os.WriteFile(filepath.Join(dir, "test.txt"), []byte("hello"), 0644)

    // Capture output
    var buf bytes.Buffer
    rootCmd.SetOut(&buf)
    rootCmd.SetArgs([]string{"list", dir})

    err := rootCmd.Execute()
    if err != nil {
        t.Fatal(err)
    }

    if !strings.Contains(buf.String(), "test.txt") {
        t.Error("expected test.txt in output")
    }
}
```
