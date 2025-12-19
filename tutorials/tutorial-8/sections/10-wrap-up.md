# Next Steps & Wrap-up

**Duration:** 2-3 minutes

## Recap What Was Covered:
- Arrays (fixed, value type)
- Slices (dynamic, reference semantics)
- Slice internals (header, backing array)
- Common slice patterns
- Maps (key-value, hash table)
- Map patterns (sets, grouping)
- Performance tips
- In-memory database example

## Homework/Practice Suggestions:
1. **Easy:** Implement a sliding window function
2. **Medium:** Build a word frequency counter
3. **Challenge:** LRU cache with map and linked list
4. **Advanced:** Time-series database with efficient queries

## Cheat Sheet

```
Slice:
  make([]T, len)           Create with length
  make([]T, len, cap)      Create with capacity
  append(s, v)             Append (assign result!)
  copy(dst, src)           Copy elements
  s[low:high]              Slice expression
  s[low:high:max]          Full slice expression

Map:
  make(map[K]V)            Create map
  make(map[K]V, hint)      Create with size hint
  m[k] = v                 Set
  v := m[k]                Get (zero if missing)
  v, ok := m[k]            Get with existence check
  delete(m, k)             Delete
  len(m)                   Size
```
