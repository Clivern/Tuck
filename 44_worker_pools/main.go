// Worker Pools
//
// In this example we'll look at how to implement a worker pool using goroutines and
// channels.
//
// Run: go run ./44_worker_pools
package main

import (
	"fmt"
	"slices"
	"time"
)

// Here's the worker, of which we'll run several concurrent instances. These workers
// will receive work on the jobs channel and send the corresponding results on
// results. We'll sleep a little per job to simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// In order to use our pool of workers we need to send them work and collect
	// their results. We make 2 channels for this.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 jobs and then close that channel to indicate that's all the
	// work we have.
	start := time.Now()
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work. This also ensures that the
	// worker goroutines have finished.
	var all []int
	for range numJobs {
		all = append(all, <-results)
	}
	slices.Sort(all)
	fmt.Println("results:", all)

	// Our running program shows the 5 jobs being executed by various workers. The
	// program only takes about 200ms despite doing about 500ms of total work,
	// because there are 3 workers operating concurrently.
	fmt.Println("took about", time.Since(start).Round(100*time.Millisecond))
}
