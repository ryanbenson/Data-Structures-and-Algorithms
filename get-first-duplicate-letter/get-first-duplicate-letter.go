package getfirstduplicateletter

import "strings"

// Provides the first letter that is a repeated letter in the word given
// Underlying: it leverages a map to manage the letters used or not
func getFirstDuplicateLetterMap(word string) string {
	letters := strings.Split(word, "")

	letterMap := make(map[string]int)

	for _, letter := range letters {
		_, ok := letterMap[letter]
		if ok {
			return letter
		}
		letterMap[letter] = 1
	}

	// never found one
	return ""
}

// Provides the first letter that is a repeated letter in the word given
// Underlying: It goes through each letter and finds/counts all letters matching
func getFirstDuplicateLetterSearch(word string) string {
	letters := strings.Split(word, "")

	for _, letter := range letters {
		count := strings.Count(word, letter)
		if count > 1 {
			return letter
		}
	}

	return ""
}
