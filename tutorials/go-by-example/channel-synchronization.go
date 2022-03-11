package main

import (
	"fmt"
	"time"
)

/*
Chennels can be used
to synchronize execution
across goroutines.

Waitgroup is an alternative which
can be used when waiting for
multiple goroutines to finish.
*/

// Done channel will be used to notify
// another goroutine
// When routine is done
func worker(done chan bool) {
	fmt.Print("working...")

	time.Sleep(time.Second)

	fmt.Println("done")
	// Send a value to notify that we are done
	// to channel
	done <- true
}

func main() {

	// Start a worker goroutine, giving it channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// BLOCKS until we receive a notification from the worker
	// on the channel.

	// If you removed the <- done line from this program, the
	// program exits before the worker even started
	<-done
}
