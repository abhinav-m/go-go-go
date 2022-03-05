package main

import "fmt"

func main() {
	// Creating a map of string and integer
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	// Printing key value accessed above
	fmt.Println("v1", v1)
	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// Optional second value from map
	// First  value is value obtained from map
	// Second value is if the key exists in map
	// This distinguishes empty keys like 0 and ""
	// in map
	ig, prs := m["k3"]

	fmt.Println("prs:", prs)
	// Note ig still logs a value of 0 thus distinguishing empty key
	fmt.Println("_", ig)

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println("map2", n)

}
