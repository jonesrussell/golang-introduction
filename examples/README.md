# Code Examples

This directory contains code examples for each tutorial in the series.

## Directory Structure

```
examples/
├── 01-basics/           # Variables, types, control flow
├── 02-structs/          # Struct definition, methods
├── 03-embedding/        # Composition and embedding
├── 04-pointers/         # Pointers and memory
├── 05-interfaces/       # Interface patterns
├── 06-error-handling/   # Error types and handling
├── 07-concurrency/      # Goroutines and channels
├── 08-slices-maps/      # Collections
├── 09-dependency-injection/ # DI patterns
├── 10-anti-patterns/    # What to avoid
├── 11-logging/          # Structured logging
├── 12-cli-tools/        # CLI development
└── 13-packages/         # Module organization
```

## Running Examples

Each example directory contains standalone Go files that can be run with:

```bash
cd examples/01-basics
go run main.go
```

Or for directories with multiple files:

```bash
go run .
```

## Prerequisites

- Go 1.21 or later
- Some examples require additional dependencies (documented in each directory)
