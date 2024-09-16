// Sync Once and Pool
//
// The sync package has a few more tools beyond mutexes and wait groups. sync.Once runs
// a function exactly once no matter how many goroutines call it, which suits lazy
// initialization. sync.Pool keeps a set of reusable objects to cut allocations.
//
// Run: go run ./50_sync_once
package main

import (
	"bytes"
	"fmt"
	"sync"
)

// A config that is expensive to load, loaded on first use.
var (
	once   sync.Once
	config map[string]string
)

func loadConfig() map[string]string {
	once.Do(func() {
		fmt.Println("loading config (only once)")
		config = map[string]string{"env": "prod"}
	})
	return config
}

// sync.OnceValue wraps a function so it runs once and caches its result.
var port = sync.OnceValue(func() int {
	fmt.Println("computing port (only once)")
	return 8080
})

// A pool of buffers. New is called when the pool has nothing to hand out.
var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func render(name string) string {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()
	buf.WriteString("hello, ")
	buf.WriteString(name)
	return buf.String()
}

func main() {
	// Many goroutines ask for the config at the same time; it loads once.
	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func() { loadConfig() })
	}
	wg.Wait()
	fmt.Println(loadConfig()["env"])

	fmt.Println(port(), port())

	// Reusing buffers from the pool; the caller never sees the difference.
	fmt.Println(render("gopher"))
	fmt.Println(render("world"))
}
