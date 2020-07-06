package binaryqueue

import (
	"testing"
)

func TestBinaryQueue_Success(t *testing.T) {
	num := 10
	result := binaryQueue(num)
	expected := []int{1, 10, 11, 100, 101, 110, 111, 1000, 1001, 1010, 1011, 1100, 1101, 1110, 1111, 10000}

	for _, resultVal := range result {
		found := false
		for _, expectedVal := range expected {
			if resultVal == expectedVal {
				found = true
			}
		}
		// anything not found? something went wrong
		if found == false {
			t.Errorf("When getting a binary queue, incorrect queue given, got: %v, expected: %v", result, expected)
		}
	}
}
