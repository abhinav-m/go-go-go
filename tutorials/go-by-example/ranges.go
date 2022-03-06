package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}

	//range iterates over elements in a variety of data structures

	sum := 0

	for _, num := range nums {
		sum += num
	}

	// Range provides both index and value
	// operating on arrays and slices

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Range on maps iterates over key value pairs

	myMap := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range myMap {
		// Formatting key value pairs
		fmt.Printf("%s -> %s\n", k, v)
	}
	// Iterating only over keys
	for k := range myMap {
		fmt.Println("key:", k)
	}

	// Range on strings iterate over unicode code points
	// First value -> Starting BYTE index of the rune
	// Second value -> rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
