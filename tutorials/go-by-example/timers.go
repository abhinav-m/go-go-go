package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Timers represent a single event in the future.
		You tell the timer how long you want to wait,
		and it provides a channel that will be notified
		at that time. This timer will wait 2 seconds.

	*/
	timer1 := time.NewTimer(2 * time.Second)

	// Blocks on the timer channel, C till timer is fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Alterative to time.Sleep
	// However, timers can be CANCELLED
	timer2 := time.NewTimer(time.Second)

	// Wait for it to fire
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// Stop it earlier
	stop2 := timer2.Stop()

	/*
			Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.

		    time.Sleep(2 * time.Second)
			}
			The first timer will fire ~2s after we start the program,
			 but the second should be stopped before it has a chance to fire


	*/

	// We notice here that timer 2 stops before firing since above
	// code executes first, instead of the goroutine
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
