// Stateful Goroutines
//
// In the previous example we used explicit locking with mutexes to synchronize
// access to shared state across multiple goroutines. Another option is to use the
// built-in synchronization features of goroutines and channels to achieve the same
// result. This channel-based approach aligns with Go's ideas of sharing memory by
// communicating and having each piece of data owned by exactly one goroutine.
//
// Run: go run ./49_stateful_goroutines
package main

import (
	"fmt"
	"math/rand/v2"
	"sync/atomic"
	"time"
)

// In this example our state will be owned by a single goroutine. This will guarantee
// that the data is never corrupted with concurrent access. In order to read or write
// that state, other goroutines will send messages to the owning goroutine and
// receive corresponding replies. These readOp and writeOp structs encapsulate those
// requests and a way for the owning goroutine to respond.
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
	// As before we'll count how many operations we perform.
	var readOps atomic.Uint64
	var writeOps atomic.Uint64

	// The reads and writes channels will be used by other goroutines to issue read
	// and write requests, respectively.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Here is the goroutine that owns the state, which is a map as in the previous
	// example but now private to the stateful goroutine. This goroutine repeatedly
	// selects on the reads and writes channels, responding to requests as they
	// arrive.
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// This starts 100 goroutines to issue reads to the state-owning goroutine via
	// the reads channel. Each read requires constructing a readOp, sending it over
	// the reads channel, and then receiving the result over the provided resp
	// channel.
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.IntN(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				readOps.Add(1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We start 10 writes as well, using a similar approach.
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.IntN(5),
					val:  rand.IntN(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				writeOps.Add(1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the goroutines work for a second.
	time.Sleep(time.Second)

	// Finally, capture and report the op counts.
	fmt.Println("many reads:", readOps.Load() > 1000)
	fmt.Println("many writes:", writeOps.Load() > 100)

	// For this particular case the goroutine-based approach was a bit more involved
	// than the mutex-based one. It might be useful in certain cases though, for
	// example where you have other channels involved or when managing multiple such
	// mutexes would be error-prone. Use whichever approach feels most natural.
}
