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

func getSumSequential(list []int, subsetCount int) ([]int, error) {
	listLength := len(list)
	if listLength < subsetCount {
		return nil, errors.New("Subset count cannot be larger than the list size")
	}
	// idea 2: this can do it in one go with no overhead of more data to be created/processed
	// go a for loop for the length of the list
	// keep track of the highest number and subset array
	// as you increment like normal, keep track of where you start and the subset count use
	// incremment using the subset count first, then when you hit the subsetcount, bump
	// where you were in the loop, and reset back with the subset count
	// keep going until you hit the end and the i + subsetcount exceed the length
	// rough outline:
	/*
	biggest count = 0
	bigger subset = []{int}
	current count = 0
	current place in list = 0
	for i = 0; i < len(list); i++ {
		current count += list[i]
		if i >= subsetcount {
			i = 0
		}
		current place in list ++
	}
	*/
	return []int{}, nil
}