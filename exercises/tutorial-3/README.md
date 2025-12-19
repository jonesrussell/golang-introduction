# Tutorial 3 Exercises: Struct Embedding and Composition

Practice exercises for struct composition and embedding patterns.

## How to Use

1. Navigate to an exercise folder
2. Read the instructions in the `main.go` file
3. Complete the TODO sections
4. Run your code: `go run main.go`
5. Check against `solution/main.go` if stuck

## Exercises

| # | Exercise | Topics |
|---|----------|--------|
| 1 | Basic Composition | named fields, has-a relationships |
| 2 | Struct Embedding | anonymous fields, field promotion |
| 3 | Method Promotion | accessing embedded methods directly |
| 4 | Multiple Embedding | embedding multiple types, conflicts |
| 5 | Practical Patterns | real-world embedding patterns |

## Running Exercises

```bash
cd 01-basic-composition
go run main.go

# Check solution
go run solution/main.go
```

## Key Concepts

- **Composition** uses named fields: `Address Address`
- **Embedding** uses anonymous fields: `Address` (no field name)
- **Field promotion** lets you access embedded fields directly
- **Method promotion** lets you call embedded methods directly
- Embedding looks like inheritance but IS composition
