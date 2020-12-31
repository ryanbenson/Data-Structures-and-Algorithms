package passportprocessing

import (
	"regexp"
	"strconv"
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
	pieces := getKeyValues(passportPieces)
	isValid := hasValidKeys(pieces)
	return isValid
}

func getKeyValues(passportPieces []string) [][]string {
	allPieces := [][]string{}
	for _, piece := range passportPieces {
		pieces := strings.Split(piece, ":")
		allPieces = append(allPieces, pieces)
	}
	return allPieces
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
func hasValidKeys(pieces [][]string) bool {
	hasByr := false
	hasIyr := false
	hasEyr := false
	hasHgt := false
	hasHcl := false
	hasEcl := false
	hasPid := false
	for _, piece := range pieces {
		if piece[0] == "byr" {
			hasByr = isValidByr(piece[1])
		}
		if piece[0] == "iyr" {
			hasIyr = isValidIyr(piece[1])
		}
		if piece[0] == "eyr" {
			hasIyr = isValidEyr(piece[1])
		}
		if piece[0] == "hgt" {
			hasHgt = isValidHgt(piece[1])
		}
		if piece[0] == "hcl" {
			hasHcl = isValidHcl(piece[1])
		}
		if piece[0] == "ecl" {
			hasEcl = isValidEcl(piece[1])
		}
		if piece[0] == "pid" {
			hasPid = true
		}
	}
	if hasByr == false || hasIyr == false || hasEyr == false || hasHgt == false || hasHcl == false || hasEcl == false || hasPid == false {
		return false
	}
	return true
}

func isValidByr(yearString string) bool {
	year, _ := strconv.Atoi(yearString)
	if year >= 1920 && year <= 2002 {
		return true
	}
	return false
}

func isValidIyr(yearString string) bool {
	year, _ := strconv.Atoi(yearString)
	if year >= 2010 && year <= 2020 {
		return true
	}
	return false
}

func isValidEyr(yearString string) bool {
	year, _ := strconv.Atoi(yearString)
	if year >= 2020 && year <= 2030 {
		return true
	}
	return false
}

func isValidHgt(height string) bool {
	// last two chars should be the measurement type
	measurementType := height[len(height)-2:]
	measurementValue := height[0 : len(height)-2]
	value, _ := strconv.Atoi(measurementValue)
	if measurementType == "in" {
		if value >= 59 && value <= 76 {
			return true
		}
	}
	if measurementType == "cm" {
		if value >= 150 && value <= 193 {
			return true
		}
	}
	// if it's not inches or centimeters, it's invalid
	return false
}

func isValidHcl(color string) bool {
	// must start with #, and followed by 0-9, a-f
	re := regexp.MustCompile("#[0-9a-f]+")
	return re.MatchString(color)
}

func isValidEcl(color string) bool {
	// must be one of:
	validOptions := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, option := range validOptions {
		if option == color {
			return true
		}
	}
	return false
}
