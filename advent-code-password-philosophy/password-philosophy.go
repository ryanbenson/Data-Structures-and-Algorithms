package passwordphilosophy

import (
	"regexp"
	"strconv"
	"strings"
)

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
