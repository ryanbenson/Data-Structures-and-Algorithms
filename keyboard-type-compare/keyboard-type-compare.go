package keyboardtype

import (
	"strings"
)

// determines if two keyboard inputs are the same
func isEqual(str1 string, str2 string) bool {
	finalStr1 := inputToString(str1)
	finalStr2 := inputToString(str2)
	return finalStr1 == finalStr2
}

// converts an input of keyboard character sequence to a final string
func inputToString(input string) string {
	letters := strings.Split(input, "")
	var finalStr []string
	inputLen := len(letters)
	deletes := 0

	for i := 0; i < inputLen; i++ {
		letter := letters[i]
		if letter == "#" {
			// if we have any letters, remove the last one
			if len(finalStr) > 0 {
				finalStr = finalStr[:len(finalStr)-1]
			}
			continue
		}
		if letter == "%" {
			deletes++
			continue
		}
		if deletes > 0 {
			deletes--
			continue
		}
		finalStr = append(finalStr, letter)
	}
	returnStr := strings.Join(finalStr, "")
	return returnStr
}
