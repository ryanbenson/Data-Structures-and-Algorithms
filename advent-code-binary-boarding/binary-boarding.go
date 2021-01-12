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

	row := getRow(frontToBack)
	column := getColumn(leftToRight)
	num := getNum(row, column)
	return row, column, num
}

func getRow(letters string) int {
	// limit := 127
	list := strings.Split(letters, "")
	for _, letter := range list {
		fmt.Println(letter)
	}
	return 0
}

func getColumn(letters string) int {
	// limit := 7
	return 0
}

func getNum(row int, column int) int {
	return row*8 + column
}
