package keyboardtype

// determines if two keyboard inputs are the same
func isEqual(str1 string, str2 string) bool {
	finalStr1 := inputToString(str1)
	finalStr2 := inputToString(str2)
	return finalStr1 == finalStr2
}

// converts an input of keyboard character sequence to a final string
func inputToString(input string) string {
	return ""
}
