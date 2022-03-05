package main

import "fmt"

func main() {
	// Inferred to string
	var a = "initial"
	fmt.Println(a)
	// Declared and set as integers
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Unitialized variables are zero valued
	var e int
	fmt.Println(e)

	// Initialized to empty string
	var t string
	fmt.Println(t)

	// Simultaneously declaring and assigning variable

	f := "apple"
	fmt.Println(f)

}
