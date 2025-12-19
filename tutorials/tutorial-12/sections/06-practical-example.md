# Practical Example: File Utility CLI

**Duration:** 10-12 minutes

```go runnable
package main

import (
    "fmt"
    "io"
    "io/fs"
    "os"
    "path/filepath"
    "strings"

    "github.com/spf13/cobra"
)

var (
    verbose bool
    rootCmd = &cobra.Command{
        Use:   "fileutil",
        Short: "A file utility tool",
        Long:  "A CLI tool for common file operations",
    }
)

func init() {
    rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

    rootCmd.AddCommand(listCmd)
    rootCmd.AddCommand(copyCmd)
    rootCmd.AddCommand(searchCmd)
    rootCmd.AddCommand(statsCmd)
}

// List command
var showHidden, longFormat bool
var listCmd = &cobra.Command{
    Use:     "list [directory]",
    Aliases: []string{"ls", "l"},
    Short:   "List files in a directory",
    Args:    cobra.MaximumNArgs(1),
    RunE:    runList,
}

func init() {
    listCmd.Flags().BoolVarP(&showHidden, "all", "a", false, "Show hidden files")
    listCmd.Flags().BoolVarP(&longFormat, "long", "l", false, "Long format")
}

func runList(cmd *cobra.Command, args []string) error {
    dir := "."
    if len(args) > 0 {
        dir = args[0]
    }

    entries, err := os.ReadDir(dir)
    if err != nil {
        return fmt.Errorf("reading directory: %w", err)
    }

    for _, entry := range entries {
        name := entry.Name()
        if !showHidden && strings.HasPrefix(name, ".") {
            continue
        }

        if longFormat {
            info, _ := entry.Info()
            fmt.Printf("%s %10d %s %s\n",
                info.Mode(),
                info.Size(),
                info.ModTime().Format("Jan 02 15:04"),
                name)
        } else {
            if entry.IsDir() {
                fmt.Printf("%s/\n", name)
            } else {
                fmt.Println(name)
            }
        }
    }
    return nil
}

// Copy command
var force, recursive bool
var copyCmd = &cobra.Command{
    Use:     "copy <source> <destination>",
    Aliases: []string{"cp"},
    Short:   "Copy files or directories",
    Args:    cobra.ExactArgs(2),
    RunE:    runCopy,
}

func init() {
    copyCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing")
    copyCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Copy directories recursively")
}

func runCopy(cmd *cobra.Command, args []string) error {
    src, dst := args[0], args[1]

    srcInfo, err := os.Stat(src)
    if err != nil {
        return fmt.Errorf("source: %w", err)
    }

    if srcInfo.IsDir() {
        if !recursive {
            return fmt.Errorf("source is directory, use -r")
        }
        return copyDir(src, dst)
    }

    return copyFile(src, dst)
}

func copyFile(src, dst string) error {
    if !force {
        if _, err := os.Stat(dst); err == nil {
            return fmt.Errorf("destination exists, use -f to overwrite")
        }
    }

    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    if verbose {
        fmt.Printf("Copied: %s -> %s\n", src, dst)
    }
    return err
}

func copyDir(src, dst string) error {
    return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return err
        }

        relPath, _ := filepath.Rel(src, path)
        dstPath := filepath.Join(dst, relPath)

        if info.IsDir() {
            return os.MkdirAll(dstPath, info.Mode())
        }
        return copyFile(path, dstPath)
    })
}

// Search command
var searchRecursive bool
var searchCmd = &cobra.Command{
    Use:     "search <pattern> [directory]",
    Aliases: []string{"find"},
    Short:   "Search for files",
    Args:    cobra.RangeArgs(1, 2),
    RunE:    runSearch,
}

func init() {
    searchCmd.Flags().BoolVarP(&searchRecursive, "recursive", "r", false, "Search recursively")
}

func runSearch(cmd *cobra.Command, args []string) error {
    pattern := args[0]
    dir := "."
    if len(args) > 1 {
        dir = args[1]
    }

    count := 0
    err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return nil
        }

        if !searchRecursive && filepath.Dir(path) != dir {
            if info.IsDir() {
                return filepath.SkipDir
            }
            return nil
        }

        matched, _ := filepath.Match(pattern, info.Name())
        if matched {
            fmt.Println(path)
            count++
        }
        return nil
    })

    if verbose {
        fmt.Printf("\nFound %d files\n", count)
    }
    return err
}

// Stats command
var statsCmd = &cobra.Command{
    Use:   "stats <directory>",
    Short: "Show directory statistics",
    Args:  cobra.ExactArgs(1),
    RunE:  runStats,
}

func runStats(cmd *cobra.Command, args []string) error {
    dir := args[0]

    var totalSize int64
    fileCount, dirCount := 0, 0
    extCounts := make(map[string]int)

    err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return nil
        }

        if info.IsDir() {
            dirCount++
        } else {
            fileCount++
            totalSize += info.Size()
            ext := filepath.Ext(info.Name())
            if ext != "" {
                extCounts[ext]++
            }
        }
        return nil
    })

    if err != nil {
        return err
    }

    fmt.Printf("Directory: %s\n", dir)
    fmt.Printf("Files: %d\n", fileCount)
    fmt.Printf("Directories: %d\n", dirCount)
    fmt.Printf("Total Size: %s\n", formatBytes(totalSize))
    fmt.Println("\nFile types:")
    for ext, count := range extCounts {
        fmt.Printf("  %s: %d\n", ext, count)
    }

    return nil
}

func formatBytes(b int64) string {
    const unit = 1024
    if b < unit {
        return fmt.Sprintf("%d B", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```
