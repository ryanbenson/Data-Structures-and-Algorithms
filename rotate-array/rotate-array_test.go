package rotatearray

import (
	"testing"
)

func TestRotateArray_Success(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := 2
	result := rotateArray(arr, index)

	expected := []int{4, 5, 1, 2, 3}

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
		t.Errorf("When rotating an array, incorrect queue given, got: %v, expected: %v", result, expected)
	}
}
