// Sorting
//
// Go's slices package implements sorting for builtins and user-defined types. We'll
// look at sorting for builtins first.
//
// Run: go run ./52_sorting
package main

import (
	"fmt"
	"slices"
)

func main() {
	// Sorting functions are generic, and work for any ordered built-in type. For a
	// list of ordered types, see cmp.Ordered.
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting ints.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use the slices package to check if a slice is already in sorted
	// order.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)

	// A sorted slice can be searched quickly with a binary search.
	i, found := slices.BinarySearch(ints, 4)
	fmt.Println("Found 4:", found, "at", i)

	// Reverse, Max, and Min round out the basics.
	slices.Reverse(strs)
	fmt.Println("Reversed:", strs, slices.Max(ints), slices.Min(ints))

	// slices.Sorted collects and sorts the values of an iterator.
	fmt.Println(slices.Sorted(slices.Values([]float64{2.5, -1, 0})))
}
