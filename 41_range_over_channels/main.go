// Range over Channels
//
// In an earlier example we saw how for and range provide iteration over basic data
// structures. We can also use this syntax to iterate over values received from a
// channel.
//
// Run: go run ./41_range_over_channels
package main

import "fmt"

// gen returns a channel and a goroutine that fills it, closing it when done. This
// generator shape is very common.
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func main() {
	// We'll iterate over 2 values in the queue channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterates over each element as it's received from queue. Because
	// we closed the channel above, the iteration terminates after receiving the 2
	// elements.
	for elem := range queue {
		fmt.Println(elem)
	}

	// This example also showed that it's possible to close a non-empty channel but
	// still have the remaining values be received.
	total := 0
	for n := range gen(1, 2, 3, 4) {
		total += n
	}
	fmt.Println("total:", total)
}
