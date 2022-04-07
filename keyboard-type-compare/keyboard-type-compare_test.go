package keyboardtype

import "testing"

// is golden ratio tests
func TestIsEqual(t *testing.T) {
	result := isEqual("a##x", "#a#x")
	expected := true

	if result != expected {
		t.Errorf("When two keyboard inputs are given, and should be equal: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func BenchmarkIsGoldenRatioArray(b *testing.B) {
    for i := 0; i < b.N; i++ {
		isEqual("a##x", "#a#x")
	}
}