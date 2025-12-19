package main

import "fmt"

type Counter struct {
	Value int
}

// Pointer receiver - modifies the original
func (c *Counter) Increment() {
	c.Value++
}

// Pointer receiver - modifies the original
func (c *Counter) Add(n int) {
	c.Value += n
}

// Pointer receiver - modifies the original
func (c *Counter) Reset() {
	c.Value = 0
}

func main() {
	counter := Counter{}

	fmt.Println("Initial:", counter.Value)

	counter.Increment()
	fmt.Println("After Increment:", counter.Value)

	counter.Add(5)
	fmt.Println("After Add(5):", counter.Value)

	counter.Reset()
	fmt.Println("After Reset:", counter.Value)
}
