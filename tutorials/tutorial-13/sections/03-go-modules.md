# Go Modules

**Duration:** 7-8 minutes

## Topics:
- [Module initialization](https://go.dev/doc/modules/gomod-ref)
- [go.mod file](https://go.dev/doc/modules/gomod-ref)
- Dependencies
- Versioning

## Commands

```bash
# Initialize a new module
go mod init github.com/username/project

# Add a dependency
go get github.com/spf13/cobra@latest
go get github.com/spf13/cobra@v1.7.0

# Update dependencies
go get -u ./...

# Remove unused dependencies
go mod tidy

# Download dependencies
go mod download

# Verify dependencies
go mod verify

# View dependency graph
go mod graph
```

## go.mod File

```go
module github.com/username/project

go 1.21

require (
    github.com/spf13/cobra v1.7.0
    github.com/spf13/viper v1.16.0
)

require (
    // Indirect dependencies (transitive)
    github.com/fsnotify/fsnotify v1.6.0 // indirect
)

// Replace for local development
replace github.com/username/otherproject => ../otherproject

// Exclude problematic versions
exclude github.com/old/package v1.0.0
```

## go.sum File
- Cryptographic checksums
- Ensures reproducible builds
- Commit to version control

## Key teaching points:
- [Modules](https://go.dev/doc/modules/gomod-ref) are collections of Go packages
- [`go mod init`](https://pkg.go.dev/cmd/go#hdr-Initialize_new_module_in_current_directory) creates a new module
- [`go.mod`](https://go.dev/doc/modules/gomod-ref) defines module path and dependencies
- [`go.sum`](https://go.dev/doc/modules/gomod-ref) ensures dependency integrity
- Use [`go mod tidy`](https://pkg.go.dev/cmd/go#hdr-Add_missing_and_remove_unused_modules) to clean dependencies
