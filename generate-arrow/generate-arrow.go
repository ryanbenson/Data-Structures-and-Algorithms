package generatearrow

import "strings"

// Creates an ASCII arrow pointing in the direction provided
// The size dictates how many characters to the side the point will go
func printArrow(direction string, size int) string {
	if size == 0 {
		return ""
	}
	if size == 1 {
		return "*"
	}
	currentRow := 0
	if direction == "right" {
		currentRow = 1
	} else if direction == "left" {
		currentRow = size - 1
	}
	reachedPeak := false
	arrow := ""
	arrowChar := "*"
	spacer := "  "
	for {
		if direction == "left" {
			if currentRow >= 0 {
				arrow = arrow + strings.Repeat(spacer, currentRow) + arrowChar
				if (reachedPeak == true) {
					currentRow++
				} else {
					currentRow--
				}
				if (currentRow == 0) {
					reachedPeak = true
				}
				if currentRow >= size && reachedPeak == true {
					break
				} else {
					arrow = arrow + "\n"
				}
			}
		} else if direction == "right" {
			if currentRow <= size {
				arrow = arrow + strings.Repeat(spacer, currentRow - 1) + arrowChar
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
	}

	return arrow
}
