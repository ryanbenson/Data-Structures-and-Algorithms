package reportrepair

// provide a list of ints where two numbers will
// sum to 2020, using those two numbers,
// multiply them together. that's the return
func repair(numbers []int) int {
	x := 0
	y := 0
	z := 0
	totalNumbers := len(numbers)
	goal := 2020

	for i, num := range numbers {
		// don't bother going back in the list of numbers, start at i
		for k := i; k < totalNumbers; k++ {
			for j := k; j < totalNumbers; j++ {
				sum := num + numbers[k] + numbers[j]
				if sum == goal {
					x = num
					y = numbers[k]
					z = numbers[j]
					break
				}
			}
		}
	}

	result := x * y * z
	return result
}
