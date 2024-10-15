// Package notice has no API. The packages example blank-imports it so init runs
// for its side effect.
package notice

import "github.com/clivern/tuck/90_packages_and_modules/geometry"

// init runs because of that blank import. geometry is imported here, so its
// inits have already finished.
func init() {
	geometry.Steps = append(geometry.Steps, "notice")
}
