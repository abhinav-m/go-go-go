package main

import (
	"fmt"
	"sort"
)

/*
In order to sort by a custom function in Go, we need a corresponding type.
 Here we’ve created a byLength type that is just an alias for the builtin []string type.
*/

// Need a corresponding type to sort by custom function
// byLength is an alieas for builtin []string type
type byLength []string

/*
We implement sort.Interface - Len, Less, and Swap -
on our type so we can use the sort package’s generic Sort function.
Len and Swap will usually be similar across types and Less will hold the actual custom sorting logic.
*/

// Interface sort implementation (len function)
func (s byLength) Len() int {
	return len(s)
}

// Interface sort implementaiton swap function
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Interface sort Less (does the actual sorting)
func (s byLength) Less(i, j int) bool {
	// Reverse to sort in descending
	return len(s[i]) < len(s[j])
}

func main() {

	// Slice of strings
	fruits := []string{"peach", "banana", "kiwi"}
	// Convert to type byLength and sort by sort.Sort
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	/*
		By following this same pattern of creating a custom type,
		 implementing the three Interface methods on that type, and then calling sort.
		Sort on a collection of that custom type, we can sort Go slices by arbitrary functions.
	*/
}
