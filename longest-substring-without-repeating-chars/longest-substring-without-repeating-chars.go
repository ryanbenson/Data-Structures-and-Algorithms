package longestsubstringwithoutrepeatingchars

// Gets the longest substring of letters (its length) that does not have a repeated character in it
func getLongestSubstrLengthWithoutRepeats(haystack string) int {
	// don't need to bother figuring it out if it's a single string
	haystackLen := len(haystack)
	if haystackLen == 1 {
		return 1
	}

	// keeps track of our longest substring length
	longestSubstringLen := 0
	// the starting char index of our window
	startingCharIndex := 0
	// the leading char index of our window
	charIndex := 0

	for startingCharIndex != haystackLen -1 {
		// get our window of the substring, and check if it's a unique list of letters
		stringGroup := string(haystack[startingCharIndex:charIndex])
		isUnique := isStringUniqueChars(stringGroup)

		// if it's a unique list of letters and it's the longest one so far, update our longest int
		if isUnique == true && len(stringGroup) > longestSubstringLen {
			longestSubstringLen = len(stringGroup)
		}

		// update the right side of our window
		charIndex++

		// if our right side of the window is longer than the window bounds, reset the window to move up a letter
		if charIndex > haystackLen {
			startingCharIndex++
			// always reset the charindex to our starting index, not 0
			charIndex = startingCharIndex
		}
	}

	return longestSubstringLen
}

// Determines if a string is made up of unique letters only or not
func isStringUniqueChars(sequence string) bool {
	// using a map since it's resource beneficial and easy to lookup
	letterMap := make(map[string]int)
	// for all of the letters, check if we have a letter in our map
	for i, letter := range sequence {
		stringLetter := string(letter)
		_, ok := letterMap[stringLetter]

		// if we found a matching letter, it's not unique so we can exit
		if ok {
			return false
		}

		// update our map so we can check subsequent letters
		letterMap[stringLetter] = i
	}

	return true
}