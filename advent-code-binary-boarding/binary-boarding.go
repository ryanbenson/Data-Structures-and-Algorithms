package binaryboarding

import (
	"fmt"
	"math"
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
	start := 0.0
	end := 127.0
	num := 0.0
	list := strings.Split(letters, "")
	for i, letter := range list {
		if letter == "F" {
			if i == 6 {
				num = start
			}
			end = math.Floor((start + end) / 2)
		}
		if letter == "B" {
			if i == 6 {
				num = end
			}
			start = math.Ceil((start + end) / 2)
		}
	}
	return int(num)
}

func getColumn(letters string) int {
	start := 0.0
	end := 7.0
	num := 0.0
	list := strings.Split(letters, "")
	for i, letter := range list {
		if letter == "L" {
			if i == 2 {
				num = start
			}
			end = math.Floor((start + end) / 2)
		}
		if letter == "R" {
			if i == 2 {
				num = end
			}
			start = math.Ceil((start + end) / 2)
		}
	}
	return int(num)
}

func getNum(row int, column int) int {
	return row*8 + column
}
