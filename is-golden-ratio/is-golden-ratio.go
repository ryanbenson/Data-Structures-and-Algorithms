package isgoldenratio

import "math"

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

func floatRemainder(f float64) float64 {
	return math.Round(f*1000)/1000
}