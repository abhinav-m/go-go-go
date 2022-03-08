package main

import "fmt"

// Collection of fields
// Useful for grouping data together
type person struct {
	name string
	age  int
}

// Constructs a new person struct
// with given name
func newPerson(name string) *person {
	p := person{name: name}

	p.age = 28

	// You can safely return
	// pointer to a local variable
	// as a local variable can survive
	// scope of a function
	return &p
}

func main() {
	// Creating a new struct
	fmt.Println(person{"bob", 20})

	// Naming fields while creating struct
	fmt.Println(person{name: "Navneet", age: 29})

	// Neglected fields while initializing are zero valued
	fmt.Println(person{name: "Fred"})

	// & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 30})

	// Idiomatic to encapsulate new struct creation
	// in constructor functions
	fmt.Println(newPerson("John"))

	// Fields are accessed with dots .
	s := person{name: "Sean", age: 50}

	fmt.Println(s.name)

	// Struct pointers
	// You can also use dots with struct pointers -
	// the pointers are automatically dereferenced.
	sp := &s
	// Dereferenced to get age
	fmt.Println(sp.age)

	// Structs are mutable
	// Setting age
	sp.age = 51

	fmt.Println(sp.age)

}
