package numbinaryonescount

import (
	"strconv"
	"strings"
)

// getBinaryOnesCount determines how many 1 characters
// are in the binary representation of a number
func getBinaryOnesCount(num int64) int {
	binaryOfInt := strconv.FormatInt(num, 2)
	return strings.Count(binaryOfInt, "1")
}