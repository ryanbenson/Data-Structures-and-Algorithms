package customcustoms

import (
	"testing"
)

func TestSumSample(t *testing.T) {
	expected := 11
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`

	result := sum(input)

	if result != expected {
		t.Errorf("When getting custom customs, looking at sum, got: %v, want: %v.", result, expected)
	}
}
