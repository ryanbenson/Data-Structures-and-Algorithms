package twosum

import "fmt"

func main() {
	numbers := []int{9, 1, 2, 3, 4, 5}
	target := 5
	result := twoSum(numbers, target)
	fmt.Println(result)
}

func twoSum(list []int, target int) [2]int {
	var result [2]int
	found := false

	for i, v := range list {
		if v >= target {
			continue
		}

		for k, val := range list {
			if val >= target {
				continue
			}
			sum := v + val

			if sum == target {
				result[0] = i
				result[1] = k
				found = true
				break
			}
		}

		if found == true {
			break
		}
	}
	return result
}
