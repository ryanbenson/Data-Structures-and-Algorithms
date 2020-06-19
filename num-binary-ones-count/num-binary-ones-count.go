package numbinaryonescount

import (
	"strconv"
	"strings"
)

func getBinaryOnesCount(num int64) int {
	binaryOfInt := strconv.FormatInt(num, 2)
	return strings.Count(binaryOfInt, "1")
}