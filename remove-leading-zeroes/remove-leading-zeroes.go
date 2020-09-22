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
