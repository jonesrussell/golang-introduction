# Next Steps & Wrap-up

**Duration:** 2-3 minutes

## Recap What Was Covered:
- [Pointer basics](https://go.dev/ref/spec#Pointer_types) (`&`, `*`, nil)
- [Pass by value](https://go.dev/doc/faq#pass_by_value) vs pass by pointer
- [Pointers to structs](https://go.dev/ref/spec#Pointer_types)
- When to use pointers
- Common gotchas and safety
- [Reference types](https://go.dev/ref/spec#Slice_types) (slices, maps)
- Best practices

## Preview Next Topics:
- [Interfaces](https://go.dev/ref/spec#Interface_types) and [polymorphism](https://go.dev/doc/faq#polymorphism)
- [Error handling](https://go.dev/doc/effective_go#errors)
- Concurrency (pointers with goroutines)

## Practice Suggestions:
1. **Easy:** Implement a swap function using pointers
2. **Medium:** Build a doubly linked list
3. **Challenge:** Implement a binary tree with insert/search
4. **Advanced:** Build a memory pool using pointers

## Resources:
- [Effective Go on Pointers](https://go.dev/doc/effective_go#pointers): go.dev/doc/effective_go#pointers
- [Go FAQ on pass by value](https://go.dev/doc/faq#pass_by_value): go.dev/doc/faq#pass_by_value
- [Go Tour - Pointers](https://go.dev/tour/moretypes/1): go.dev/tour/moretypes/1

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
- [Pointers](https://go.dev/ref/spec#Pointer_types) are essential for mutation and efficiency
- Go makes pointers [safer than C](https://go.dev/doc/faq#Pointers)
- Know when to use value vs pointer semantics
- Always check for [nil](https://go.dev/ref/spec#The_zero_value) when working with pointers
