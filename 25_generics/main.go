// Generics
//
// Go supports generics, also known as type parameters: functions and types that work
// with any type that satisfies a constraint.
//
// Run: go run ./25_generics
package main

import (
	"cmp"
	"fmt"
)

// SlicesIndex takes a slice of any comparable type and an element of that type, and
// returns the index of the first occurrence of v in s, or -1 if not present. The
// comparable constraint means we can compare values of this type with == and !=.
// The type parameter S ~[]E says S is any slice type with elements of type E.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Map works on any input and output element types; any means no constraint at all.
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

// A constraint can list types. ~int also accepts types whose underlying type is int.
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](nums ...T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// cmp.Ordered allows every type that supports < and >.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// A constraint can require methods. fmt.Stringer is any type with String.
func Quote[T fmt.Stringer](v T) string {
	return fmt.Sprintf("%q", v.String())
}

type label string

func (l label) String() string { return string(l) }

// Zero has no argument the compiler can infer from, so the caller names T.
func Zero[T any]() (z T) { return }

// As an example of a generic type, List is a singly-linked list with values of any
// type.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Methods on generic types are declared like regular ones, but the receiver keeps
// the type parameters: List[T], not List.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements returns all the List elements as a slice. The next example shows a
// more idiomatic way of iterating over custom types.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// When invoking generic functions, we can often rely on type inference. We don't
	// have to specify the types for S and E: the compiler infers them.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// We could also specify them explicitly.
	_ = SlicesIndex[[]string, string](s, "zoo")

	fmt.Println(Map([]int{1, 2, 3}, func(n int) string { return fmt.Sprint(n * n) }))
	fmt.Println(Sum(1, 2, 3), Sum(1.5, 2.5), Max("apple", "pear"))
	fmt.Println("quoted:", Quote(label("go")))
	fmt.Println("zero:", Zero[int]())
	// fmt.Println(Zero())
	// error: in call to Zero, cannot infer T

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
