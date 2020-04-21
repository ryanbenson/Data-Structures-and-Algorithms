package longestsubstringwithoutrepeatingchars

import "fmt"

func getLongestSubstrLengthWithoutRepeats(haystack string) int {
	// don't need to bother figuring it out if it's a single string
	haystackLen := len(haystack)
	if haystackLen == 1 {
		return 1
	}

	startingCharIndex := 0
	charIndex := 0
	for startingCharIndex != haystackLen -1 {

		fmt.Println(haystack[startingCharIndex:charIndex])
		charIndex += 1

		// TODO: check if string we get from haystack is unique list of chars or not

		if charIndex >= haystackLen -1 {
			startingCharIndex += 1
			// always reset the charindex to our starting index, not 0
			charIndex = startingCharIndex
		}

		if startingCharIndex >= haystackLen -1 {
			break
		}
	}

	return longestSubstringLen
}