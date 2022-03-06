package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	fmt.Println(fact(7))

	// For closures to be recursive
	// must be declared explicitly
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	// Since declared in main, go knows what to call
	fmt.Println(fib(7))
}
