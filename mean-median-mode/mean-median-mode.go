package meanmedianmode

import (
	"math"
	"sort"
)

type meanMedianMode struct {
	mean   float64
	median float64
	mode   []int
}

// getMeanMedianMode provides a struct of the mean, median and mode of an array of ints
func getMeanMedianMode(numbers []int) meanMedianMode {
	// first we need to sort our list of numbers
	sort.Ints(numbers)

	mean := mean(numbers)
	median := median(numbers)
	mode := mode(numbers)

	return meanMedianMode{
		mean:   mean,
		median: median,
		mode:   mode,
	}
}

// mean returns the mean of a sorted array of ints
func mean(numbers []int) float64 {
	totalNumbers := len(numbers)
	if totalNumbers == 1 {
		return float64(numbers[0])
	}

	var totalValue float64
	totalValue = 0

	for _, num := range numbers {
		totalValue = float64(num) + totalValue
	}

	return math.Round(totalValue / float64(totalNumbers))
}

// median returns the median of a sorted array of ints
func median(numbers []int) float64 {
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

// mode returns the mode of an array of ints
func mode(numbers []int) []int {
	if len(numbers) == 1 {
		return []int{numbers[0]}
	}

	numberMap := make(map[int]int)

	for _, num := range numbers {
		numberMap[num] = numberMap[num] + 1
	}

	highestOccured := 0
	highestNumber := 0
	modes := []int{}
	for number, occured := range numberMap {
		if occured > highestOccured {
			highestOccured = occured
			highestNumber = number

			// we have a new highest, reset our modes
			modes = []int{highestNumber}
			continue
		}
		// if we have a matching number, add it since we can have more than one mode
		if occured == highestOccured {
			modes = append(modes, highestNumber)
		}
	}

	// if there are no repeats in the list, then there's no mode
	if highestOccured == 1 {
		return nil
	}
	return modes
}
