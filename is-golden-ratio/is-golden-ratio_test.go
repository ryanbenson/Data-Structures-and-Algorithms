package isgoldenratio

import "testing"

func TestIsGoldenRatio_ValidOne(t *testing.T) {
	result := isGoldenRatio(19.416, 12)
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