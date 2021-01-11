package binaryboarding

import (
	"fmt"
	"strings"
)

// Returns the max id amongs a list of passes
func getMaxId(passes string) int {
	maxId := 0
	passList := strings.Split(passes, "\n")
	for _, pass := range passList {
		row, column, id := getSeat(pass)
		fmt.Println(row)
		fmt.Println(column)
		if id > maxId {
			maxId = id
		}
	}
	return maxId
}

// Gets the seat data for a specific pass
func getSeat(pass string) (int, int, int) {
	frontToBack := pass[0:7]
	leftToRight := pass[7:10]
	fmt.Println(frontToBack, leftToRight)
	letters := strings.Split(pass, "")
	for _, letter := range letters {
		fmt.Println(letter)
	}
	return 0, 0, 0
}
