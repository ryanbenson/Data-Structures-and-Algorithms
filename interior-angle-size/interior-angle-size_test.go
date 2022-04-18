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

func TestGetInteriorAngleBelowMin(t *testing.T) {
	result0 := getInteriorAngle(0)
	result1 := getInteriorAngle(1)
	result2 := getInteriorAngle(2)
	expected := 0

	if result0 != expected {
		t.Errorf("When getting the bigger interior angle size, but too small: is incorrect, got: %v, want: %v.", result0, expected)
	}

	if result1 != expected {
		t.Errorf("When getting the bigger interior angle size, but too small: is incorrect, got: %v, want: %v.", result1, expected)
	}

	if result2 != expected {
		t.Errorf("When getting the bigger interior angle size, but too small: is incorrect, got: %v, want: %v.", result2, expected)
	}
}

func BenchmarkIsEqualSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getInteriorAngle(8)
	}
}