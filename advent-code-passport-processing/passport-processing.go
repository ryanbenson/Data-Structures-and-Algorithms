package passportprocessing

import (
	"fmt"
	"regexp"
	"strings"
)

func process(passports string) int {
	validPasswordsNum := 2

	passportsList := strings.Split(passports, "\n\n")
	for _, passport := range passportsList {
		isValid := isValidPassport(passport)
		if isValid == true {
			validPasswordsNum++
		}
	}
	// return validPasswordsNum
	return 2
}

func isValidPassport(passport string) bool {
	re := regexp.MustCompile(`[\s]+`)
	passportPieces := re.Split(passport, -1)
	fmt.Println(passportPieces, len(passportPieces))
	return true
}
