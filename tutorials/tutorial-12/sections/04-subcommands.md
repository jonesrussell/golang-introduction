# Subcommands

**Duration:** 6-7 minutes

## Root Command

```go
// cmd/root.go
var rootCmd = &cobra.Command{
    Use:   "fileutil",
    Short: "A file utility tool",
}
```

## List Command

```go
// cmd/list.go
var listCmd = &cobra.Command{
    Use:   "list [directory]",
    Short: "List files in a directory",
    Args:  cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        dir := "."
        if len(args) > 0 {
            dir = args[0]
        }

        showHidden, _ := cmd.Flags().GetBool("all")
        listFiles(dir, showHidden)
    },
}

func init() {
    listCmd.Flags().BoolP("all", "a", false, "Show hidden files")
    rootCmd.AddCommand(listCmd)
}
```

## Copy Command

```go
// cmd/copy.go
var copyCmd = &cobra.Command{
    Use:   "copy <source> <destination>",
    Short: "Copy a file",
    Args:  cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        src, dst := args[0], args[1]
        force, _ := cmd.Flags().GetBool("force")

        return copyFile(src, dst, force)
    },
}

func init() {
    copyCmd.Flags().BoolP("force", "f", false, "Overwrite existing files")
    rootCmd.AddCommand(copyCmd)
}
```

## Search Command

```go
// cmd/search.go
var searchCmd = &cobra.Command{
    Use:   "search <pattern> [directory]",
    Short: "Search for files matching pattern",
    Args:  cobra.RangeArgs(1, 2),
    Run: func(cmd *cobra.Command, args []string) {
        pattern := args[0]
        dir := "."
        if len(args) > 1 {
            dir = args[1]
        }

        recursive, _ := cmd.Flags().GetBool("recursive")
        searchFiles(pattern, dir, recursive)
    },
}

func init() {
    searchCmd.Flags().BoolP("recursive", "r", false, "Search recursively")
    rootCmd.AddCommand(searchCmd)
}
```

## Usage

```bash
fileutil list
fileutil list -a /home/user
fileutil copy src.txt dst.txt -f
fileutil search "*.go" ./src -r
```

## Key teaching points:
- Use [`AddCommand()`](https://pkg.go.dev/github.com/spf13/cobra#Command.AddCommand) to add subcommands
- Use [`Args`](https://pkg.go.dev/github.com/spf13/cobra#Command.Args) validators for argument validation
- [`RunE`](https://pkg.go.dev/github.com/spf13/cobra#Command.RunE) returns errors for better error handling
- Subcommands organize complex CLIs
