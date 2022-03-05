package main

import "fmt"

func main() {
	// Both assignment and condition with if
	// No need for brackets
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// THERE IS NO TERNARY IN GO
}
