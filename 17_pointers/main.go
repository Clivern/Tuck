// Pointers
//
// Go supports pointers, which let you pass references to values within your program.
// There is no pointer arithmetic, and the garbage collector frees what is no longer
// referenced.
//
// Run: go run ./17_pointers
package main

import (
	"fmt"
	"unsafe"
)

// zeroval has an int parameter, so the argument is passed by value: it gets a copy.
func zeroval(ival int) {
	ival = 0
}

// zeroptr has an *int parameter, a pointer to an int. *iptr dereferences the pointer,
// and assigning to it changes the value at that address.
func zeroptr(iptr *int) {
	*iptr = 0
}

type point struct{ X, Y int }

// Returning the address of a local variable is safe: the compiler moves it to the
// heap.
func newPoint(x, y int) *point {
	p := point{x, y}
	return &p
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i gives the memory address of i, a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("pointer is set:", &i != nil)

	// Fields are reached through a pointer with a plain dot.
	p := newPoint(1, 2)
	p.X = 10
	fmt.Println(*p)

	// new allocates a zeroed value and returns its address. new(expr) also sets an
	// initial value.
	n := new(int)
	m := new(42)
	fmt.Println(*n, *m)

	// The zero value of a pointer is nil, and dereferencing nil panics.
	var q *int
	fmt.Println("nil?", q == nil)

	// unsafe.Pointer converts between pointer types. Ordinary code keeps typed
	// pointers. unsafe describes layout, and talks to memory from outside Go.
	fmt.Println("byte:", unsafe.Sizeof(byte(0)))
	fmt.Println("X then Y:", unsafe.Offsetof(p.X) < unsafe.Offsetof(p.Y))
	fmt.Println("through unsafe:", *(*int)(unsafe.Pointer(&i)))
}
