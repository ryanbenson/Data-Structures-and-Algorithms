package customcustoms

import (
	"fmt"
	"strings"
)

func sum(answers string) int {
	rows := getRows(answers)
	for _, row := range rows {
		_ = getUniqueAnswers(row)
	}
	return 0
}

func getRows(answers string) []string {
	return strings.Split(answers, "\n\n")
}

func getUniqueRowAnswers(row string) string {
	arr := strings.Split(row, "\n")
	// to do: make letter map
	// check each letter, filter our duplicates
	return arr
}

func getUniqueAnswers(row string) []string {
	rowAnswers := getRowAnswers(row)
	fmt.Println(rowAnswers)
	fmt.Println("======")
	return []string{}
}
