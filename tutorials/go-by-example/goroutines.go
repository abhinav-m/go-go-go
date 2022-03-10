package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// A goroutine is a lightweight thread of execution
func main() {

	// Direct call of  function , synchronous
	f("direct")

	// Invoking function with goroutine
	// using go keyword. Executes concurrently
	// with calling one
	go f("goroutine")

	// Calling anonymous function as a goroutine
	// sort of like IIFE
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	/*
		When we run this program, we see the output of the blocking call first,
		then the output of the two goroutines. The goroutines’ output may be interleaved,
		 because goroutines are being run concurrently by the Go runtime.

	*/

	// Two functions are running asynchronously
	// in separate goroutines, Wait for them to finish
	// Using sleep (NOT A ROBUST APPROACH)
	time.Sleep(time.Second)
	fmt.Println("done")
}
