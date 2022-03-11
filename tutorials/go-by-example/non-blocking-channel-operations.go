package main

import "fmt"

/*
	Basic sends and receives on channels are blocking.
	However, we can use select with a default clause
	to implement non-blocking sends, receives and even non-
	blocking multi-way selects
*/
func main() {

	messages := make(chan string)
	signals := make(chan bool)

	/*
		Non blocking receive. If a value is available
		on messages, then select will take <- messages case
		with that value. If not, it will take default case immediately
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	/*
		Non blocking send, message can not be sent
		to messages channel, because channel has no
		buffer and no receiver (blocking operation). Therefore, default
		case is selected
	*/
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	/*
		Select multiple cases above default clause
		implementing a multi-way non-blocking select.
		Here we attempt non-blocking receives on both messages
		and signals.
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
