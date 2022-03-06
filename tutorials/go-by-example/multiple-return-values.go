package main

import "fmt"

// Function signature
// signifying returning of two
// ints
func vals() (int, int) {
	// Sort of like a tuple
	return 3, 7
}

func main() {
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)
	// Using blank identifier
	_, c := vals()
	fmt.Println(c)
}
