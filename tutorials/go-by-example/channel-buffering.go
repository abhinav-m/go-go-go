package main

import "fmt"

/*
By default channels are unbuffered,
meaning they only accept sends (chan <-) if there is a corresponding recieve (<- chan)
*/
func main() {
	// Here we make a channel of strings buffering up to 2 values.

	messages := make(chan string, 2)

	// Sending two messages to buffered channel
	// Arrow pointer IMPORTANT SIGNIFIES recieving and sending
	// Because it is a buffered channel we can send values into channel
	// without corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive those two values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
