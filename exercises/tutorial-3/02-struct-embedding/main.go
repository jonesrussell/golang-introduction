package main

import "fmt"

// Exercise 2: Struct Embedding
//
// Your task:
// 1. Define a User struct with: ID (int), Username, Email (strings)
// 2. Define an Admin struct that EMBEDS User (anonymous field)
//    - Add Permissions field ([]string)
// 3. Create an Admin and access User fields directly (promoted)
//
// Key difference from composition:
// - Composition: admin.User.Username (explicit)
// - Embedding: admin.Username (promoted)
//
// Expected output:
//   Admin ID: 1
//   Username: admin
//   Email: admin@example.com
//   Permissions: [read write delete]
//   Via User field: admin
//
// Run with: go run main.go

// TODO: Define User struct

// TODO: Define Admin struct that EMBEDS User (no field name!)
// type Admin struct {
//     User              // <-- This is embedding (anonymous field)
//     Permissions []string
// }

func main() {
	// TODO: Create an Admin
	// admin := Admin{
	//     User: User{
	//         ID:       1,
	//         Username: "admin",
	//         Email:    "admin@example.com",
	//     },
	//     Permissions: []string{"read", "write", "delete"},
	// }

	// Uncomment when ready:
	// Fields are promoted - access directly!
	// fmt.Println("Admin ID:", admin.ID)
	// fmt.Println("Username:", admin.Username)
	// fmt.Println("Email:", admin.Email)
	// fmt.Println("Permissions:", admin.Permissions)
	//
	// Can still access via type name
	// fmt.Println("Via User field:", admin.User.Username)

	_ = fmt.Println
}
