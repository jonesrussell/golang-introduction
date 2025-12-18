## **Video Tutorial Plan: Building CLI Tools in Go**

### **Video Metadata**
- **Title:** Building CLI Tools in Go: Cobra and Flag Packages
- **Duration Target:** 40-50 minutes
- **Difficulty:** Intermediate
- **Prerequisites:** Go Basics, Structs, Interfaces

---

## **Video Structure**

### **1. Introduction (3-4 min)**
- Why Go excels at CLI tools
- Overview of options (flag, pflag, Cobra)
- What we'll build: A file utility CLI
- Preview of final tool

---

### **2. Standard Library flag Package (6-7 min)**

**Code Examples:**
```go
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

// Usage:
// ./greet -name=Alice -count=3 -verbose
// ./greet --name Alice --count 3

// Flag variations
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

**Limitations:**
- No subcommands
- No short flags (-v)
- Basic type support

---

### **3. Cobra Introduction (5-6 min)**

**Topics:**
- Cobra features
- Command structure
- Installation and setup

**Code Examples:**
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

// Adding flags
var verbose bool
var config string

func init() {
    // Persistent flags (available to this command and all subcommands)
    rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

    // Local flags (only this command)
    rootCmd.Flags().StringVarP(&config, "config", "c", "", "config file path")
}
```

---

### **4. Subcommands (6-7 min)**

**Code Examples:**
```go
// cmd/root.go
var rootCmd = &cobra.Command{
    Use:   "fileutil",
    Short: "A file utility tool",
}

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

// Usage:
// fileutil list
// fileutil list -a /home/user
// fileutil copy src.txt dst.txt -f
// fileutil search "*.go" ./src -r
```

---

### **5. Configuration with Viper (5-6 min)**

**Code Examples:**
```go
import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var cfgFile string

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
    rootCmd.PersistentFlags().String("database-url", "", "database connection string")

    // Bind flag to viper
    viper.BindPFlag("database-url", rootCmd.PersistentFlags().Lookup("database-url"))
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        home, _ := os.UserHomeDir()
        viper.AddConfigPath(home)
        viper.AddConfigPath(".")
        viper.SetConfigName(".myapp")
        viper.SetConfigType("yaml")
    }

    // Environment variables
    viper.SetEnvPrefix("MYAPP")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}

// Access configuration
func runServer(cmd *cobra.Command, args []string) {
    dbURL := viper.GetString("database-url")
    port := viper.GetInt("port")
    debug := viper.GetBool("debug")

    // Priority: flags > env vars > config file > defaults
}

// Config file (.myapp.yaml):
// database-url: postgres://localhost/mydb
// port: 8080
// debug: true

// Environment: MYAPP_DATABASE_URL=postgres://...
// Flag: --database-url=postgres://...
```

---

### **6. Practical Example: File Utility CLI (10-12 min)**

```go
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
            return nil // Skip errors
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

---

### **7. User Experience (4-5 min)**

```go
// Progress indicators
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

// Confirmation prompts
func confirm(prompt string) bool {
    fmt.Printf("%s [y/N]: ", prompt)
    var response string
    fmt.Scanln(&response)
    return strings.ToLower(response) == "y"
}

// Color output
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

---

### **8. Testing CLI Tools (3-4 min)**

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

---

### **9. Wrap-up (2-3 min)**

**Key takeaways:**
- Use Cobra for complex CLIs
- Implement subcommands for organization
- Provide good help text
- Add verbose/quiet modes
- Use Viper for configuration

**Homework:**
1. Add more commands to file utility
2. Implement config file support
3. Add shell completion
4. Publish to Homebrew/Snap

---

## **Supplementary Materials**

**Cobra Cheat Sheet:**
```go
// Root command
var rootCmd = &cobra.Command{Use: "app"}

// Subcommand
var subCmd = &cobra.Command{Use: "sub", RunE: run}
rootCmd.AddCommand(subCmd)

// Flags
cmd.Flags().StringVarP(&var, "name", "n", "default", "description")
cmd.PersistentFlags()  // Inherited by subcommands

// Arguments
cobra.NoArgs
cobra.ExactArgs(n)
cobra.MinimumNArgs(n)
cobra.MaximumNArgs(n)
cobra.RangeArgs(min, max)
```

---

This tutorial provides practical guidance on building professional CLI tools with Go.
