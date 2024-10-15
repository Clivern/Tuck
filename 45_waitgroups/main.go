// WaitGroups
//
// To wait for multiple goroutines to finish, we can use a wait group.
//
// Run: go run ./45_waitgroups
package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// This is the function we'll run in every goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by
	// pointer.
	var wg sync.WaitGroup

	// Launch several goroutines using WaitGroup.Go, which adds one to the counter,
	// runs the function in a goroutine, and marks it done when the function returns.
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// Block until all the goroutines started by wg are done.
	wg.Wait()

	// The classic form calls Add before starting the goroutine and Done inside it.
	// Add must happen before go, not inside the goroutine, or Wait may return early.
	results := make([]int, 3)
	for i := range results {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = i * i
		}()
	}
	wg.Wait()
	fmt.Println(results)

	// The order of workers starting up and finishing is likely to be different for
	// each invocation.

	// errgroup runs workers and returns the first error one of them reports.
	var g errgroup.Group
	for i := 1; i <= 3; i++ {
		g.Go(func() error {
			if i == 2 {
				return fmt.Errorf("worker %d failed", i)
			}
			return nil
		})
	}
	fmt.Println("errgroup:", g.Wait())
}
