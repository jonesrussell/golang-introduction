# Practical Example: Building a Linked List

**Duration:** 8-10 minutes

## Build Together

A singly linked list demonstrating pointer concepts.

```go runnable
package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
    Value int
    Next  *Node  // Pointer to next node (or nil for end)
}

// LinkedList represents the entire list
type LinkedList struct {
    Head *Node
    Size int
}

// NewLinkedList creates an empty linked list
func NewLinkedList() *LinkedList {
    return &LinkedList{
        Head: nil,
        Size: 0,
    }
}

// Append adds a value to the end of the list
func (ll *LinkedList) Append(value int) {
    newNode := &Node{Value: value, Next: nil}

    if ll.Head == nil {
        // Empty list - new node becomes head
        ll.Head = newNode
    } else {
        // Traverse to end
        current := ll.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    ll.Size++
}

// Prepend adds a value to the beginning of the list
func (ll *LinkedList) Prepend(value int) {
    newNode := &Node{
        Value: value,
        Next:  ll.Head,  // Point to current head
    }
    ll.Head = newNode  // New node becomes head
    ll.Size++
}

// Get returns the value at the given index
func (ll *LinkedList) Get(index int) (int, bool) {
    if index < 0 || index >= ll.Size {
        return 0, false
    }

    current := ll.Head
    for i := 0; i < index; i++ {
        current = current.Next
    }
    return current.Value, true
}

// Remove removes the first occurrence of value
func (ll *LinkedList) Remove(value int) bool {
    if ll.Head == nil {
        return false
    }

    // Special case: removing head
    if ll.Head.Value == value {
        ll.Head = ll.Head.Next
        ll.Size--
        return true
    }

    // Find node before the one to remove
    current := ll.Head
    for current.Next != nil {
        if current.Next.Value == value {
            current.Next = current.Next.Next  // Skip the node
            ll.Size--
            return true
        }
        current = current.Next
    }

    return false  // Value not found
}

// Contains checks if value exists in list
func (ll *LinkedList) Contains(value int) bool {
    current := ll.Head
    for current != nil {
        if current.Value == value {
            return true
        }
        current = current.Next
    }
    return false
}

// ToSlice converts the list to a slice
func (ll *LinkedList) ToSlice() []int {
    result := make([]int, 0, ll.Size)
    current := ll.Head
    for current != nil {
        result = append(result, current.Value)
        current = current.Next
    }
    return result
}

// String returns a string representation
func (ll *LinkedList) String() string {
    if ll.Head == nil {
        return "[]"
    }

    result := "["
    current := ll.Head
    for current != nil {
        result += fmt.Sprintf("%d", current.Value)
        if current.Next != nil {
            result += " -> "
        }
        current = current.Next
    }
    return result + "]"
}

// Reverse reverses the list in place
func (ll *LinkedList) Reverse() {
    var prev *Node = nil
    current := ll.Head

    for current != nil {
        next := current.Next  // Save next
        current.Next = prev   // Reverse pointer
        prev = current        // Move prev forward
        current = next        // Move current forward
    }

    ll.Head = prev
}

func main() {
    fmt.Println("=== Linked List Demo ===\n")

    // Create list
    list := NewLinkedList()

    // Append values
    list.Append(10)
    list.Append(20)
    list.Append(30)
    fmt.Println("After appending 10, 20, 30:")
    fmt.Println(list)  // [10 -> 20 -> 30]

    // Prepend
    list.Prepend(5)
    fmt.Println("\nAfter prepending 5:")
    fmt.Println(list)  // [5 -> 10 -> 20 -> 30]

    // Get value
    if val, ok := list.Get(2); ok {
        fmt.Printf("\nValue at index 2: %d\n", val)  // 20
    }

    // Contains
    fmt.Printf("Contains 20: %v\n", list.Contains(20))  // true
    fmt.Printf("Contains 99: %v\n", list.Contains(99))  // false

    // Size
    fmt.Printf("Size: %d\n", list.Size)  // 4

    // Remove
    list.Remove(20)
    fmt.Println("\nAfter removing 20:")
    fmt.Println(list)  // [5 -> 10 -> 30]

    // Reverse
    list.Reverse()
    fmt.Println("\nAfter reversing:")
    fmt.Println(list)  // [30 -> 10 -> 5]

    // Convert to slice
    slice := list.ToSlice()
    fmt.Printf("\nAs slice: %v\n", slice)  // [30 10 5]
}
```

## Walk Through:
- Node struct uses pointer to next node
- LinkedList tracks head pointer and size
- Traversal uses pointer following
- Nil represents end of list
- All methods use pointer receivers (modify state)
- Reverse demonstrates pointer manipulation
