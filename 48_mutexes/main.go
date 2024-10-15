// Mutexes
//
// In the previous example we saw how to manage simple counter state using atomic
// operations. For more complex state we can use a mutex to safely access data across
// multiple goroutines.
//
// Run: go run ./48_mutexes
package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; since we want to update it concurrently from
// multiple goroutines, we add a Mutex to synchronize access. Note that mutexes must
// not be copied, so if this struct is passed around, it should be done by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Lock the mutex before accessing counters; unlock it at the end of the function
// using a defer statement.
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// An RWMutex allows many readers at once, or a single writer.
type Config struct {
	mu   sync.RWMutex
	vals map[string]string
}

func (c *Config) Get(k string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.vals[k]
}

func main() {
	// Note that the zero value of a mutex is usable as-is, so no initialization is
	// required here.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// This function increments a named counter in a loop.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// Run several goroutines concurrently; note that they all access the same
	// Container, and two of them access the same counter.
	wg.Go(func() { doIncrement("a", 10000) })
	wg.Go(func() { doIncrement("a", 10000) })
	wg.Go(func() { doIncrement("b", 10000) })

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println(c.counters)

	cfg := &Config{vals: map[string]string{"env": "prod"}}
	fmt.Println(cfg.Get("env"))

	// sync.Map is a map that is safe for concurrent use without an extra mutex.
	// A plain map plus a mutex is still the usual choice when the lock can cover
	// a larger update.
	var scores sync.Map
	scores.Store("a", 1)
	scores.Store("a", 2)
	v, ok := scores.Load("a")
	fmt.Println("sync map:", v, ok)

	// Cond waits for a condition protected by a mutex. Wait unlocks while the
	// goroutine sleeps and locks again before it returns.
	var mu sync.Mutex
	var queue []int
	cond := sync.NewCond(&mu)
	mu.Lock()
	go func() {
		mu.Lock()
		queue = append(queue, 1)
		cond.Signal()
		mu.Unlock()
	}()
	for len(queue) == 0 {
		cond.Wait()
	}
	fmt.Println("cond:", queue[0])
	mu.Unlock()
}
