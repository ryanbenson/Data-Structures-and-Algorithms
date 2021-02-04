package customcustoms

import (
	"testing"
)

func TestSum(t *testing.T) {
	expected := 0
	result := sum("hello")

	if result != expected {
		t.Errorf("When getting custom customs, looking at sum, got: %v, want: %v.", result, expected)
	}
}
