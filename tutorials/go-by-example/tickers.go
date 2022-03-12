package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Do something repeatedly at intervals

	*/

	// Tickers use a similar mechanism to timers,
	// a channel that is sent values. Here we will use
	// builtin channel to wait every 500 ms for values
	ticker := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Sleep this thread for 1600 ms
	// so ticker works 3 times
	time.Sleep(1600 * time.Millisecond)
	// Stop ticker
	ticker.Stop()
	// Pass true to channel to stop goroutine above
	done <- true
	fmt.Println("Ticker stopped")
}
