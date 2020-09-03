package getfirstduplicateletter

import "testing"

func TestGetFirstDuplicateLetterMap_Success(t *testing.T) {
	result := getFirstDuplicateLetterMap("hello")
	expected := "l"

	if result != expected {
		t.Errorf("When the first duplicate letter, incorrect result given, got: %v, expected: %v", result, expected)
	}

	result2 := getFirstDuplicateLetterMap("abcdefabc")
	expected2 := "a"

	if result2 != expected2 {
		t.Errorf("When the first duplicate letter (2), incorrect result given, got: %v, expected: %v", result2, expected2)
	}
}

func TestGetFirstDuplicateLetterMap_NoDuplicates(t *testing.T) {
	result := getFirstDuplicateLetterMap("asdfghjkl")
	expected := ""

	if result != expected {
		t.Errorf("When the first duplicate letter, and no dupliate letters, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func TestGetFirstDuplicateLetterSearch_Success(t *testing.T) {
	result := getFirstDuplicateLetterSearch("hello")
	expected := "l"

	if result != expected {
		t.Errorf("When the first duplicate letter, incorrect result given, got: %v, expected: %v", result, expected)
	}

	result2 := getFirstDuplicateLetterSearch("abcdefabc")
	expected2 := "a"

	if result2 != expected2 {
		t.Errorf("When the first duplicate letter (2), incorrect result given, got: %v, expected: %v", result2, expected2)
	}
}

func TestGetFirstDuplicateLetterSearch_NoDuplicates(t *testing.T) {
	result := getFirstDuplicateLetterSearch("asdfghjkl")
	expected := ""

	if result != expected {
		t.Errorf("When the first duplicate letter, and no dupliate letters, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkGetFirstDuplicateLetterMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFirstDuplicateLetterMap("abcdefabc")
	}
}

func BenchmarkGetFirstDuplicateLetterSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFirstDuplicateLetterSearch("abcdefabc")
	}
}
