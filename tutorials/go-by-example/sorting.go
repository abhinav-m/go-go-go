package main

import (
	"fmt"
	"sort"
)

/*
Go’s sort package implements
 sorting for builtins and user-defined types.
 We’ll look at sorting for builtins first.
*/
func main() {
	// Sorting ints and strings
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)

	fmt.Println("Ints:   ", ints)

	// Checking if slice is sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
