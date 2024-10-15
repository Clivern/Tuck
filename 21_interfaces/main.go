// Interfaces
//
// Interfaces are named collections of method signatures. A type implements an
// interface simply by having its methods: there is no implements keyword.
//
// Run: go run ./21_interfaces
package main

import (
	"fmt"
	"math"
)

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// We'll implement this interface on rect and circle types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to implement all of its methods.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for circles.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// scale has a pointer receiver, so it can change the circle it is called on.
func (c *circle) scale(k float64) {
	c.radius *= k
}

// scaler is implemented only by types whose method set includes scale.
type scaler interface {
	scale(float64)
}

// A variable with an interface type can call the methods in the interface. Here's a
// generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Sometimes it's useful to know the runtime type of an interface value. A type
// assertion does that; the comma-ok form reports whether it worked.
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

// This line fails to compile if rect ever stops implementing geometry.
// Value methods are in the method set of both rect and *rect.
var _ geometry = rect{}

// A pointer method is only in the method set of the pointer. *circle implements
// scaler. A circle value does not, even though a call on a variable still works.
var _ scaler = (*circle)(nil)

// var _ scaler = circle{}
// error: cannot use circle{} (value of struct type circle) as scaler value in variable declaration: circle does not implement scaler (method scale has pointer receiver)

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The circle and rect types both implement the geometry interface, so we can use
	// instances of them as arguments to measure.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)

	// A slice of an interface type can hold values of different types.
	shapes := []geometry{r, c, rect{1, 1}}
	total := 0.0
	for _, s := range shapes {
		total += s.area()
	}
	fmt.Printf("total area %.2f\n", total)

	// The empty interface, any, is satisfied by every type.
	var anything any = 42
	fmt.Println(anything)
	anything = "now a string"
	fmt.Println(anything)

	// Calling a pointer method on a variable is fine: Go takes its address.
	// That does not put the method in the value's method set, which is why
	// circle itself is not a scaler above.
	grown := circle{radius: 5}
	grown.scale(2)
	fmt.Println("scaled radius:", grown.radius)

	// An interface is nil only when it holds neither a type nor a value.
	var noShape geometry
	fmt.Println("nil interface:", noShape, noShape == nil)

	// A nil pointer stored in an interface is not a nil interface. The type is
	// *circle, so the comparison is false, and calling a method panics.
	var hole *circle
	var asShape geometry = hole
	fmt.Println("typed nil:", asShape, asShape == nil)
	func() {
		defer func() { fmt.Println("nil call panicked:", recover() != nil) }()
		asShape.area()
	}()
}
