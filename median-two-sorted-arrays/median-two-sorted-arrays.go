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


// getMedianTwoArraysOptimized returns the median of two arrays, but using optimized efficiency
// it avoids merging the two arrays and re-sorting them
func getMedianTwoArraysOptimized(arr1 []int, arr2 []int) float64 {
	// get the target index(es) based on the length of both arrays
	arr1Len := len(arr1)
	arr2Len := len(arr2)
	totalLen := arr1Len + arr2Len
	var targetNumbersCount int
	isEven := false
	
	if totalLen %2 == 0 {
		isEven = true
		targetNumbersCount = totalLen / 2 + 1 // use the upper index, we'll -1 later if it's even. by default it rounds down
	}
	
	curArr1Index := 0
	curArr2Index := 0
	numbersCounted := 0
	
	previousNumber := 0 // used if we have an even number of items
	medianNumber := 0
	
	// loop through our arrays using the smallest ints we can find until we get our median index(es)
	for numbersCounted <= targetNumbersCount {
		val1 := arr1[curArr1Index]
		val2 := arr1[curArr2Index]
		
		if (val1 <= val2) {
			// if we're at the index we want, we have our median
			if numbersCounted == targetNumbersCount {
				medianNumber = val1
			} else {
				// otherwise hold onto the previous number in case it's an even count and move on
				previousNumber = val1
				curArr1Index++
			}
		} else {
			// same as before if val1 was lower, but using val2
			if numbersCounted == targetNumbersCount {
				medianNumber = val2
			} else {
				previousNumber = val2
				curArr2Index++
			}
		}
		numbersCounted++
	}
	
	// if it's an even number, we have our solution. if not divide our two
	if isEven == false {
		return float64(medianNumber)
	}
	return float64((medianNumber + previousNumber) / 2)
}
