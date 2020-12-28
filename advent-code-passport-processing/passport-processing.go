package passportprocessing

import (
	"regexp"
	"strings"
)

func process(passports string) int {
	validPasswordsNum := 0

	passportsList := strings.Split(passports, "\n\n")
	for _, passport := range passportsList {
		isValid := isValidPassport(passport)
		if isValid == true {
			validPasswordsNum++
		}
	}
	return validPasswordsNum
}

func isValidPassport(passport string) bool {
	re := regexp.MustCompile(`[\s]+`)
	passportPieces := re.Split(passport, -1)
	keys := getKeys(passportPieces)
	isValid := hasValidKeys(keys)
	return isValid
}

func getKeys(passportPieces []string) []string {
	keys := []string{}
	for _, piece := range passportPieces {
		pieces := strings.Split(piece, ":")
		key := pieces[0]
		keys = append(keys, key)
	}
	return keys
}

// Pieces of a passport. All fields are requied except cid
// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)
func hasValidKeys(keys []string) bool {
	hasByr := false
	hasIyr := false
	hasEyr := false
	hasHgt := false
	hasHcl := false
	hasEcl := false
	hasPid := false
	for _, key := range keys {
		if key == "byr" {
			hasByr = true
		}
		if key == "iyr" {
			hasIyr = true
		}
		if key == "eyr" {
			hasEyr = true
		}
		if key == "hgt" {
			hasHgt = true
		}
		if key == "hcl" {
			hasHcl = true
		}
		if key == "ecl" {
			hasEcl = true
		}
		if key == "pid" {
			hasPid = true
		}
	}
	if hasByr == false || hasIyr == false || hasEyr == false || hasHgt == false || hasHcl == false || hasEcl == false || hasPid == false {
		return false
	}
	return true
}
