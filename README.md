# Go Fundamentals & Best Practices

A comprehensive video tutorial series for learning Go (Golang) from basics to advanced concepts.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org/)
[![YouTube](https://img.shields.io/badge/YouTube-@fullstackdev42-red?logo=youtube)](https://www.youtube.com/@fullstackdev42)

## About

This repository contains a complete video tutorial series for [FullStackDev42](https://www.youtube.com/@fullstackdev42) YouTube channel. It includes:

- Video tutorial plans and scripts
- Code examples for each tutorial
- Practice exercises with solutions
- Cheat sheets and quick references
- Interactive learning materials

## Tutorial Series

### Beginner Level

| # | Tutorial | Duration | Topics |
|---|----------|----------|--------|
| 1 | [Go Basics: Variables, Types, and Control Flow](tutorials/Tutorial-1-Go-Basics-for-Beginners.md) | 25-35 min | Variables, types, if/else, loops, switch |
| 2 | [Go Structs: Definition, Initialization, and Methods](tutorials/Tutorial-2-Go-Structs-Definition-Initialization-and-Methods.md) | 30-40 min | Struct definition, initialization, methods, receivers |
| 3 | [Struct Embedding and Composition in Go](tutorials/Tutorial-3-Struct-Embedding-and-Composition-in-Go.md) | 35-45 min | Composition, embedding, mixins, patterns |

### Intermediate Level

| # | Tutorial | Duration | Topics |
|---|----------|----------|--------|
| 4 | [Pointers in Go: When to Use *Type vs Type](tutorials/Tutorial-4-Pointers-in-Go.md) | 30-40 min | Pointers, memory, pass by value/reference |
| 5 | [Understanding Go Interfaces](tutorials/Tutorial-5-Go-Interfaces.md) | 40-50 min | Interface definition, implicit implementation, polymorphism |
| 6 | [Error Handling in Go](tutorials/Tutorial-6-Error-Handling-in-Go.md) | 35-45 min | Error types, wrapping, custom errors, patterns |
| 7 | [Go Concurrency: Goroutines and Channels](tutorials/Tutorial-7-Go-Concurrency.md) | 45-55 min | Goroutines, channels, select, sync patterns |
| 8 | [Go Slices and Maps: Internals and Best Practices](tutorials/Tutorial-8-Slices-and-Maps.md) | 35-45 min | Slice internals, maps, iteration, common patterns |

### Advanced Level

| # | Tutorial | Duration | Topics |
|---|----------|----------|--------|
| 9 | [Dependency Injection in Go](tutorials/Tutorial-9-Dependency-Injection.md) | 40-50 min | DI patterns, testability, wire |
| 10 | [Avoiding Common Go Anti-Patterns](tutorials/Tutorial-10-Go-Anti-Patterns.md) | 35-45 min | context.Value, global state, interface pollution |
| 11 | [Structured Logging with Zap](tutorials/Tutorial-11-Structured-Logging-Zap.md) | 30-40 min | Zap logger, structured logging, production practices |
| 12 | [Building CLI Tools in Go](tutorials/Tutorial-12-Building-CLI-Tools.md) | 40-50 min | Cobra, flags, configuration, UX |
| 13 | [Go Packages and Modules](tutorials/Tutorial-13-Packages-and-Modules.md) | 35-45 min | Module system, visibility, organization |

## Repository Structure

```
golang-introduction/
├── README.md
├── LICENSE
├── tutorials/           # Video tutorial plans
│   ├── Tutorial-1-Go-Basics-for-Beginners.md
│   ├── Tutorial-2-Go-Structs-Definition-Initialization-and-Methods.md
│   └── ...
├── examples/            # Code examples by tutorial
│   ├── 01-basics/
│   ├── 02-structs/
│   └── ...
├── exercises/           # Practice exercises
│   ├── 01-basics/
│   ├── 02-structs/
│   └── ...
└── cheatsheets/         # Quick reference guides
    ├── variables-types.md
    ├── control-flow.md
    └── ...
```

## Getting Started

### Prerequisites

- [Go 1.21+](https://golang.org/dl/) installed
- A code editor (VS Code or Cursor recommended)
- Basic programming knowledge (helpful but not required)

### Quick Start

```bash
# Clone the repository
git clone https://github.com/jonesrussell/golang-introduction.git
cd golang-introduction

# Navigate to an example
cd examples/01-basics

# Run the example
go run main.go
```

## Learning Path

### For Absolute Beginners
Start with Tutorial 1 and progress sequentially through Tutorials 1-3. These cover the fundamental building blocks.

### For Developers from Other Languages
If you have experience with another programming language:
- Skim Tutorial 1 for Go syntax specifics
- Focus on Tutorials 2-3 (structs and composition - Go's approach to OOP)
- Pay special attention to Tutorial 4 (pointers) and Tutorial 5 (interfaces)

### For Intermediate Go Developers
Jump to the Advanced section (Tutorials 9-13) for production patterns and best practices.

## Resources

### Official Go Resources
- [Go Tour](https://tour.golang.org) - Interactive Go tutorial
- [Go Playground](https://play.golang.org) - Run Go code in browser
- [Effective Go](https://golang.org/doc/effective_go) - Official best practices
- [Go by Example](https://gobyexample.com) - Annotated example programs

### This Series
- [YouTube Channel](https://www.youtube.com/@fullstackdev42)
- [Cheat Sheets](cheatsheets/)
- [Code Examples](examples/)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Russell Jones** - [FullStackDev42](https://www.youtube.com/@fullstackdev42)

## Acknowledgments

- The Go team for creating an amazing language
- The Go community for excellent documentation and examples
- All viewers and contributors who help improve this series
