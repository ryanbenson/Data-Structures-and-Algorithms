package keyboardtype

import "testing"

func TestIsEqual(t *testing.T) {
	result := isEqual("a##x", "#a#x")
	expected := true

	if result != expected {
		t.Errorf("When two keyboard inputs are given, and should be equal: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsEqualWillNoBeEqual(t *testing.T) {
	result := isEqual("fi##f%%%th %%year #time###", "fifth year time")
	expected := false

	if result != expected {
		t.Errorf("When two keyboard inputs are given, and should be not equal: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func BenchmarkIsEqualSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEqual("a##x", "#a#x")
	}
}

func BenchmarkComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEqual("fi##f%%%th %%year #time###", "fifth year time")
	}
}