# Pointers vs Reference Types

**Duration:** 4-5 minutes

## Topics to cover:
- Slices, maps, channels are "reference types"
- What that actually means
- When you still need pointers with these types

## Code Examples

```go
// Slices - header is copied, backing array is shared
func modifySlice(s []int) {
    s[0] = 999  // Modifies original backing array
    s = append(s, 100)  // Does NOT affect original slice header
}

func main() {
    nums := []int{1, 2, 3}
    modifySlice(nums)
    fmt.Println(nums)  // [999 2 3] - element modified, but no 100!
}

// To modify the slice itself (not just contents), use pointer
func modifySlicePtr(s *[]int) {
    *s = append(*s, 100)  // Affects original
}

func main() {
    nums := []int{1, 2, 3}
    modifySlicePtr(&nums)
    fmt.Println(nums)  // [1 2 3 100]
}

// Maps - similar behavior
func addToMap(m map[string]int) {
    m["new"] = 42  // Modifies original map
}

// But can't reassign the map itself without pointer
func replaceMap(m *map[string]int) {
    *m = make(map[string]int)  // Replaces entire map
}

// Summary:
// - Slice/map: pass by value copies the header, shares data
// - Modifying contents: no pointer needed
// - Replacing entire slice/map: pointer needed
// - Appending to slice: pointer needed if caller needs to see new length

// Common pattern: return modified slice
func appendSafe(s []int, v int) []int {
    return append(s, v)  // Caller uses returned slice
}

func main() {
    nums := []int{1, 2, 3}
    nums = appendSafe(nums, 4)  // Reassign
    fmt.Println(nums)  // [1 2 3 4]
}
```

## Key teaching points:
- Slices/maps/channels have internal pointers
- Modifying contents works without `*`
- Replacing the whole thing needs `*` or return value
- Return value pattern is often cleaner than `*[]Type`
- Know the difference: modifying vs replacing
