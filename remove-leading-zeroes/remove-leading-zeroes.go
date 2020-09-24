package removeleadingzeroes

import (
	"sort"
)

// returns the list of numbers without all leading zeroes
// zeroes after a n>=1 is fine
func removeSearchList(list []int) []int {
	index := sort.Search(len(list), func(i int) bool { return list[i] > 0 })
	return list[index:len(list)]
}

// returns the list of numbers without all leading zeroes
// zeroes after n>=1 is fine, using a for loop
func removeForLoop(list []int) []int {
	index := 0
	for i, num := range list {
		if num > 0 {
			index = i
			break
		}
	}
	return list[index:len(list)]
}
