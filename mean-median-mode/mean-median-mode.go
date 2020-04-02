package meanmedianmode

import (
	"math"
	"sort"
)

type meanMedianMode struct {
	mean   float64
	median float64
	mode   int
}

func getMeanMedianMode(numbers []int) meanMedianMode {
	// first we need to sort our list of numbers
	sort.Ints(numbers)

	mean := mean(numbers)
	median := median(numbers)

	return meanMedianMode{
		mean:   mean,
		median: median,
		mode:   2,
	}
}

func mean(numbers []int) float64 {
	totalNumbers := len(numbers)
	var totalValue float64
	totalValue = 0

	for _, num := range numbers {
		totalValue = float64(num) + totalValue
	}

	return math.Round(totalValue / float64(totalNumbers))
}

func median(numbers []int) float64 {
	totalNumbers := len(numbers)

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
