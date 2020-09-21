package removeleadingzeroes

import "testing"

func TestRemove(t *testing.T) {
	result := remove([]int{0, 0, 0, 1, 0, 2, 3})
	expected := []int{1, 0, 2, 3}

	if len(result) != len(expected) {
		t.Errorf("When removing leading zeroes, incorrect length given, got: %v, expected: %v", len(result), len(expected))
		return
	}

	allMatched := true
	for i, val := range expected {
		// if any fail, then it will error out
		if val != result[i] {
			allMatched = false
		}
	}
	// anything not found? something went wrong
	if allMatched == false {
		t.Errorf("When removing leading zeroes, incorrect values given, got: %v, expected: %v", result, expected)
		return
	}
}
