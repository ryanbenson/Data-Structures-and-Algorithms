package binaryboarding

import (
	"fmt"
	"strings"
)

func getMaxId(passes string) int {
	maxId := 0
	passList := strings.Split(passes, "\n")
	for _, pass := range passList {
		id := getSeatId(pass)
		if id > maxId {
			maxId = id
		}
	}
	return maxId
}

func getSeatId(pass string) int {
	// totalRows := 128 // 0-127
	letters := strings.Split(pass, "")
	for _, letter := range letters {
		fmt.Println(letter)
	}
	return 0
}
