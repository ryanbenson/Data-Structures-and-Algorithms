package customcustoms

import (
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

func getUniqueRowAnswers(row string) string {
	arr := strings.Split(row, "\n")
	letterMap := make(map[string]int)
	allAnswers := ""
	for _, r := range arr {
		allAnswers = allAnswers + r
	}
	uniqueAnswers := ""

	for _, letter := range allAnswers {
		_, ok := letterMap[string(letter)]
		if ok {
			continue
		}
		letterMap[string(letter)] = 1
		uniqueAnswers = uniqueAnswers + string(letter)
	}
	return uniqueAnswers
}

func getUniqueAnswersCount(row string) int {
	rowAnswers := getUniqueRowAnswers(row)
	return len(rowAnswers)
}
