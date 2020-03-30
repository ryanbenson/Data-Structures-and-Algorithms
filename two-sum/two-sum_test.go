package twosum

import "testing"

func TestTwoSum_HasMatches(t *testing.T) {
	results := twoSum([]int{9, 1, 2, 3, 4, 5}, 9)
	expected := []int{4, 5}

	if results[0] != expected[0] {
		t.Errorf("First int is incorrect, got: %d, want: %d.", results[0], expected[0])
	}
	if results[1] != expected[1] {
		t.Errorf("Second int is incorrect, got: %d, want: %d.", results[1], expected[1])
	}
}

func TestTwoSum_HasMatchesLong(t *testing.T) {
	results := twoSum([]int{9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 100, 10000, 100000000, 432423423}, 100000001)
	expected := []int{1, 12}

	if results[0] != expected[0] {
		t.Errorf("First int is incorrect, got: %d, want: %d.", results[0], expected[0])
	}
	if results[1] != expected[1] {
		t.Errorf("Second int is incorrect, got: %d, want: %d.", results[1], expected[1])
	}
}

func TestTwoSum_NoMatches(t *testing.T) {
	results := twoSum([]int{9, 1, 2, 3, 4, 5}, 555)

	if results != nil {
		t.Errorf("No matches is incorrect, got: %d, want: %v.", results, nil)
	}
}
