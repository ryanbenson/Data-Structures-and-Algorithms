package largestsubarraysum

import (
	"testing"
)

// func TestGetSum(t *testing.T) {
// 	arr := []int{1,2,3,4,5,6}
// 	result, err := getSum(arr, 3)

// 	expected := []int{4,5,6}
// 	if err != nil {
// 		t.Errorf("When finding the largest subset array, incorrect error given, got: %v, expected: %v", err, expected)
// 	}
// 	if result != expected {
// 		t.Errorf("When finding the largest subset array, incorrect values given, got: %v, expected: %v", result, expected)
// 	}
// }

func TestGetSumErr(t *testing.T) {
	arr := []int{1,2,3,4,5,6}
	result, err := getSum(arr, 10)

	if err == nil {
		t.Errorf("When finding the largest subset array, incorrect error given, got: %v", err)
	}
	if result != nil {
		t.Errorf("When finding the largest subset array, incorrect values given, got: %v, expected: %v", result, nil)
	}
}

func BenchmarkLargestRect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1,2,3,4,5,6}
		getSum(arr, 10)
	}
}