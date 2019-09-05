package geogo

import "math"

// radians  - converts degrees to radians.
func radians(degree float64) float64 {
	return degree * math.Pi / 180
}
