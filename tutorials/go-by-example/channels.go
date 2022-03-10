package main

import "fmt"

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and
receive those values into another goroutine.

*/
func main() {
	// Make a channel with make and val-type
	// Channels are typed by the values that they convey
	messages := make(chan string)

	// Send a value to the channel with the <- syntax
	// Here we send a message "ping" to the messages
	// channel we made above, from a new goroutine
	go func() { messages <- "ping " }()

	/*
		The <-channel syntax receives a value from the channel.
		Here we’ll receive the "ping" message we sent above and print it out.
	*/

	// Receiving a value from the messages channel
	// This is in another goroutine thus recieved in different one
	msg := <-messages
	fmt.Println(msg)

	/*

		When we run the program the "ping" message is successfully passed from one
		 goroutine to another via our channel.


		By default sends and receives BLOCK until both the sender and receiver are READY.
	 	This property allowed us to wait at the end of our program
	 	for the "ping" message without having to use any other synchronization.
	*/
}
