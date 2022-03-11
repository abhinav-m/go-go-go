package main

import "fmt"

/*
Closing a channel indicates that no more values will be sent on it.
This can be useful to communicate completion to the channel’s receivers.
*/

func main() {
	// Jobs channel will be used to communicate
	// work done from the main() goroutine to a worker
	// goroutine. When no more jobs for the workers exist,
	// job channels will be closed
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Worker goroutine
	go func() {
		// Repeatedly receives from jobs with j, more := <-jobs.
		// In this special 2-value form of receive, the "more" value will
		// be false if jobs has been closed and all values in channel
		// have already been received.
		// We use this to notify done after we have worked on all jobs
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	// Sends 3 jobs to jobs channel, then closes it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
