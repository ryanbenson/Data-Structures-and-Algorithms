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

func getRowAnswers(row string) []string {
	return strings.Split(row, "\n")
}

func getUniqueAnswers(row string) []string {
	rowAnswers := getRowAnswers(row)
	fmt.Println(rowAnswers)
	fmt.Println("======")
	return []string{}
}
