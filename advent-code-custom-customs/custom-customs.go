package customcustoms

import (
	"fmt"
	"strings"
)

func sum(questions string) int {
	rows := getRows(questions)
	for _, row := range rows {
		rowAnswers := getRowAnswers(row)
		fmt.Println(rowAnswers)
	}
	return 0
}

func getRows(questions string) []string {
	return strings.Split(questions, "\n\n")
}

func getRowAnswers(row string) []string {
	return strings.Split(row, "\n")
}
