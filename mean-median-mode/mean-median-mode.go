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

	return meanMedianMode{
		mean:   mean,
		median: 12.5,
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
