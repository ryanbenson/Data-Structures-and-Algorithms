package isgoldenratio

import "testing"

func TestIsGoldenRatio_ValidOne(t *testing.T) {
	result := isGoldenRatio(19.416, 12)
	expected := true

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatio_InvalidOne(t *testing.T) {
	result := isGoldenRatio(1, 2)
	expected := false

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatio_InvalidTwo(t *testing.T) {
	result := isGoldenRatio(2, 5)
	expected := false

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatio_ValidTwo(t *testing.T) {
	result := isGoldenRatio(1, 0.618)
	expected := true

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatio_ValidThree(t *testing.T) {
	result := isGoldenRatio(61.77, 38.176)
	expected := true

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func BenchmarkIsGoldenRatio(b *testing.B) {
    for i := 0; i < b.N; i++ {
		isGoldenRatio(19.416, 12)
	}
}