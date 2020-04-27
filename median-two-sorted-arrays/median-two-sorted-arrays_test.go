package mediantwosortedarrays

import (
	"testing"
)

func TestGetMedianTwoArrays_evenNumberOfTotalValues(t *testing.T) {
	results := getMedianTwoArrays([]int{1, 2, 3, 4, 5, 9}, []int{1, 2, 3, 4, 5, 7})
	var expected float64
	expected = 3

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestGetMedianTwoArrays_oddNumberOfTotalValues(t *testing.T) {
	results := getMedianTwoArrays([]int{1, 2, 3, 4, 5, 9, 11}, []int{1, 2, 3, 4, 5, 7})
	var expected float64
	expected = 4

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestGetMedianTwoArrays_oneValue(t *testing.T) {
	results := getMedianTwoArrays([]int{7}, []int{})
	var expected float64
	expected = 7

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestGetMedianTwoArraysOptimized_evenNumberOfTotalValues(t *testing.T) {
	results := getMedianTwoArraysOptimized([]int{1, 2, 3, 4, 5, 9}, []int{1, 2, 3, 4, 5, 7})
	var expected float64
	expected = 3

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestGetMedianTwoArraysOptimized_oddNumberOfTotalValues(t *testing.T) {
	results := getMedianTwoArraysOptimized([]int{1, 2, 3, 4, 5, 9, 11}, []int{1, 2, 3, 4, 5, 7})
	var expected float64
	expected = 4

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestGetMedianTwoArraysOptimized_oneValue(t *testing.T) {
	results := getMedianTwoArraysOptimized([]int{7}, []int{})
	var expected float64
	expected = 7

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}
