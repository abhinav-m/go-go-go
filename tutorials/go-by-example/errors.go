package main

import (
	"errors"
	"fmt"
)

/*

In Go it’s idiomatic to communicate errors via an explicit,
 separate return value. This contrasts with the exceptions
 used in languages like Java and Ruby and the overloaded single result
  / error value sometimes used in C. Go’s approach makes it easy to see
  which functions return errors and to handle them using the same
language constructs employed for any other, non-error tasks.
*/

// Idiomatic to return a pair of value and error
// Error is explicit, separate return value
func f1(arg int) (int, error) {

	if arg == 42 {
		//By convention errors are the last return value
		// errors.new creates a new error are  of type error, a built in interface
		return -1, errors.New("can't work with 42")
	}
	// Nil value indicating no error
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// Implementing custom errors is possible
// By implementing the Error() method on them
// (Interface usage as shown in previous tutorials)
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// Using dereferencing &argError syntax to build new struct
		// with supplied values of arg and error
		// and print above implemented interface with it
		return -1, &argError{arg, "can't work with it custom"}
	}

	return arg + 3, nil
}

/*
The two loops below test out each of our error-returning functions.
Note that the use of an inline error check on the if line is a common idiom in Go code.

*/

func main() {
	for _, i := range []int{7, 42} {
		// Use of an inline error check
		// Assigning and checking in same line
		// Common idiom in go.
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)

	// If data from custom error is needed
	// Need to get the error as an instance
	// of the custom error type via type
	// assertion

	// Asserting the error to take out the argument
	// and problem parameters in struct
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
