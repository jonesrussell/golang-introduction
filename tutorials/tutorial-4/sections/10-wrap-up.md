# Next Steps & Wrap-up

**Duration:** 2-3 minutes

## Recap What Was Covered:
- Pointer basics (`&`, `*`, nil)
- Pass by value vs pass by pointer
- Pointers to structs
- When to use pointers
- Common gotchas and safety
- Reference types (slices, maps)
- Best practices

## Preview Next Topics:
- Interfaces and polymorphism
- Error handling
- Concurrency (pointers with goroutines)

## Practice Suggestions:
1. **Easy:** Implement a swap function using pointers
2. **Medium:** Build a doubly linked list
3. **Challenge:** Implement a binary tree with insert/search
4. **Advanced:** Build a memory pool using pointers

## Resources:
- Effective Go on Pointers
- Go FAQ on pass by value

## Cheat Sheet

```
&x          Get address of x
*ptr        Dereference pointer
*Type       Pointer type
new(Type)   Allocate and return pointer
nil         Zero value for pointers

Value receiver:   func (t Type) Method()
Pointer receiver: func (t *Type) Method()
```

## Key teaching points:
- Pointers are essential for mutation and efficiency
- Go makes pointers safer than C
- Know when to use value vs pointer semantics
- Always check for nil when working with pointers
