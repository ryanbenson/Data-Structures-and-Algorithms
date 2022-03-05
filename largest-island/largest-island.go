package largestisland

import (
	"strconv"
)

// largestRect finds the largest rectangle island
// assumes rectangles are filled
// islands are declared by {int} 1, otherwise {int} 0
// the dimensions of the largest island will be provided
func largestRect(islands [][]int) string {
	biggestIslandDimensions := ""
	biggestIslandSize := 0
	// go through each row, if you find a 1
	for i, row := range islands {
		for k, cell := range row {
			// see if the next item in the row is a 1
			if cell == 1 && k+1 < len(row) && row[k+1] == 1 {
				// find the size given the coords
				dimensions, size := findRectSize(k, i, islands)
				if size > biggestIslandSize {
					biggestIslandDimensions = dimensions
					biggestIslandSize = size
				}
			}
		}
	}

	return biggestIslandDimensions
}

// finds the rectangle size of a given start point (x, y)
// assumption:
// going to assume no cut out shapes
// 1,1,1
// 1,1,0
// also no donuts
// 1,1,1
// 1,0,1
// 1,1,1
func findRectSize(x int, y int, islands [][]int) (string, int) {
	maxX := 0
	maxY := 0
	// go through the column to find the last 1
	for i := y; i < len(islands); i++ {
		if islands[i][x] == 1 {
			maxY++
		}
		if islands[i][x] == 0 {
			break
		}
	}

	// go through the row to find the last 1
	rowToCheck := islands[y]
	for k := x; k < len(rowToCheck); k++ {
		if rowToCheck[k] == 1 {
			maxX++
		}
		if rowToCheck[k] == 0 {
			break;
		}
	}
	// ensure the last corner exists
	if islands[maxY-1][maxX-1] != 1 {
		return "", 0
	}
	dimensions := strconv.Itoa(maxX) + "x" + strconv.Itoa(maxY)
	size := maxX * maxY
	// return the size in a string (e.g. "2x2")
	// also return the full size (e.g. 4)
	return dimensions, size
}
