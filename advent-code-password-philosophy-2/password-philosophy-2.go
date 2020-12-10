package passwordphilosophy2

import (
	"regexp"
	"strconv"
)

// Determines how many passwords are valid given a list
// of passwords with their way to validate them
func passwordsValidNum(passwords []string) int {
	validCount := 0

	for _, password := range passwords {
		isValid := isValidPassword(password)
		if isValid == true {
			validCount = validCount + 1
		}
	}

	return validCount
}

// Determines if a single password is valid or not
// note: index is 1 based, not 0
// example: 1-3 a: abcde = valid: pos 1 has a. pos 3 does not
// example: 1-3 b: cdefg = invalid: pos 1 & 3 does not have b
// example: 2-9 c: ccccccccc = invalid: both pos 1 & 3 has b
func isValidPassword(password string) bool {
	re := regexp.MustCompile(`([\d]+)-([\d]+)[\s]([a-z]):\s([a-z]+)`)
	m := re.FindSubmatch([]byte(password))

	// normally want to catch the error >_>;
	pos1, _ := strconv.Atoi(string(m[1]))
	pos2, _ := strconv.Atoi(string(m[2]))
	letter := string(m[3])
	phrase := string(m[4])

	letter1 := string(phrase[pos1-1])
	letter2 := string(phrase[pos2-1])

	letter1Match := letter1 == letter
	letter2Match := letter2 == letter

	if letter1Match == letter2Match {
		return false
	}

	return true
}
