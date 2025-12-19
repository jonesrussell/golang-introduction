# Project Layout

**Duration:** 8-10 minutes

## Standard Layout

```
project/
├── cmd/                    # Main applications
│   ├── api/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/               # Private packages
│   ├── auth/
│   ├── database/
│   └── handlers/
├── pkg/                    # Public packages (library)
│   ├── client/
│   └── models/
├── api/                    # API definitions (OpenAPI, proto)
├── web/                    # Web assets
├── configs/                # Configuration files
├── scripts/                # Build/deploy scripts
├── test/                   # Additional test data
├── docs/                   # Documentation
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Key Directories

```go
// cmd/ - Entry points
// cmd/api/main.go
package main

import "github.com/username/project/internal/app"

func main() {
    app.Run()
}

// internal/ - Private to this module
// internal/database/db.go
package database

// Can only be imported by github.com/username/project/...
// NOT importable by other modules!

// pkg/ - Public library code
// pkg/client/client.go
package client

// Importable by anyone
type Client struct {
    baseURL string
}
```

## Key teaching points:
- [`cmd/`](https://github.com/golang-standards/project-layout) contains main applications
- [`internal/`](https://go.dev/doc/go1.4#internalpackages) packages are private to the module
- [`pkg/`](https://github.com/golang-standards/project-layout) contains public library code
- Follow [standard project layout](https://github.com/golang-standards/project-layout) conventions
