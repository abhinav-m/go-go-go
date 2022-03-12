package main

import "fmt"

func main() {
	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"

	close(queue)

	// Ranges can iterate over queues
	// It gets each element as it's received from
	// channel. Because we closed the channel above,
	// the iteration terminates after receiving 2 elements
	for elem := range queue {
		fmt.Println(elem)
	}

	/*
		Example also showed that it's possible to close a
		non-empty channel but still have the remaining
		values be received
	*/
}
