# Go Tutorial Exercises

Practice exercises to reinforce concepts from each tutorial.

## Available Exercises

| Tutorial | Topic | Exercises |
|----------|-------|-----------|
| 1 | Go Basics | 6 exercises: variables, types, control flow |
| 2 | Structs | 6 exercises: definition, methods, constructors |
| 3 | Embedding & Composition | 5 exercises: composition, embedding, promotion |

## Structure

```
exercises/
├── tutorial-1/          # Go Basics
│   ├── 01-hello-world/
│   ├── 02-variables/
│   ├── 03-types/
│   ├── 04-conditionals/
│   ├── 05-loops/
│   └── 06-switch/
├── tutorial-2/          # Structs
│   ├── 01-defining-structs/
│   ├── 02-initialization/
│   ├── 03-nested-structs/
│   ├── 04-value-receiver-methods/
│   ├── 05-pointer-receiver-methods/
│   └── 06-constructor-pattern/
└── tutorial-3/          # Embedding & Composition
    ├── 01-basic-composition/
    ├── 02-struct-embedding/
    ├── 03-method-promotion/
    ├── 04-multiple-embedding/
    └── 05-practical-patterns/
```

## How to Use

1. Navigate to an exercise: `cd tutorial-1/01-hello-world`
2. Read the instructions in `main.go`
3. Complete the TODO sections
4. Run: `go run main.go`
5. Compare with `solution/main.go` if stuck

## Tips

- Start with exercise 01 and work through in order
- Read error messages carefully - Go is helpful
- Use `go fmt main.go` to auto-format your code
- Experiment! Modify the code and see what happens
