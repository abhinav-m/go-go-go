package main

import (
	"fmt"
	"time"
)

/*
Timeouts are important for programs that connect to external resources or
 that otherwise need to bound execution time. (Asynchronous?)
Implementing timeouts in Go is easy and elegant thanks to channels and select.


*/
func main() {
	// Buffered channel
	// So send in goroutine is NON BLOCKING
	// Common pattern to prevent goroutine leaks in case
	// channel is never read
	c1 := make(chan string, 1)

	// Simulating async using goroutine
	// by sleeping 2s ( assume this is an external call)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Select implementing a timeout.
	// We await
	// res := <-c1 awaits a result and <-time.After
	// awaits a value to be sent after the timeout of 1s
	// Since select proceeds with the first receive thats ready
	// timeout case is accepted if the operation takes more
	// than allowed 1s.
	select {
	case res := <-c1:
		fmt.Println(res)
	// Accepts timeout 1
	// IF operation doesn't succeed within 1 second
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	/*
		If we allow a longer timeout of 3s,
		then the receive from c2 will succeed and we’ll print the result.
	*/
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

//Running this program shows the first operation timing out and the second succeeding
