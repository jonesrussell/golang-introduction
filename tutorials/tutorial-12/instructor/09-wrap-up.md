# Instructor Notes: Wrap Up

## Recap What Was Covered
- Flag package (simple tools)
- Cobra (complex tools)
- Subcommands
- Configuration with Viper
- Practical example
- User experience
- Testing

## Key Takeaways to Emphasize
- Flag package for simple tools
- Cobra for complex tools
- Subcommands organize functionality
- Good UX matters

## Preview Next Tutorial
- "Next: Packages and Modules - organizing code"
- "You'll learn Go's module system"
- "We'll build on everything you learned today"

## Practice Recommendations
- **Easy**: Build a simple CLI tool with flag
- **Medium**: Build a tool with Cobra
- **Challenge**: Build a complex tool with subcommands

## Cheat Sheet Highlights
- Flag: `flag.String("name", "default", "usage")`
- Cobra: `cobra.Command{Use: "cmd", Run: func}`
- Subcommands: `rootCmd.AddCommand(subCmd)`
- Viper: `viper.BindPFlags(cmd.Flags())`

## Engagement
- "What CLI tool would you build?"
- "Try building a tool with Cobra"
- "Questions? Let's address them now"

## Closing
- "You've mastered CLI tools in Go"
- "Go is excellent for building CLI tools"
- "Next tutorial: Packages and Modules - see you there!"
