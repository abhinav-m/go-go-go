package main

import (
	"fmt"
	"time"
)

/*

Rate limiting is an important mechanism for controlling resource utilization
 and maintaining quality of service.
Go elegantly supports rate limiting with goroutines, channels, and tickers.

*/
func main() {
	// Dummy requests channel with 5 requests
	requests := make(chan int, 5)

	// Pass requests to channel with same name
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	// Close channel to indicate everything sent to channel
	close(requests)

	// Limiter channel recieves a value every 200 ms
	// Regulator in limiting scheme
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on recieve via limiter channel,
	// We achieve limiting requests to 1 request every 200 ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
		/*
			Running our program we see
			the first batch of requests
			handled once every ~200 milliseconds as desired.
		*/
	}

	// Allowing short bursts of requests in our rate limiting
	// scheme while preserving the overall rate limit.
	// We can accomplish this by buffering our limiter channel
	// Burstylimiter channel will allow bursts of upto 3 events
	burstyLimiter := make(chan time.Time, 3)

	// Filling channel to represent allowed bursts
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 ms, try to add a new value to burstyLimiter
	// Upto its limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			// Send value to bursty limiter
			burstyLimiter <- t
		}
	}()

	// Simulate 5 more requests
	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		// First three will benifit from the burst capability
		// of burstyLimiter
		burstyRequests <- i
	}

	close(burstyRequests)

	/*
		For the second batch of requests we serve the first 3 immediately because
		 of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.
	*/
	for req := range burstyRequests {
		// Wait for burstyLimiter Only first will benifit from 3 concurrent limits
		// Second round will delay each request by 200 ms
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
