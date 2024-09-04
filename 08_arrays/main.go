// Arrays
//
// An array is a numbered sequence of elements with a fixed length. The length is part
// of the type, so [5]int and [4]int are different types. In typical Go code, slices
// are much more common; arrays are useful in some special cases.
//
// Run: go run ./08_arrays
package main

import "fmt"

func main() {
	// This array holds exactly 5 ints. Elements start at the zero value, 0.
	var a [5]int
	fmt.Println("emp:", a)

	// Set and get a value with array[index].
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len returns the length of an array.
	fmt.Println("len:", len(a))

	// Declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// ... lets the compiler count the elements.
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// An index with : sets that element; the ones in between are zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Arrays are values: assigning one copies all of its elements.
	c := b
	c[0] = 0
	fmt.Println("copy:", b[0], c[0])

	// Arrays of the same type can be compared with ==.
	fmt.Println("equal:", [2]int{1, 2} == [2]int{1, 2})

	// Array types are one-dimensional, but you can compose them into
	// multi-dimensional structures.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
