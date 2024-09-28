// Package units formats lengths. It lives under internal/, so only code inside the
// geometry directory tree can import it.
package units

import "fmt"

// Centimeters formats a length in centimeters.
func Centimeters(v float64) string {
	return fmt.Sprintf("%.1fcm", v)
}
