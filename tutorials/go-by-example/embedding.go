package main

import "fmt"

type base struct {
	num int
}

// Go supports embedding of structs and interfaces
// to express seamless COMPOSITION of types (vs inheritance)
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// Container embeds a base
// Embedding looks like a field without a name
type container struct {
	base
	str string
}

func main() {

	// When creating structs with literals,
	//we have to initialize the embedding explicitly;
	// here the embedded type serves as the field name
	co := container{
		base: base{num: 1}, str: "some name",
	}

	// Base field can be directly accessed
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Base field can also be accessed by full path
	fmt.Println("also num", co.base.num)

	// Since container embeds base, the methods of base also become methods of a container.
	// Here we invoke a method that was embedded from base directly on co.
	// Using base method on container
	fmt.Println("describe:", co.describe())

	// Embedding structs with methods
	// may be used to bestow
	// interface implementations onto other structs
	type describer interface {
		// Adding method describe
		// Which gives interface implementation
		// for base to its implementors
		describe() string
	}

	// DOUBT TODO
	var d describer = co
	fmt.Println("describer:", d.describe())
}
