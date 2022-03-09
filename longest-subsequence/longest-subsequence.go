package longestsubsequence

import (
	"sort"
)

// finds the longest subsequence of numbers within a list of numbers
// providing the length of the longest subsequence found
func longestSubSeq(sequence []int) int {
	longest := 0
	currentSequence := 1
	sort.Ints(sequence)
	for i, num := range sequence {
		// if we've run out of room, short circuit
		if i+1 >= len(sequence) {
			break
		}
		// is the next item in the array our adjacent sibling number?
		if num + 1 == sequence[i+1] {
			currentSequence++
		} else {
			// that's a no. we've hit the end of our subsequence
			if currentSequence > longest {
				longest = currentSequence
			}
			// if we're too far into the list of numbers
			// and there can't be a bigger sequence, then short circuit
			if i+longest > len(sequence) {
				break
			}
			currentSequence = 1
		}
	}
	return longest
}