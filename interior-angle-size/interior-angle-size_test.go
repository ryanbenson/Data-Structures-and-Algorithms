package interioranglesize

import "testing"

func TestGetInteriorAngle(t *testing.T) {
	result := getInteriorAngle(3)
	expected := 60

	if result != expected {
		t.Errorf("When getting the interior angle size: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestGetInteriorAngleHigh(t *testing.T) {
	result := getInteriorAngle(8)
	expected := 135

	if result != expected {
		t.Errorf("When getting the bigger interior angle size: is incorrect, got: %v, want: %v.", result, expected)
	}
}

// func BenchmarkIsEqualSimple(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		isEqual("a##x", "#a#x")
// 	}
// }