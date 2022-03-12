package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function to simulate
// work in separate goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	//Sleep to simulate an expensive task.

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Waitgroup is used to wait for all goroutines launched here to finish
	// If a waitgroup is explicitly passed to functions, it should be done
	// by pointer
	var wg sync.WaitGroup

	// Launch several goroutines, increasing waitgroup counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// Avoid using same i value in each goroutine closure ??
		// Sort of like let, giving its own copy of i?
		i := i
		// Wrap worker call in a closure
		// making sure to tell the waitgroup that the worker is done
		// This way worker doesn't have to know about concurrency
		// primitives involved in execution
		go func() {

			defer wg.Done()
			worker(i)
		}()
	}
	/*
		Block until the waitgroup counter goes back to 0;
		All workers are notified they're done
	*/

	// Note this approach has no straightforward way to
	// propogate errors from workers. For more advanced use cases, consider using errgroup
	wg.Wait()
}
