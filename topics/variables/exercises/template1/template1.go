// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/d2M0Q3mRnd

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var a int
	var b bool
	var c string

	// Display the value of those variables.
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	v1 := 15
	v2 := false
	v3 := "Whatever"

	// Display the value of those variables.
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", pi, pi)
}
