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

func TestGetMedianTwoArraysOptimized_oddNumberOfTotalValuesRightSide(t *testing.T) {
	results := getMedianTwoArraysOptimized([]int{9, 99, 999, 9999, 99999}, []int{1, 1, 1, 1, 1, 2, 3, 4, 5, 7})
	var expected float64
	expected = 4

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}

// Benchmarking the unoptimized and optimized
func BenchmarkGetMedianTwoArrays(b *testing.B) {
    for i := 0; i < b.N; i++ {
        getMedianTwoArrays([]int{1, 2, 3, 4, 5, 9}, []int{1, 2, 3, 4, 5, 7})
    }
}

func BenchmarkGetMedianTwoArraysOptimized(b *testing.B) {
    for i := 0; i < b.N; i++ {
        getMedianTwoArraysOptimized([]int{1, 2, 3, 4, 5, 9}, []int{1, 2, 3, 4, 5, 7})
    }
}
