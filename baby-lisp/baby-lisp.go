package babylisp

import (
	"errors"
	"strconv"
	"strings"
)

var operations = []string{"add", "subtract", "multiply", "divide"}

func process(equation string) (int, error) {
	if len(equation) == 0 {
		return 0, nil
	}

	if len(equation) == 1 {
		return 0, errors.New("Equation too short")
	}

	openChar := "("
	closeChar := ")"

	if string(equation[0]) != openChar {
		return 0, errors.New("Invalid start of equation")
	}

	if string(equation[len(equation)-1]) != closeChar {
		return 0, errors.New("Invalid end of equation")
	}

	eq := equation[1 : len(equation)-1]
	elements := strings.Fields(eq)

	if len(elements) != 3 {
		return 0, errors.New("Invalid number of equation elements")
	}

	operand := elements[0]
	if isValidOperand(operand, operations) == false {
		return 0, errors.New("Invalid operand")
	}

	result := 0

	x, err := strconv.Atoi(elements[1])
	y, err := strconv.Atoi(elements[2])

	if err != nil {
		return 0, errors.New("Invalid numbers provided")
	}

	if operand == operations[0] {
		result = x + y
	}
	if operand == operations[1] {
		result = x - y
	}
	if operand == operations[2] {
		result = x * y
	}
	if operand == operations[3] {
		result = x / y
	}

	return result, nil
}

func isValidOperand(operand string, operations []string) bool {
	for _, validOperand := range operations {
		if validOperand == operand {
			return true
		}
	}
	return false
}
