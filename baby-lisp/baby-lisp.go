package babylisp

import (
	"errors"
	"strconv"
	"strings"
)

// the list of possible operations
var operations = []string{"add", "subtract", "multiply", "divide"}

// Takes a lispy string that's a simple calculation and processes it
func process(equation string) (int, error) {
	// if an equation has no length
	if len(equation) == 0 {
		return 0, nil
	}

	// smallest equation size "(add 1 1)"
	minEquationSize := 9
	if len(equation) < minEquationSize {
		return 0, errors.New("Equation too short")
	}

	// check that the equation is wrapped in ()
	openChar := "("
	closeChar := ")"

	if string(equation[0]) != openChar {
		return 0, errors.New("Invalid start of equation")
	}

	if string(equation[len(equation)-1]) != closeChar {
		return 0, errors.New("Invalid end of equation")
	}

	// take everything in between the ()
	eq := equation[1 : len(equation)-1]
	// split the string by spaces
	elements := strings.Fields(eq)

	// if we have more or less than 3 items, then the equation is invalid
	if len(elements) != 3 {
		return 0, errors.New("Invalid number of equation elements")
	}

	// validate the operand
	operand := elements[0]
	if isValidOperand(operand, operations) == false {
		return 0, errors.New("Invalid operand")
	}

	result := 0

	// convert the numbers into ints, check for any issues
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

// Determines if the operand provided is in the whitelist
func isValidOperand(operand string, operations []string) bool {
	for _, validOperand := range operations {
		if validOperand == operand {
			return true
		}
	}
	return false
}
