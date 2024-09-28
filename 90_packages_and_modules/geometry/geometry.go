// Package geometry is a tiny library used by the Packages and Modules example. The
// package name is the last element of its import path.
package geometry

import (
	"math"

	"github.com/clivern/tuck/90_packages_and_modules/geometry/internal/units"
)

// Names starting with an upper case letter are exported: other packages can use
// them.
type Circle struct {
	Radius float64
}

// Area is exported, so callers can use it.
func (c Circle) Area() float64 {
	return round(math.Pi * c.Radius * c.Radius)
}

// Diameter returns the diameter formatted in centimeters.
func (c Circle) Diameter() string {
	return units.Centimeters(2 * c.Radius)
}

// round is lower case, so it is private to this package.
func round(f float64) float64 {
	return math.Round(f*100) / 100
}

// init runs once, when the package is first imported, before main starts.
func init() {
	Default = Circle{Radius: 1}
}

// Default is set by init.
var Default Circle
