package babylisp

import (
	"testing"
)

func TestProcess(t *testing.T) {
	result := process("(add 1 2)")
	expected := 3

	if result != expected {
		t.Errorf("When processing a calculation, got: %v, want: %v.", result, expected)
	}
}
