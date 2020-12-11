package toboggantrajectory

import (
	"strings"
)

func findTrees(mapContent string) int {
	mapData := strings.Split(mapContent, "\n")

	right := 3
	down := 1
	tree := "#"

	totalLines := len(mapData)
	lineLimit := len(mapData[0])
	currentLine := 0
	curLinePos := 0
	totalTrees := 0

	for {
		// go right
		curLinePos = curLinePos + right
		// if we reach the end of the right side of the map
		// then get the remainder and restart since it repeats
		if curLinePos >= lineLimit {
			curLinePos = curLinePos - lineLimit
		}

		// go down the map
		currentLine = currentLine + down
		// if we reach the end of the map, exit
		if currentLine >= totalLines {
			break
		}

		// did we find a tree?
		mapItem := string(mapData[currentLine][curLinePos])
		if mapItem == tree {
			totalTrees = totalTrees + 1
		}
	}

	return totalTrees
}
