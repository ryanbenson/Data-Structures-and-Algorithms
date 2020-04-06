package comparewithbackspace

// Compare two strings that might have backspaces in it ("#")
// to see if they match after processing them
func compareWithBackspace(sequence1 string, sequence2 string) bool {
	return buildString(sequence1) == buildString(sequence2)
}

// process a string sequence, which includes deleting a character
// using "#", and return the final result
func buildString(sequence string) string {
	finalString := ""
	for _, letter := range sequence {
		stringLetter := string(letter)
		if stringLetter == "#" {
			strLen := len(finalString)
			if strLen > 0 {
				finalString = finalString[:strLen-1]
			}
		} else {
			finalString = finalString + stringLetter
		}
	}
	return finalString
}