# Wrap-up

**Duration:** 2-3 minutes

## Key Takeaways:
- Use Cobra for complex CLIs
- Implement subcommands for organization
- Provide good help text
- Add verbose/quiet modes
- Use Viper for configuration

## Homework:
1. Add more commands to file utility
2. Implement config file support
3. Add shell completion
4. Publish to Homebrew/Snap

## Cobra Cheat Sheet

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
