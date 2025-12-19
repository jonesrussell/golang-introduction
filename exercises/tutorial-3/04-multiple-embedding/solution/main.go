package main

import "fmt"

type Reader struct {
	Name string
}

func (r Reader) Read() string {
	return "Reading data..."
}

type Writer struct {
	Name string
}

func (w Writer) Write() string {
	return "Writing data..."
}

// ReadWriter embeds both Reader and Writer
// Both have Name field - must access via type name or shadow with own field
type ReadWriter struct {
	Name   string // Shadows embedded Name fields
	Reader        // Embedded
	Writer        // Embedded
}

func main() {
	rw := ReadWriter{
		Name:   "rw",
		Reader: Reader{Name: "reader"},
		Writer: Writer{Name: "writer"},
	}

	// ReadWriter's own Name field takes precedence
	fmt.Println("ReadWriter Name:", rw.Name)

	// Can still access embedded Names via type
	fmt.Println("Reader Name:", rw.Reader.Name)
	fmt.Println("Writer Name:", rw.Writer.Name)

	// Methods are still promoted (no conflict - different names)
	fmt.Println(rw.Read())
	fmt.Println(rw.Write())
}
