# Go Fundamentals & Best Practices

An interactive tutorial platform for learning Go (Golang) from basics to advanced concepts.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://golang.org/)

## Features

- **Interactive tutorials** - Step-by-step lessons with runnable code examples
- **In-browser code execution** - Run Go code directly in your browser
- **Progress tracking** - Track your learning progress across tutorials
- **Instructor mode** - Teaching notes and tips for educators

## Quick Start

```bash
# Clone the repository
git clone https://github.com/jonesrussell/go-fundamentals-best-practices.git
cd go-fundamentals-best-practices

# Start the backend (port 8080)
go run cmd/server/main.go

# In another terminal, start the frontend (port 5173)
cd frontend
npm install
npm run dev
```

Open http://localhost:5173 in your browser.

### Using Docker

```bash
docker compose up
```

## Tutorial Series

### Beginner

| # | Tutorial | Duration |
|---|----------|----------|
| 1 | Go Basics: Variables, Types, and Control Flow | 25-35 min |
| 2 | Go Structs: Definition, Initialization, and Methods | 30-40 min |
| 3 | Struct Embedding and Composition | 35-45 min |

### Intermediate

| # | Tutorial | Duration |
|---|----------|----------|
| 4 | Pointers in Go | 30-40 min |
| 5 | Understanding Go Interfaces | 40-50 min |
| 6 | Error Handling in Go | 35-45 min |
| 7 | Go Concurrency: Goroutines and Channels | 45-55 min |
| 8 | Go Slices and Maps | 35-45 min |

### Advanced

| # | Tutorial | Duration |
|---|----------|----------|
| 9 | Dependency Injection in Go | 40-50 min |
| 10 | Avoiding Common Go Anti-Patterns | 35-45 min |
| 11 | Structured Logging with Zap | 30-40 min |
| 12 | Building CLI Tools in Go | 40-50 min |
| 13 | Go Packages and Modules | 35-45 min |

## Project Structure

```
├── cmd/server/          # Backend API server
├── internal/
│   ├── api/             # HTTP handlers and routes
│   ├── executor/        # Go code execution service
│   ├── parser/          # Markdown tutorial parser
│   └── storage/         # Progress tracking
├── frontend/            # Vue.js frontend
├── tutorials/           # Tutorial content (markdown)
└── cheatsheets/         # Quick reference guides
```

## Tech Stack

**Backend:** Go 1.25+, standard library `net/http`

**Frontend:** Vue 3, TypeScript, Vite, Pinia, Tailwind CSS, Shiki

## Development

```bash
# Run with task runner
task dev:backend    # Start backend
task dev:frontend   # Start frontend
task lint           # Run linters
task test           # Run tests
```

## Resources

- [Go Tour](https://tour.golang.org) - Interactive Go tutorial
- [Go Playground](https://play.golang.org) - Run Go in browser
- [Effective Go](https://golang.org/doc/effective_go) - Official best practices

## License

MIT License - see [LICENSE](LICENSE) for details.

## Author

**Russell Jones** - [FullStackDev42](https://www.youtube.com/@fullstackdev42)
