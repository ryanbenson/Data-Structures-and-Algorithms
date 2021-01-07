package binaryboarding

import (
	"fmt"
	"strings"
)

func getMaxId(passes string) int {
	maxId := 0
	passList := strings.Split(passes, "\n")
	for _, pass := range passList {
		row, column, id := getSeat(pass)
		if id > maxId {
			maxId = id
		}
	}
	return maxId
}

func getSeat(pass string) (int, int, int) {
	// totalRows := 128 // 0-127
	letters := strings.Split(pass, "")
	for _, letter := range letters {
		fmt.Println(letter)
	}
	return 0, 0, 0
}
