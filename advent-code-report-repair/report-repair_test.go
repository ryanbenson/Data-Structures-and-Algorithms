package reportrepair

import (
	"testing"
)

func TestProcessExample(t *testing.T) {
	list := []int{1721, 979, 366, 299, 675, 1456}
	result := repair(list)
	expected := 1

	if result != expected {
		t.Errorf("When repairing a report, got: %v, want: %v.", result, expected)
	}
}
