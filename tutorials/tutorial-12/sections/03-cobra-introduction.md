# Cobra Introduction

**Duration:** 5-6 minutes

## Topics:
- [Cobra](https://pkg.go.dev/github.com/spf13/cobra) features
- Command structure
- Installation and setup

## Code Examples

```go
// Install: go get -u github.com/spf13/cobra/cobra

package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

// Root command
var rootCmd = &cobra.Command{
    Use:   "myapp",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines
and provides detailed information about your application.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello from myapp!")
    },
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

## Adding Flags

```go
var verbose bool
var config string

func init() {
    // Persistent flags (available to this command and all subcommands)
    rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

    // Local flags (only this command)
    rootCmd.Flags().StringVarP(&config, "config", "c", "", "config file path")
}
```

## Key teaching points:
- [Cobra](https://pkg.go.dev/github.com/spf13/cobra) provides subcommands, flags, and help generation
- [Persistent flags](https://pkg.go.dev/github.com/spf13/cobra#Command.PersistentFlags) are inherited by subcommands
- [Local flags](https://pkg.go.dev/github.com/spf13/cobra#Command.Flags) are command-specific
- Use [`VarP`](https://pkg.go.dev/github.com/spf13/cobra#Command.Flags) for flags with short and long forms
