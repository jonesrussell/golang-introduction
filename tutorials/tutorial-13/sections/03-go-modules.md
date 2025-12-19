# Go Modules

**Duration:** 7-8 minutes

## Topics:
- Module initialization
- go.mod file
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
