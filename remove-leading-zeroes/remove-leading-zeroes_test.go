package removeleadingzeroes

import (
	"reflect"
	"testing"
)

func TestRemoveSearchList(t *testing.T) {
	result := removeSearchList([]int{0, 0, 0, 1, 0, 2, 3})
	expected := []int{1, 0, 2, 3}

	if len(result) != len(expected) {
		t.Errorf("When removing leading zeroes, incorrect length given, got: %v, expected: %v", len(result), len(expected))
		return
	}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("When removing leading zeroes, incorrect values given, got: %v, expected: %v", result, expected)
		return
	}
}

func BenchmarkRemoveSearchList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeSearchList([]int{0, 0, 0, 1, 0, 2, 3})
	}
}
