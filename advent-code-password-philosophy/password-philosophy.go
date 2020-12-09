package passwordphilosophy

import (
	"regexp"
	"strconv"
	"strings"
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
// example: 1-3 x: abcxyz = valid (1-3 instances of x in the passphrase)
// example: 4-12 y: abcxyz = invalid, only one y
func isValidPassword(password string) bool {
	re := regexp.MustCompile(`([\d]+)-([\d]+)[\s]([a-z]):\s([a-z]+)`)
	m := re.FindSubmatch([]byte(password))

	// normally want to catch the error >_>;
	min, _ := strconv.Atoi(string(m[1]))
	max, _ := strconv.Atoi(string(m[2]))
	letter := string(m[3])
	phrase := string(m[4])

	// count how many letters that are that we need
	countOfLetter := strings.Count(phrase, letter)

	// min / max can equal the count
	if countOfLetter >= min && countOfLetter <= max {
		return true
	}
	return false
}
