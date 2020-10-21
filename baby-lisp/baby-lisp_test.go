package babylisp

import (
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
