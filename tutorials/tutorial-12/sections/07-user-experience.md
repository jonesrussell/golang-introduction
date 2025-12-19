# User Experience

**Duration:** 4-5 minutes

## Progress Indicators

## Key teaching points:
- Provide feedback for long-running operations
- Use confirmation prompts for destructive actions
- Color output improves readability (use libraries like [fatih/color](https://github.com/fatih/color))

```go
func copyWithProgress(src, dst string, size int64) error {
    srcFile, _ := os.Open(src)
    defer srcFile.Close()
    dstFile, _ := os.Create(dst)
    defer dstFile.Close()

    var written int64
    buf := make([]byte, 32*1024)

    for {
        n, err := srcFile.Read(buf)
        if n > 0 {
            dstFile.Write(buf[:n])
            written += int64(n)
            printProgress(written, size)
        }
        if err == io.EOF {
            break
        }
    }
    fmt.Println()
    return nil
}

func printProgress(current, total int64) {
    percent := float64(current) / float64(total) * 100
    fmt.Printf("\rProgress: %.1f%%", percent)
}
```

## Confirmation Prompts

```go
func confirm(prompt string) bool {
    fmt.Printf("%s [y/N]: ", prompt)
    var response string
    fmt.Scanln(&response)
    return strings.ToLower(response) == "y"
}
```

## Color Output

```go
import "github.com/fatih/color"

var (
    success = color.New(color.FgGreen).SprintFunc()
    warning = color.New(color.FgYellow).SprintFunc()
    danger  = color.New(color.FgRed).SprintFunc()
)

fmt.Println(success("Operation completed"))
fmt.Println(warning("Warning: file exists"))
fmt.Println(danger("Error: permission denied"))
```
