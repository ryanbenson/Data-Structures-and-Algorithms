package rotatearray

import (
	"bytes"
	"testing"
)

func TestRotateArray_Success(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := 2
	result := rotateArray(arr, index)

	expected := []string{4, 5, 1, 2, 3}

	matched := bytes.Compare(result, expected)
	// if matched == 0, they match, otherwise they failed to match
	if matched != 0 {
		t.Errorf("When rotating an array, incorrect queue given, got: %v, expected: %v", result, expected)
	}
}
