package main

import (
	"fmt"
	"time"
)

// Select allows for waiting on multiple channel oerations
// Combining goroutines and channels with select is a powerful feature
func main() {

	// Defining two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channels will receive a value after
	// some amount of time to simulate 2 different
	// asynch operations eg blocking RPC operations
	// executing in concurrent goroutines
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Select used to await both of these values,
	// printing each as they arrive
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recieved", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	// Note total execution time is only ~2 seconds since
	// both the 1 and 2 seconds sleep execute concurrently

}
