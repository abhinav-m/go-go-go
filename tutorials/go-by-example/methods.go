package main

import "fmt"

// Methods are supported on struct types

// Struct defining rectangle with width, height
type rect struct {
	width, height int
}

// Methods can be defined by either pointer
// or value reciever types

// Defines a reciever type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// Defines a value reciever

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Creates copy of value
	// since method is of value reciever type
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls
	//  or to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
