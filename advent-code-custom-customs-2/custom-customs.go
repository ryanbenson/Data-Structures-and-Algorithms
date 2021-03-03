package customcustoms

import (
	"fmt"
	"strings"
)

func sum(answers string) int {
	rows := getRows(answers)
	total := 0
	for _, row := range rows {
		count := getUniqueAnswersCount(row)
		total = total + count
	}
	return total
}

func getRows(answers string) []string {
	return strings.Split(answers, "\n\n")
}

func getRowAnswersMap(row string) map[string]int {
	arr := strings.Split(row, "\n")
	letterMap := make(map[string]int)
	allAnswers := ""
	for _, r := range arr {
		allAnswers = allAnswers + r
	}
	uniqueAnswers := ""

	for _, letter := range allAnswers {
		num, ok := letterMap[string(letter)]
		if ok {
			letterMap[string(letter)] = num + 1
			continue
		} else {
			letterMap[string(letter)] = 1
		}
		uniqueAnswers = uniqueAnswers + string(letter)
	}
	return letterMap
}

func getNumberOfPeopleInRow(row string) int {
	arr := strings.Split(row, "\n")
	return len(arr)
}

func getUniqueAnswersCount(row string) int {
	m := getRowAnswersMap(row)
	peopleInRow := getNumberOfPeopleInRow(row)
	fmt.Println(peopleInRow)
	fmt.Println(m)
	// to do, go through the map, and see which letters have
	// every person in the row answering that question
	// add up all of the letters that have the letter == peopleInRow
	return 1
}
