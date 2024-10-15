// Packages and Modules
//
// Go code is organized in packages, and packages in modules. A package is a
// directory of .go files that share a package clause. A module is a tree of
// packages with a go.mod file at its root, which names the module path and lists
// its dependencies. This repository is the module github.com/clivern/tuck.
//
// Run: go run ./90_packages_and_modules
// Run: go list ./90_packages_and_modules/...
// Run: head -3 go.mod
package main

// A package is imported by its full path: the module path plus the directory. The
// geometry package lives in the geometry/ folder next to this file. An alias, geo,
// is the name used in this file. A blank import runs init and gives no name to call.
import (
	"fmt"

	geo "github.com/clivern/tuck/90_packages_and_modules/geometry"
	_ "github.com/clivern/tuck/90_packages_and_modules/notice"
)

func main() {
	// Exported names are used through the name chosen at the import, here geo.
	c := geo.Circle{Radius: 2}
	fmt.Println(c.Area(), c.Diameter())

	// Package-level variables are initialized, and init functions run, before main.
	// Steps lists them: geometry's own inits, in source order, then the package
	// that was blank-imported.
	fmt.Println(geo.Default)
	fmt.Println("init:", geo.Steps)

	// Lower-case names are private to their package.
	// geometry.round(1.234)
	// error: name round not exported by package geometry

	// Third-party packages are added with go get, which records them in go.mod and
	// go.sum, for example: go get golang.org/x/sync/errgroup. go mod tidy removes
	// dependencies you no longer use.
	fmt.Println("see the geometry folder for the package source")
}
