package longestsubsequence

import (
	"testing"
)

func TestLongestSubSeq(t *testing.T) {
	arr := []int{1, 9, 87, 3, 10, 4, 20, 2, 45}
	result := longestSubSeq(arr)
	expected := 5

	if result != expected {
		t.Errorf("When finding the largest subsequence, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkLargestRect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1, 9, 87, 3, 10, 4, 20, 2, 45}
		_ = longestSubSeq(arr)
	}
}
