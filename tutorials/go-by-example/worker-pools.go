package main

import (
	"fmt"
	"time"
)

/*
	Worker
	which will run several concurrent instances
	These workers will receive work on jobs channel
	and send corresponding results to result

*/

// Jobs channel that outputs jobs results channel that receives result
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		// Simulate time with sleep
		time.Sleep(time.Second)
		fmt.Println("worker", id, "started job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5

	// Two buffered channels for jobs and results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Spawn 3 worker goroutines ( Blocked intially, no jobs yet)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Add more to jobs
	// channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// close jobs ( to indicate all work we have)
	close(jobs)

	// Collect results
	// Also ensures that worker goroutines have finished
	// Alternative way is use waitgroup
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	// Workers execute the 5 jobs concurrently.
	// Program takes about 2.5 seconds for doing 5 seconds of total work
	// Because of 3 workers operating concurrently.
}
