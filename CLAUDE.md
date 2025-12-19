# CLAUDE.md

This file provides context for Claude (AI assistant) when working with this codebase.

## Project Overview

**Go Fundamentals & Best Practices** - An interactive tutorial platform for learning Go (Golang). Features a Go backend API and Vue.js frontend that allows users to read tutorials, run code examples, and track progress.

**Repository:** `github.com/jonesrussell/go-fundamentals-best-practices`

## Quick Commands

```bash
# Development
task dev:backend          # Start Go backend server (port 8080)
task dev:frontend         # Start Vue frontend dev server (port 5173)
task dev                  # Start both (requires two terminals)

# Build
task build               # Build both backend and frontend
go build ./...           # Build Go code only

# Linting & Quality
task lint                # Run all linters (Go + frontend type-check)
task fmt                 # Format Go code
golangci-lint run ./...  # Run Go linter directly

# Testing
task test                # Run Go tests
task test:coverage       # Run tests with coverage report

# Docker
docker compose up        # Start containerized services
```

## Architecture

### Backend (Go)

```
cmd/server/main.go        # Entry point
internal/
├── api/
│   ├── handlers.go       # HTTP request handlers
│   └── routes.go         # Route definitions
├── executor/
│   ├── runner.go         # Code execution service
│   └── wrapper.go        # Snippet wrapping (auto-adds package main)
├── parser/
│   ├── markdown.go       # Legacy single-file parser
│   ├── section.go        # New directory-based parser
│   ├── tutorial.go       # Tutorial orchestrator
│   └── goldmark_parser.go # AST-based markdown parsing
└── storage/
    └── progress.go       # User progress persistence
pkg/models/
├── tutorial.go           # Tutorial, Section, CodeExample, Exercise
└── progress.go           # UserProgress model
```

### Frontend (Vue 3 + TypeScript)

```
frontend/src/
├── components/
│   ├── TutorialViewer.vue   # Main tutorial display
│   ├── TutorialList.vue     # Sidebar tutorial list
│   ├── SectionViewer.vue    # Section content renderer
│   ├── CodeRunner.vue       # Code execution UI
│   ├── CodeEditor.vue       # Code editing textarea
│   ├── InstructorPanel.vue  # Instructor notes display
│   └── ExerciseView.vue     # Exercise component
├── composables/
│   ├── useTutorial.ts       # Tutorial data fetching/caching
│   ├── useCodeExecution.ts  # Code execution logic
│   └── useSyntaxHighlight.ts # Shiki syntax highlighting
├── stores/
│   └── progress.ts          # Pinia store for progress tracking
├── services/
│   └── api.ts               # Axios API client
├── router/
│   └── index.ts             # Vue Router configuration
└── views/
    ├── HomeView.vue         # Home page
    └── TutorialView.vue     # Tutorial route wrapper
```

## Tutorial Content Structure

### New Directory-Based Format (Preferred)

```
tutorials/
└── tutorial-1/
    ├── tutorial.yaml         # Tutorial metadata
    ├── sections/
    │   ├── 01-introduction.md
    │   ├── 02-hello-world.md
    │   └── ...
    └── instructor/
        ├── 01-introduction.md  # Instructor-only notes
        ├── 02-hello-world.md
        └── ...
```

**tutorial.yaml:**
```yaml
id: "1"
title: "Go Basics: Variables, Types, and Control Flow"
duration: "25-35 minutes"
difficulty: "Beginner"
level: "Beginner"
prerequisites:
  - "Basic programming concepts helpful but not required"
```

### Legacy Single-File Format

```
tutorials/
├── Tutorial-1-Go-Basics-for-Beginners.md
├── Tutorial-2-Go-Structs-Definition-Initialization-and-Methods.md
└── ...
```

The parser auto-detects format and prefers directory-based when available.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/tutorials` | List all tutorials (metadata only) |
| GET | `/api/tutorials/:id` | Get full tutorial with sections |
| GET | `/api/tutorials/:id?instructor=true` | Include instructor notes |
| GET | `/api/tutorials/:id/sections` | Get sections only |
| POST | `/api/execute` | Execute Go code |
| GET | `/api/progress` | Get user progress |
| POST | `/api/progress` | Update user progress |
| POST | `/api/progress/section` | Mark section complete |
| GET | `/api/exercises/:tutorialId` | Get exercises for tutorial |

## Code Examples in Markdown

Use these attributes in code blocks:

```markdown
<!-- Full program - auto-detected as runnable -->
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
```

<!-- Snippet - auto-wrapped with package main -->
```go snippet
fmt.Println("Hello, World!")
```

<!-- Explicit runnable flag -->
```go runnable
package main
// ...
```
```

