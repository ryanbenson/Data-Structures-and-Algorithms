package isgoldenratio

import (
	"math"
	"sort"
)

// isGoldenRatio determines if the two points meets the golden ratio or not
func isGoldenRatio(x float64, y float64) bool {
	a := x
	b := y
	if y > x {
		a = y
		b = x
	}
	left := floatRemainder((a + b) / a)
	right := floatRemainder(a / b)
	return left == right
}

// isGoldenRatioArray determines if the two points meets the golden ratio or not
// using an array based algorithm
func isGoldenRatioArray(x float64, y float64) bool {
	numbers := []float64{x, y}
	sort.Float64s(numbers)

	a := numbers[1]
	b := numbers[0]

	left := floatRemainder((a + b) / a)
	right := floatRemainder(a / b)
	return left == right
}

// floatRemainder returns a float with a specific remainder
func floatRemainder(f float64) float64 {
	return math.Round(f*1000)/1000
}
