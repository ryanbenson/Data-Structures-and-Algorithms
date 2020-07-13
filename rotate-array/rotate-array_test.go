package rotatearray

import (
	"testing"
)

func TestRotateArray_Success(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := 2
	result, err := rotateArray(arr, index)

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
	if err != nil {
		t.Errorf("When rotating an array, incorrect error given, got: %v, expected: %v", result, nil)
	}
}
func TestRotateArray_SuccessLastIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := 4
	result, err := rotateArray(arr, index)

	expected := []int{1, 2, 3, 4, 5}

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
	if err != nil {
		t.Errorf("When rotating an array, incorrect error given, got: %v, expected: %v", result, nil)
	}
}
func TestRotateArray_IndexExceedsBounds(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := 90
	result, err := rotateArray(arr, index)

	if result != nil {
		t.Errorf("When rotating an array with a too high index, incorrect queue given, got: %v, expected: %v", result, nil)
	}
	if err == nil {
		t.Errorf("When rotating an array with a too high index, incorrect error given, got: %v, expected: %v", result, nil)
	}
}

// loop flavor implementation

func TestRotateArrayLoop_Success(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := 2
	result, err := rotateArrayLoop(arr, index)

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
	if err != nil {
		t.Errorf("When rotating an array, incorrect error given, got: %v, expected: %v", result, nil)
	}
}
func TestRotateArrayLoop_SuccessLastIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := 4
	result, err := rotateArrayLoop(arr, index)

	expected := []int{1, 2, 3, 4, 5}

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
	if err != nil {
		t.Errorf("When rotating an array, incorrect error given, got: %v, expected: %v", result, nil)
	}
}
func TestRotateArrayLoop_IndexExceedsBounds(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := 90
	result, err := rotateArrayLoop(arr, index)

	if result != nil {
		t.Errorf("When rotating an array with a too high index, incorrect queue given, got: %v, expected: %v", result, nil)
	}
	if err == nil {
		t.Errorf("When rotating an array with a too high index, incorrect error given, got: %v, expected: %v", result, nil)
	}
}

func BenchmarkRotateArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1, 2, 3, 4, 5}
		index := 2
		_, _ = rotateArray(arr, index)
	}
}
func BenchmarkRotateArrayLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1, 2, 3, 4, 5}
		index := 2
		_, _ = rotateArrayLoop(arr, index)
	}
}
