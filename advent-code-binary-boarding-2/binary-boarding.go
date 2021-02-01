package binaryboarding

import (
	"math"
	"strings"
)

func find(passes string) int {
	maxID := getMaxID(passes)
	availableSeats := getAvailableSeats(maxID)
	passList := strings.Split(passes, "\n")
	for _, pass := range passList {
		_, _, id := getSeat(pass)
		availableSeats = removeAvailableSeat(availableSeats, id)
	}
	return availableSeats[0]
}

func removeAvailableSeat(seats []int, seat int) []int {
	matchedIndex := 0
	for i, v := range seats {
		if seat == v {
			matchedIndex = i
			break
		}
	}
	ret := make([]int, 0)
	ret = append(ret, seats[:matchedIndex]...)
	return append(ret, seats[matchedIndex+1:]...)
}

// Returns the list of all available seatas
func getAvailableSeats(max int) []int {
	seats := []int{}
	// ignore the first set, 0-7
	for i := 7; i < max; i++ {
		seats = append(seats, i)
	}
	return seats
}

// Returns the max id among a list of passes
func getMaxID(passes string) int {
	maxID := 0
	passList := strings.Split(passes, "\n")
	for _, pass := range passList {
		_, _, id := getSeat(pass)
		if id > maxID {
			maxID = id
		}
	}
	return maxID
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

// Gets the row based on the letter sequence provided
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

// Gets the column based on the letter sequence provided
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

// Gets the number of the seat
func getNum(row int, column int) int {
	return row*8 + column
}
