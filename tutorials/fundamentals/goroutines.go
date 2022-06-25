package main

import (
	"fmt"
	"sync"
	"time"
)

func count(thing string) {
	
	 
 for  i := 0;i <5; i++ {
	fmt.Println(i,thing)
	time.Sleep(time.Millisecond * 800)
	
 }
}

func main() {
	//1. Waits for count function to finish and then moves on
	// This never ends
	// count("sheep")
	// count("cat")

	//1. Runs goroutine 1 and executes both
	// side by side. (Can be multiple goroutines)
	// // This never ends
	// go count("sheep")
	// count("cat")

	//2. Run goroutine 1 and 2 and exit because
	// main goroutine exits (function main) both
	// go count("sheep")
	// go count("cat")

	// Block with scanln to run 2.
	// Or any other blocking code
	// fmt.Scanln()

	// 1. Waitgroup to handle goroutine as its 
	// working 
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		// Terminate one waitgroup 
		wg.Done()
	}()
	//Wait for function to complete
	// Block till goroutine is done
	wg.Wait()
	// go count("sheep")
	// go count("cat")

}