package mediantwosortedarrays

import (
	"math"
	"sort"
)

// getMedianTwoArrays returns the median of two arrays
func getMedianTwoArrays(arr1 []int, arr2 []int) float64 {
	numbers := append(arr1, arr2...)
	// have to re-sort our array
	sort.Ints(numbers)

	totalNumbers := len(numbers)
	if totalNumbers == 1 {
		return float64(numbers[0])
	}

	if totalNumbers%2 == 0 {
		upperIndex := totalNumbers / 2
		lowerIndex := upperIndex - 1
		upperNumber := numbers[upperIndex]
		lowerNumber := numbers[lowerIndex]
		return float64((upperNumber + lowerNumber) / 2)
	}

	index := math.Floor(float64(totalNumbers / 2))
	return float64(numbers[int8(index)])

}