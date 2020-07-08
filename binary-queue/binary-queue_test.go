package binaryqueue

import (
	"testing"
)

func TestBinaryQueue_Success(t *testing.T) {
	num := 16
	result := binaryQueue(num)
	expected := []string{"1", "10", "11", "100", "101", "110", "111", "1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111", "10000"}

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
func TestBinaryQueue_Zero(t *testing.T) {
	num := 0
	result := binaryQueue(num)
	expected := []string{}

	// got anything? if so, it should be invalid
	if len(result) > 0 {
		t.Errorf("When getting a binary queue, incorrect queue given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkBinaryQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryQueue(16)
	}
}
