package longestsubstringwithoutrepeatingchars

func getLongestSubstrLengthWithoutRepeats(haystack string) int {
	// don't need to bother figuring it out if it's a single string
	if len(haystack) == 1 {
		return 1
	}
	return 0
}