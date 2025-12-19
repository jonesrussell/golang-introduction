# Standard Library flag Package

**Duration:** 6-7 minutes

## Code Examples

```go runnable
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Define flags
    name := flag.String("name", "World", "Name to greet")
    count := flag.Int("count", 1, "Number of greetings")
    verbose := flag.Bool("verbose", false, "Enable verbose output")

    // Custom usage message
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\nOptions:\n", os.Args[0])
        flag.PrintDefaults()
    }

    // Parse flags
    flag.Parse()

    // Access non-flag arguments
    args := flag.Args()
    if len(args) > 0 {
        fmt.Println("Additional arguments:", args)
    }

    // Use flags
    if *verbose {
        fmt.Println("Verbose mode enabled")
    }

    for i := 0; i < *count; i++ {
        fmt.Printf("Hello, %s!\n", *name)
    }
}
```

## Usage

```bash
./greet -name=Alice -count=3 -verbose
./greet --name Alice --count 3
```

## Flag Variations

```go
var (
    host string
    port int
)

func init() {
    flag.StringVar(&host, "host", "localhost", "Server host")
    flag.StringVar(&host, "H", "localhost", "Server host (shorthand)")
    flag.IntVar(&port, "port", 8080, "Server port")
}
```

## Limitations:
- No subcommands
- No short flags (-v)
- Basic type support

## Key teaching points:
- [`flag` package](https://pkg.go.dev/flag) is simple and sufficient for basic CLIs
- Use [`flag.Parse()`](https://pkg.go.dev/flag#Parse) to parse command-line arguments
- Access non-flag arguments with [`flag.Args()`](https://pkg.go.dev/flag#Args)
- For complex CLIs, use [Cobra](https://pkg.go.dev/github.com/spf13/cobra)
