// Range over Iterators
//
// range works with iterator functions too. An iterator is a function that takes a
// yield callback and calls it once per element. iter.Seq[V] is the standard name
// for the type func(yield func(V) bool).
//
// Run: go run ./26_range_over_iterators
package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All returns an iterator. The iterator function calls yield for every element we
// want to iterate over, and stops early if yield returns false.
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iteration doesn't need an underlying data structure, and doesn't even have to be
// finite. Here's a function returning an iterator over Fibonacci numbers: it keeps
// running as long as yield keeps returning true.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Since List.All returns an iterator, we can use it in a regular range loop.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Packages like slices have many useful functions that work with iterators.
	// Collect gathers all values into a slice.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// Once the loop hits break or an early return, yield returns false and the
	// iterator stops.
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}

	// iter.Seq2 yields pairs. slices.All yields index and value; maps.Keys and
	// slices.Sorted combine into sorted keys.
	for i, v := range slices.All([]string{"a", "b"}) {
		fmt.Println(i, v)
	}
	ages := map[string]int{"bo": 2, "ada": 1}
	fmt.Println(slices.Sorted(maps.Keys(ages)))

	// The strings package returns iterators as well.
	for part := range strings.SplitSeq("x,y,z", ",") {
		fmt.Print(part, " ")
	}
	fmt.Println()
}
