package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Value receiver - only reads data, doesn't modify
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Value receiver - only reads data
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Value receiver - returns formatted string
func (r Rectangle) String() string {
	return fmt.Sprintf("%.2f x %.2f", r.Width, r.Height)
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Println("Rectangle:", rect.String())
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
}
