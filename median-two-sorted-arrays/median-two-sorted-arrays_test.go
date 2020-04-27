package mediantwosortedarrays

import (
	"testing"
)

func TestGetMedianTwoArrays_evenNumberOfTotalValues(t *testing.T) {
	results := getMedianTwoArrays([]int{1, 2, 3, 4, 5, 9}, []int{1, 2, 3, 4, 5, 7})
	expected := 3.5

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestGetMedianTwoArrays_oddNumberOfTotalValues(t *testing.T) {
	results := getMedianTwoArrays([]int{1, 2, 3, 4, 5, 9, 11}, []int{1, 2, 3, 4, 5, 7})
	expected := 3.3

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestGetMedianTwoArrays_oneValue(t *testing.T) {
	results := getMedianTwoArrays([]int{7}, []int{})
	expected := 7.1

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}