The `snippet` attribute triggers auto-wrapping with `package main`, `func main()`, and detected imports.

## Key Models

### Tutorial
```go
type Tutorial struct {
    ID            string    `json:"id"`
    Title         string    `json:"title"`
    Duration      string    `json:"duration"`
    Difficulty    string    `json:"difficulty"`
    Level         string    `json:"level"`        // Beginner, Intermediate, Advanced
    Prerequisites []string  `json:"prerequisites"`
    Sections      []Section `json:"sections"`
}
```

### Section
```go
type Section struct {
    ID              string        `json:"id"`
    Title           string        `json:"title"`
    Topics          []string      `json:"topics"`
    CodeExamples    []CodeExample `json:"codeExamples"`
    TeachingPoints  []string      `json:"teachingPoints"`
    Order           int           `json:"order"`
    Content         string        `json:"content"`
    InstructorNotes string        `json:"instructorNotes,omitempty"`
}
```

### CodeExample
```go
type CodeExample struct {
    ID             string `json:"id"`
    Code           string `json:"code"`
    Language       string `json:"language"`
    Runnable       bool   `json:"runnable"`
    Snippet        bool   `json:"snippet,omitempty"`  // Needs wrapping
    ExpectedOutput string `json:"expectedOutput,omitempty"`
}
```

## Frontend Routes

| Path | Component | Description |
|------|-----------|-------------|
| `/` | HomeView | Welcome screen |
| `/tutorial/:id` | TutorialView | Tutorial viewer |
| `/tutorial/:id/section/:sectionIndex` | TutorialView | Direct section link |
| `/about` | AboutView | About page |

## Instructor Mode

Toggle in `TutorialViewer.vue` header. Persisted to `localStorage`. When enabled:
- API requests include `?instructor=true`
- `InstructorPanel` component displays notes
- Notes come from `tutorials/tutorial-X/instructor/*.md` files

## Code Execution Security

The executor (`internal/executor/runner.go`) validates code before running:
- **Blocked imports:** `net`, `os/exec`, `syscall`, `unsafe`, `plugin`
- **Blocked operations:** Network access, file I/O outside temp dir
- **Timeout:** 10 seconds default
- **Isolated:** Runs in temporary directory, cleaned up after

## Tech Stack

**Backend:**
- Go 1.25+
- `net/http` (standard library)
- `github.com/yuin/goldmark` (markdown parsing)
- `gopkg.in/yaml.v3` (YAML parsing)

**Frontend:**
- Vue 3 (Composition API)
- TypeScript 5.9
- Vite 7
- Pinia (state management)
- Vue Router 4
- Tailwind CSS 4
- Shiki (syntax highlighting)
- Axios (HTTP client)

## Linting

Uses `golangci-lint` with strict configuration (`.golangci.yml`). Key linters:
- `errcheck` - Check error returns
- `govet` - Go vet with shadow checking
- `gocognit` - Cognitive complexity (max 20)
- `funlen` - Function length (max 100 lines)
- `gosec` - Security checks

Frontend uses `vue-tsc` for TypeScript checking.

## Common Tasks

### Adding a New Tutorial (Directory Format)

1. Create directory: `tutorials/tutorial-N/`
2. Create `tutorial.yaml` with metadata
3. Create `sections/` with numbered markdown files
4. Optionally create `instructor/` with matching filenames

### Adding a New API Endpoint

1. Add handler function in `internal/api/handlers.go`
2. Register route in `internal/api/routes.go`
3. Update frontend `services/api.ts` if needed

### Modifying Tutorial Models

1. Update `pkg/models/tutorial.go`
2. Update parsers in `internal/parser/`
3. Update frontend types in `frontend/src/types/tutorial.ts`
4. Update API service if needed

## User Preferences (Stored in localStorage)

- `instructor-mode`: `"true"` or `"false"` - Instructor mode toggle
- `tutorial-progress`: JSON object with completion data

## Environment

- Backend runs on port 8080
- Frontend dev server runs on port 5173
- CORS enabled for cross-origin requests

## Useful Files

- `Taskfile.yml` - Task runner configuration
- `docker-compose.yml` - Docker setup
- `.golangci.yml` - Linter configuration
- `DEVELOPMENT_STATUS.md` - Feature completion status
- `TECHNICAL_STACK.md` - Detailed tech stack
- `SUCCESS_CRITERIA.md` - Project requirements
