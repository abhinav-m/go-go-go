package main

import "fmt"

func main() {

	// Contain a specific type of element without number of element
	// Can be thought of as a vector
	// Here we create a slice with length 3
	s := make([]string, 3)

	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// Append values to slice s
	// Need to accept return from append function
	// As there might be a new slice created
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Making a new slice
	c := make([]string, len(s))
	// Copying s to c
	copy(c, s)

	fmt.Println("cpy:", c)
	// Slices support "slice" operator
	// slice low:high (high not included)
	l := s[2:5]

	fmt.Println("sl1:", l)

	fmt.Println("sl2:", s[:5])
	fmt.Println("sl2:", s[2:])

	// Creating and assigning a slice in one line
	t := []string{"g", "h", "i"}

	fmt.Println("dcl:", t)

	// Creating 2D slice with 3 rows
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		// Dynamically creating variable length
		// 2d list
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}
