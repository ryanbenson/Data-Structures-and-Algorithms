package binaryqueue

import (
	"testing"
)

func TestBinaryQueue_Success(t *testing.T) {
	num := 10
	result := binaryQueue(num)
	expected := []int{1, 10, 11, 100, 101, 110, 111, 1000, 1001, 1010, 1011, 1100, 1101, 1110, 1111, 10000}

	allMatched := true
	for _, resultVal := range expected {
		found := false
		for _, expectedVal := range result {
			if resultVal == expectedVal {
				found = true
			}
		}
		// if any fail, then it will error out
		if found == false {
			allMatched = false
		}
	}
	// anything not found? something went wrong
	if allMatched == false {
		t.Errorf("When getting a binary queue, incorrect queue given, got: %v, expected: %v", result, expected)
	}
}
