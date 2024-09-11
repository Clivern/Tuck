// Goroutines
//
// A goroutine is a lightweight thread of execution, managed by the Go runtime.
// Goroutines are cheap enough to start thousands of them.
//
// Run: go run ./32_goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Suppose we have a function call f(s). Here's how we'd call that in the usual
	// way, running it synchronously.
	f("direct")

	// To invoke this function in a goroutine, use go f(s). This new goroutine will
	// execute concurrently with the calling one.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		f("goroutine")
	}()

	// You can also start a goroutine for an anonymous function call.
	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now.
	// When main returns, the program exits without waiting for them, so we wait
	// with a WaitGroup (see WaitGroups). A time.Sleep would be a fragile way to wait.
	wg.Wait()
	fmt.Println("done")

	// When we run this program, we see the output of the blocking call first, then
	// the output of the two goroutines. Their output may be interleaved, because
	// goroutines run concurrently.
	start := time.Now()
	for range 10_000 {
		wg.Go(func() { time.Sleep(10 * time.Millisecond) })
	}
	wg.Wait()
	fmt.Println("10,000 sleeping goroutines took well under a second:", time.Since(start) < time.Second)
}
