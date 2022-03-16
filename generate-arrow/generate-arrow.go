package generatearrow

import "strings"

func printArrow(direction string, size int) string {
	if size == 0 {
		return ""
	}
	if size == 1 {
		return "*"
	}
	currentRow := 1
	reachedPeak := false
	arrow := ""
	spacer := "  "
	for {
		if currentRow <= size {
			arrow = arrow + strings.Repeat(spacer, currentRow - 1) + "*"
			if (reachedPeak == true) {
				currentRow--
			} else {
				currentRow++
			}
			if (currentRow == size) {
				reachedPeak = true
			}
		}
		if currentRow == 0 {
			break
		} else {
			arrow = arrow + "\n"
		}
	}

	return arrow
}

func printArrowImplode(direction string, size int) string {
	// build an array and implode with \n
	// compare to the string builder
	return ""
}