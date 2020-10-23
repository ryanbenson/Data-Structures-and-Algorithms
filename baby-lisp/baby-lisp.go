package babylisp

import (
	"errors"
	"fmt"
	"strings"
)

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
	fmt.Println(elements)

	return 0, nil
}
