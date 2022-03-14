package main

import (
	"fmt"
	"sync"
)

/*
Container holds a map of counters; since we want to update it concurrently
from multiple goroutines, we add a Mutex to synchronize access.
IMPORTANT: Mutexes MUST not be copied (refer to one only)
so, if this struct is passed around, it should be done by pointer
*/
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Increases a containers
// counter by locking its mutex

func (c *Container) inc(name string) {
	// Lock the mutex
	// before accessing counters; unlock
	//it at the end of the function using a defer statement.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// Container with counter a, b intiialized to 0
	// Note that the zero value of a mutex is usable as-is, so no initialization is required here.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	// Waitgroup
	var wg sync.WaitGroup

	// Function to run as goroutine
	doIncrement := func(name string, n int) {
		//Named counter increased n times
		for i := 0; i < n; i++ {
			// Increment method implemented
			// on container
			c.inc(name)
		}
		// Finish current waitgroup
		wg.Done()
	}
	// Add 3 to Waitgroup
	wg.Add(3)

	// Implement concurrent goroutines
	// ON SAME container, 2 of them on same
	// key in container dictionary / map
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	// Wait for goroutines in Waitgroup to finish
	wg.Wait()
	// Print the final values of counters
	fmt.Println(c.counters)
}
