# Tutorial 2 Exercises: Go Structs

Practice exercises for struct definition, initialization, and methods.

## How to Use

1. Navigate to an exercise folder
2. Read the instructions in the `main.go` file
3. Complete the TODO sections
4. Run your code: `go run main.go`
5. Check against `solution/main.go` if stuck

## Exercises

| # | Exercise | Topics |
|---|----------|--------|
| 1 | Defining Structs | struct definition, field types, naming |
| 2 | Initialization | zero values, literal syntax, named fields |
| 3 | Nested Structs | composition, nested field access |
| 4 | Value Receiver Methods | read-only methods, Area(), Perimeter() |
| 5 | Pointer Receiver Methods | mutating methods, Scale(), Deposit() |
| 6 | Constructor Pattern | NewXxx() constructors, returning pointers |

## Running Exercises

```bash
cd 01-defining-structs
go run main.go

# Check solution
go run solution/main.go
```

## Key Concepts

- **Exported fields** start with uppercase (public)
- **Value receivers** get a copy - cannot modify original
- **Pointer receivers** get a pointer - can modify original
- **Constructors** are functions that return `*StructType`
