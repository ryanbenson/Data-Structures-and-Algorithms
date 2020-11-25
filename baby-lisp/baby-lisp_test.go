package babylisp

import (
	"errors"
	"testing"
)

func TestProcess(t *testing.T) {
	result, err := process("(add 1 2)")
	expected := 3

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
	result, err := process("add 1 2)")
	expectedErr := errors.New("Invalid start of equation")

	if result != 0 {
		t.Errorf("When processing an invalid string, got: %v, want: %v.", result, nil)
	}

	if err == nil {
		t.Errorf("When processing an invalid string and checking the error, got: %v, want: %v.", err, expectedErr)
	}
}

func TestProcessMissingClose(t *testing.T) {
	result, err := process("(add 1 2")
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
