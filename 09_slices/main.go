// Slices
//
// Slices are the everyday sequence type in Go. A slice is a view onto an underlying
// array: it has a length, a capacity, and can grow with append. The slices package
// adds many useful functions for working with them.
//
// Run: go run ./09_slices
package main

import (
	"fmt"
	"slices"
)

func main() {
	// Unlike arrays, a slice type has no length. An uninitialized slice is nil and
	// has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// make creates a slice with a length and an optional capacity.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Set and get just like arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// append returns a slice with new values added. It may allocate a new backing
	// array, so always use its result.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy copies elements from one slice into another.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice[low:high] takes a sub-slice from low up to, but excluding, high.
	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	// A sub-slice shares memory with the original: changing one changes the other.
	l[0] = "C"
	fmt.Println("shared:", s)

	// Declare and initialize in one line.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// The slices package has helpers such as Equal, Contains, Index, and Sort.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	fmt.Println(slices.Contains(t, "h"), slices.Index(t, "i"))
	nums := []int{3, 1, 2}
	slices.Sort(nums)
	fmt.Println("sorted:", nums, "max:", slices.Max(nums))

	// Slices can be composed into multi-dimensional structures. Inner slices can
	// have different lengths.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
