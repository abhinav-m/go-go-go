package main

import "fmt"

// Using channels as function parameters
// we can specify if channel is only meant to send or recieve values.
// This specificity  increases tyhe type-safety of the program

// ping function accepts a channel for SENDING values
// It would be a compile time error to try to recieve on this
// channel
// arrow towards channel signifies sending MESSAGES to a channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Arrow from channel signifies recieving message from channel

/*
The pong function accepts one channel for receives (pings) and a second for sends (pongs).
*/
func pong(pings <-chan string, pongs chan<- string) {
	// Recieves message from sender channel pings
	msg := <-pings
	// Passes message to reciever channel pongs
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
