package babylisp

import (
	"errors"
	"testing"
)

func TestProcessAdd(t *testing.T) {
	result, err := process("(add 1 2)")
	expected := 3

	if result != expected {
		t.Errorf("When processing a calculation, got: %v, want: %v.", result, expected)
	}

	if err != nil {
		t.Errorf("When processing a calculation and checking the error, got: %v, want: %v.", err, nil)
	}
}

func TestProcessSubtract(t *testing.T) {
	result, err := process("(subtract 1 2)")
	expected := -1

	if result != expected {
		t.Errorf("When processing a calculation, got: %v, want: %v.", result, expected)
	}

	if err != nil {
		t.Errorf("When processing a calculation and checking the error, got: %v, want: %v.", err, nil)
	}
}

func TestProcessDivide(t *testing.T) {
	result, err := process("(divide 99 11)")
	expected := 9

	if result != expected {
		t.Errorf("When processing a calculation, got: %v, want: %v.", result, expected)
	}

	if err != nil {
		t.Errorf("When processing a calculation and checking the error, got: %v, want: %v.", err, nil)
	}
}

func TestProcessMultiply(t *testing.T) {
	result, err := process("(multiply 5 7)")
	expected := 35

	if result != expected {
		t.Errorf("When processing a calculation, got: %v, want: %v.", result, expected)
	}

	if err != nil {
		t.Errorf("When processing a calculation and checking the error, got: %v, want: %v.", err, nil)
	}
}

func TestProcessEmptyParanthesis(t *testing.T) {
	result, err := process("()")
	expectedErr := errors.New("Equation too short")

	if result != 0 {
		t.Errorf("When processing an empty calculation, got: %v, want: %v.", result, nil)
	}

	if err == nil {
		t.Errorf("When processing an empty calculation and checking the error, got: %v, want: %v.", err, expectedErr)
	}
}

func TestProcessEmpty(t *testing.T) {
	result, err := process("")

	if result != 0 {
		t.Errorf("When processing an empty string, got: %v, want: %v.", result, nil)
	}

	if err != nil {
		t.Errorf("When processing an empty string and checking the error, got: %v, want: %v.", err, nil)
	}
}

func TestProcessMissingOpen(t *testing.T) {
	result, err := process("subtract 1 2)")
	expectedErr := errors.New("Invalid start of equation")

	if result != 0 {
		t.Errorf("When processing an invalid string, got: %v, want: %v.", result, nil)
	}

	if err == nil {
		t.Errorf("When processing an invalid string and checking the error, got: %v, want: %v.", err, expectedErr)
	}
}

func TestProcessMissingClose(t *testing.T) {
	result, err := process("(subtract 1 2")
	expectedErr := errors.New("Invalid start of equation")

	if result != 0 {
		t.Errorf("When processing an invalid string, got: %v, want: %v.", result, nil)
	}

	if err == nil {
		t.Errorf("When processing an invalid string and checking the error, got: %v, want: %v.", err, expectedErr)
	}
}

func TestProcessMissingElements(t *testing.T) {
	result, err := process("(add 1)")
	expectedErr := errors.New("Invalid number of equation elements")

	if result != 0 {
		t.Errorf("When processing an invalid string, got: %v, want: %v.", result, nil)
	}

	if err == nil {
		t.Errorf("When processing an invalid string and checking the error, got: %v, want: %v.", err, expectedErr)
	}
}

func TestProcessTooManyElements(t *testing.T) {
	result, err := process("(add 1 3 5 7 8 9 10)")
	expectedErr := errors.New("Invalid number of equation elements")

	if result != 0 {
		t.Errorf("When processing an invalid string, got: %v, want: %v.", result, nil)
	}

	if err == nil {
		t.Errorf("When processing an invalid string and checking the error, got: %v, want: %v.", err, expectedErr)
	}
}

func TestProcessInvalidNumbers(t *testing.T) {
	result, err := process("(add 1 hello)")
	expectedErr := errors.New("Invalid numbers provided")

	if result != 0 {
		t.Errorf("When processing an invalid string, got: %v, want: %v.", result, nil)
	}

	if err == nil {
		t.Errorf("When processing an invalid string and checking the error, got: %v, want: %v.", err, expectedErr)
	}
}

func TestProcessInvalidOperand(t *testing.T) {
	result, err := process("(testing 1 hello)")
	expectedErr := errors.New("Invalid operand")

	if result != 0 {
		t.Errorf("When processing an invalid string, got: %v, want: %v.", result, nil)
	}

	if err == nil {
		t.Errorf("When processing an invalid string and checking the error, got: %v, want: %v.", err, expectedErr)
	}
}
