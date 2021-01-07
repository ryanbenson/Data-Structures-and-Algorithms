package binaryboarding

import (
	"testing"
)

func TestGetMaxIdExample(t *testing.T) {
	passes := `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`
	result := getMaxId(passes)
	expected := 820

	if result != expected {
		t.Errorf("When processing boarding passes, looking for highest number, got: %v, want: %v.", result, expected)
	}
}
