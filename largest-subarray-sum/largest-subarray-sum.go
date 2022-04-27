package largestsubarraysum

import (
	"errors"
)

func getSum(list []int, subsetCount int) ([]int, error) {
	// idea 1:
	// go through the list, split it into smaller pieces by the subsetcount
	// 1,2,3,4,5,6 | 3
	// 1,2,3 | 2,3,4 | 3,4,5 | 4,5,6
	// go through those and find the largest sum
	listLength := len(list)
	if listLength < subsetCount {
		return nil, errors.New("Subset count cannot be larger than the list size")
	}
	return []int{}, nil
}