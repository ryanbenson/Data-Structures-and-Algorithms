package movezeroes

import (
	"testing"
)

func TestMoveZeroes_Success(t *testing.T) {
	result := moveZeroes([]int{1, 2, 3, 0, 4, 5, 0, 6})
	expected := []int{1, 2, 3, 4, 5, 6, 0, 0}

	allMatched := true
	for i, expectedVal := range expected {
		resultVal := result[i]
		// if any fail, then it will error out
		if expectedVal != resultVal {
			allMatched = false
		}
	}
	// anything not found? something went wrong
	if allMatched == false {
		t.Errorf("When moving zeroes on an array, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func TestMoveZeroesMerge_Success(t *testing.T) {
	result := moveZeroesMerge([]int{1, 2, 3, 0, 4, 5, 0, 6})
	expected := []int{1, 2, 3, 4, 5, 6, 0, 0}

	allMatched := true
	for i, expectedVal := range expected {
		resultVal := result[i]
		// if any fail, then it will error out
		if expectedVal != resultVal {
			allMatched = false
		}
	}
	// anything not found? something went wrong
	if allMatched == false {
		t.Errorf("When moving zeroes on an array, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moveZeroes([]int{1, 2, 3, 0, 4, 5, 0, 6})
	}
}

func BenchmarkMoveZeroesMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moveZeroesMerge([]int{1, 2, 3, 0, 4, 5, 0, 6})
	}
}
