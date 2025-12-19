# Common Pitfalls

**Duration:** 4-5 minutes

## PITFALL 1: Modifying [slice](https://go.dev/ref/spec#Slice_types) during iteration

```go
s := []int{1, 2, 3, 4, 5}
for i, v := range s {
    if v%2 == 0 {
        s = append(s[:i], s[i+1:]...)  // BAD: modifies slice!
    }
}
// Fix: iterate backwards or use filter function
```

## PITFALL 2: Append doesn't always create new array

```go
s1 := []int{1, 2, 3, 4, 5}
s2 := s1[:3]
s2 = append(s2, 100)
fmt.Println(s1)  // [1 2 3 100 5] - s1 modified!

// Fix: Use full slice expression
s2 := s1[:3:3]  // Limit capacity
s2 = append(s2, 100)  // Forces new array
```

## PITFALL 3: nil slice vs empty slice

```go
var nilSlice []int
emptySlice := []int{}
fmt.Println(nilSlice == nil)   // true
fmt.Println(emptySlice == nil) // false
// JSON: nil -> null, empty -> []
```

## PITFALL 4: [Map](https://go.dev/ref/spec#Map_types) iteration during modification

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
for k := range m {
    if k == "a" {
        delete(m, "b")  // Undefined behavior!
    }
}
// Fix: Collect keys first, then modify
```

## PITFALL 5: Range loop variable capture

```go
var funcs []func()
for _, v := range []int{1, 2, 3} {
    funcs = append(funcs, func() {
        fmt.Println(v)  // All print 3!
    })
}
// Fix: Shadow variable or pass as argument
for _, v := range []int{1, 2, 3} {
    v := v  // Shadow
    funcs = append(funcs, func() {
        fmt.Println(v)  // Works correctly
    })
}
```

## PITFALL 6: Uninitialized nested map

```go
m := make(map[string]map[string]int)
m["outer"]["inner"] = 1  // PANIC: nil map!
// Fix: Initialize inner map first
if m["outer"] == nil {
    m["outer"] = make(map[string]int)
}
m["outer"]["inner"] = 1
```
