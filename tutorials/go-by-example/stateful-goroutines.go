package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*

In the previous example we used explicit locking with mutexes to synchronize
access to shared state across multiple goroutines. Another option is to use
the built-in synchronization features of goroutines and channels to achieve
the same result. This channel-based approach aligns with Go’s ideas
of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.
*/

/*
In this example our state will be owned by a single goroutine.
This will guarantee that the data is never corrupted with concurrent access

*/

/*
	In order to read or write that state, other goroutines will send messages
	to the owning goroutine and receive corresponding replies. These readOp
	and writeOp structs encapsulate those requests and a way for the owning
     goroutine to respond.
*/
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// Counting operations performed
	var readOps uint64
	var writeOps uint64

	// Channels used to write and write requests
	reads := make(chan readOp)
	writes := make(chan writeOp)

	//  Goroutine that owns the state map
	//  but private to the stateful goroutine.

	go func() {
		var state = make(map[int]int)
		/*
			This goroutine repeatedly selects on the reads and writes channels,
			 responding to requests as they arrive. A response is executed by
			 first performing the requested operation and then sending a value on the response
			channel resp to indicate success (and the desired value in the case of reads)
		*/
		for {
			// Waits for read from reads channel
			select {
			case read := <-reads:
				// Reads and sends the value to the channel in the readOp struct
				read.resp <- state[read.key]

			case write := <-writes:
				// Writes the value received in the stateful go routine
				state[write.key] = write.val
				// Sends true to channel struct indicating success
				write.resp <- true
			}
		}

	}()

	// Runs 100 goroutines
	for r := 0; r < 100; r++ {
		go func() {
			// Constructs a readOp
			// Involving a random key
			// and a random response channel
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				// Sends readop to reads channel
				reads <- read
				// Receiving a result from provided response channel
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Similarly, write 10 times
	for w := 0; w < 10; w++ {
		go func() {
			for {
				// WriteOp type
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				// Send write to writes channel
				writes <- write
				// wait for response
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)

			}
		}()
	}
	//Let goroutines run
	time.Sleep(time.Second)

	//Capture and report final counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	fmt.Println("writeOps:", writeOpsFinal)
}

/*
Running our program shows that the goroutine-based state management example
completes about 98,000 total operations.

go run stateful-goroutines.go
readOps: 88930
writeOps: 8898
For this particular case the goroutine-based approach was a bit more involved than the mutex-based one.
It might be useful in certain cases though, for example where you have other channels involved or
when managing multiple such mutexes would be error-prone. You should use whichever approach
feels most natural, especially with respect to understanding the correctness of your program.

*/
