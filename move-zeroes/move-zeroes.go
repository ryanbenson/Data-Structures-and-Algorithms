package movezeroes

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
