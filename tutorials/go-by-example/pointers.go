package main

import "fmt"

// Passed by value
// new copy
func zeroval(ival int) {
	ival = 0
}

// Takes in an int pointer as value

func zeroptr(iptr *int) {
	// The pointer argument iptr
	// then dereferences the pointer
	// from its memory address to the current value at
	// that address using star operator
	// Assigning a value to dereferenced variable
	// Changes value at referenced address
	*iptr = 0
}

func main() {
	i := 1

	fmt.Println("initial:", i)

	// Passed by value
	zeroval(i)
	fmt.Println("zeroval", i)

	// Passing memory address of the variable i
	// using &
	zeroptr(&i)

	fmt.Println("zeroptr:", i)
	// Print memory address
	fmt.Println("pointer:", &i)
}
