# Configuration with Viper

**Duration:** 5-6 minutes

## Code Examples

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
```

## Access Configuration

```go
func runServer(cmd *cobra.Command, args []string) {
    dbURL := viper.GetString("database-url")
    port := viper.GetInt("port")
    debug := viper.GetBool("debug")

    // Priority: flags > env vars > config file > defaults
}
```

## Config File Example (.myapp.yaml)

```yaml
database-url: postgres://localhost/mydb
port: 8080
debug: true
```

## Environment Variables

```bash
MYAPP_DATABASE_URL=postgres://...
```

## Key teaching points:
- [Viper](https://pkg.go.dev/github.com/spf13/viper) provides configuration management
- Priority: flags > environment variables > config file > defaults
- Use [`BindPFlag()`](https://pkg.go.dev/github.com/spf13/viper#BindPFlag) to bind flags to viper
- Supports multiple config formats (YAML, JSON, TOML)
