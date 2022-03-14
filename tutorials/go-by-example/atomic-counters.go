package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
The primary mechanism for managing state in Go is communication over channels.
 We saw this for example with worker pools.
 There are a few other options for managing state though. Here we’ll look
 at using the sync/atomic package for atomic counters accessed by multiple goroutines.
*/

func main() {
	// Unsigned integer to represent counter
	var ops uint64

	// Waitgroup to help wait till all goroutines have finished work
	var wg sync.WaitGroup

	// 50 goroutines incrementing counter 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// Increment using memory address &
				// and atomic.AddUint
				atomic.AddUint64(&ops, 1)

				/*
					Running like this gives incosistency
					because of different goroutines
					interfering with each other
				*/
				// r := &ops
				// *r++

				/*
					We expect to get exactly 50,000 operations.
					 Had we used the non-atomic ops++ to increment the counter,
					 we’d likely get a different number, changing between runs
					 , because the goroutines would interfere with each other.
					Moreover, we’d get data race failures when running with the -race flag.
				*/
			}
			wg.Done()
		}()
	}
	// Wait till waitgroup goes back to 0
	wg.Wait()

	// Access ops after all the operations have been performed
	fmt.Println("ops:", ops)
	/*
		It’s safe to access ops now because we know no other goroutine is writing to it.
		 Reading atomics safely
		while they are being updated is also possible, using functions like atomic.LoadUint64.
	*/

}
