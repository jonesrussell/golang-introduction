# Testing CLI Tools

**Duration:** 3-4 minutes

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
