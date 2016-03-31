// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/ItPe2EEy9X

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name, email string
	age         int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a struct literal.
	u1 := user{
		name:  "David",
		email: "DavidVaini@Gmail.com",
		age:   28,
	}
	// Display the field values.

	// Declare a variable using an anonymous struct.
	customer := struct {
		name  string
		email string
		age   int
	}{
		name:  "David",
		email: "DavidVaini@gmail.com",
		age:   28,
	}

	// Display the field values.
	fmt.Println("User:", u1)
	fmt.Println("Person Name", customer.name)
	fmt.Println("Person Email", customer.email)
	fmt.Println("PersonAge", customer.age)
}
