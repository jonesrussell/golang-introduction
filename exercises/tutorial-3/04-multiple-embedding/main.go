package main

import "fmt"

// Exercise 4: Multiple Embedding
//
// Your task:
// 1. Define a Reader struct with Name string and Read() method
// 2. Define a Writer struct with Name string and Write() method
// 3. Define a ReadWriter struct that embeds BOTH Reader and Writer
// 4. Handle the Name conflict by adding a Name field to ReadWriter
//
// Expected output:
//   ReadWriter Name: rw
//   Reader Name: reader
//   Writer Name: writer
//   Reading data...
//   Writing data...
//
// Run with: go run main.go

// TODO: Define Reader struct with Name field and Read() method

// TODO: Define Writer struct with Name field and Write() method

// TODO: Define ReadWriter that embeds both Reader and Writer
// Note: Both have Name field - this causes a conflict!
// Add explicit Name field to ReadWriter to resolve

func main() {
	// TODO: Create a ReadWriter
	// rw := ReadWriter{
	//     Name: "rw",  // This shadows the embedded Name fields
	//     Reader: Reader{Name: "reader"},
	//     Writer: Writer{Name: "writer"},
	// }

	// Uncomment when ready:
	// fmt.Println("ReadWriter Name:", rw.Name)          // Uses ReadWriter's Name
	// fmt.Println("Reader Name:", rw.Reader.Name)       // Access embedded Name
	// fmt.Println("Writer Name:", rw.Writer.Name)       // Access embedded Name
	// fmt.Println(rw.Read())                            // Promoted method
	// fmt.Println(rw.Write())                           // Promoted method

	_ = fmt.Println
}
