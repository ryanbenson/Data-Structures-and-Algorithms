package movezeroes

// moves all zeroes to the end of an array
func moveZeroes(numbers []int) []int {
	numbersLen := len(numbers)
	// this will make an array for as long as the numbers is,
	// with the default value of 0 for all values
	finalNumbers := make([]int, numbersLen)

	index := 0
	// for every valid number, update the moving index so they end up being
	// added sequentially
	for _, num := range numbers {
		if num != 0 {
			finalNumbers[index] = num
			index++
		}
	}

	return finalNumbers
}

// moves all zeroes to the end of an array
// uses two arrays and merges them together in the end
func moveZeroesMerge(numbers []int) []int {
	var nonZeroes []int
	numbersLen := len(numbers)

	for _, num := range numbers {
		if num > 0 {
			nonZeroes = append(nonZeroes, num)
		}
	}

	// how many are left?
	remainingNums := numbersLen - len(nonZeroes)

	// make a new array with the remaining
	zeroes := make([]int, remainingNums)

	// now merge them
	finalNumbers := append(nonZeroes, zeroes...)

	return finalNumbers
}
