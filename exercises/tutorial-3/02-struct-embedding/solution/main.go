package main

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
}

// Admin embeds User - no field name (anonymous field)
type Admin struct {
	User        // Embedded - fields are promoted
	Permissions []string
}

func main() {
	admin := Admin{
		User: User{
			ID:       1,
			Username: "admin",
			Email:    "admin@example.com",
		},
		Permissions: []string{"read", "write", "delete"},
	}

	// Fields are promoted - access directly!
	fmt.Println("Admin ID:", admin.ID)
	fmt.Println("Username:", admin.Username)
	fmt.Println("Email:", admin.Email)
	fmt.Println("Permissions:", admin.Permissions)

	// Can still access via type name
	fmt.Println("Via User field:", admin.User.Username)
}
