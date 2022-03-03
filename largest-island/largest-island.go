package largestisland

import (
	"strconv"
)

// largestRect finds the largest rectangle island
// assumes a rect must be 2x2 or greater, so no 1x2 / 2x1
func largestRect(islands [][]int) string {
	biggestIslandDimensions := ""
	biggestIslandSize := 0
	// go through each row, if you find a 1
	for i, row := range islands {
		for k, cell := range row {
			// see if the next item in the row is a 1
			if cell == 1 && k+1 < len(row) && row[k+1] == 1 {
				// find the size given the coords
				dimensions, size := findRectSize(i, k, islands)
				if size > biggestIslandSize {
					biggestIslandDimensions = dimensions
					biggestIslandSize = size
				}
			}
		}
	}

	return biggestIslandDimensions
}

// going to assume no weird shaped rects like:
// 1,1,1
// 1,1,0
func findRectSize(x int, y int, islands [][]int) (string, int) {
	maxX := 0
	maxY := 0
	// go through the column to find the last 1
	for i := y; i < len(islands); i++ {
		if islands[i][y] == 1 {
			maxY = i+1
		}
		if islands[i][y] == 0 {
			break
		}
	}

	// go through the row to find the last 1
	rowToCheck := islands[y]
	for k := x; k < len(rowToCheck); k++ {
		if rowToCheck[k] == 1 {
			maxX = k+1
		}
		if rowToCheck[k] == 0 {
			break;
		}
	}
	dimensions := strconv.Itoa(maxX) + "x" + strconv.Itoa(maxY)
	size := maxX * maxY
	// return the size in a string (e.g. "2x2")
	// also return the full size (e.g. 4)
	return dimensions, size
}
