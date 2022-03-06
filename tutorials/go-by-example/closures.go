package main

import "fmt"

// Go supports ANONYMOUS functions
// that form closures
// Returns another function
func intSeq() func() int {
	// "Closes" over the variable i

	i := 0
	// Inline function forming closure
	return func() int {
		i++
		return i
	}
}

func main() {

	// Assigned a function
	nextInt := intSeq()

	// Stateful, preserves state
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// New state
	newInts := intSeq()
	fmt.Println(newInts())
}
