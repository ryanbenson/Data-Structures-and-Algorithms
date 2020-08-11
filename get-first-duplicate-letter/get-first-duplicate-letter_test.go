package getfirstduplicateletter

import "testing"

func TestGetFirstDuplicateLetter_Map_Success(t *testing.T) {
	result := getFirstDuplicateLetter_Map("hello")
	expected := "l"

	if result != expected {
		t.Errorf("When the first duplicate letter, incorrect result given, got: %v, expected: %v", result, expected)
	}

	result2 := getFirstDuplicateLetter_Map("abcdefabc")
	expected2 := "a"

	if result2 != expected2 {
		t.Errorf("When the first duplicate letter (2), incorrect result given, got: %v, expected: %v", result2, expected2)
	}
}

func TestGetFirstDuplicateLetter_Search_Success(t *testing.T) {
	result := getFirstDuplicateLetter_Search("hello")
	expected := "l"

	if result != expected {
		t.Errorf("When the first duplicate letter, incorrect result given, got: %v, expected: %v", result, expected)
	}

	result2 := getFirstDuplicateLetter_Search("abcdefabc")
	expected2 := "a"

	if result2 != expected2 {
		t.Errorf("When the first duplicate letter (2), incorrect result given, got: %v, expected: %v", result2, expected2)
	}
}

func BenchmarkGetFirstDuplicateLetter_Map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFirstDuplicateLetter_Map("abcdefabc")
	}
}

func BenchmarkGetFirstDuplicateLetter_Search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFirstDuplicateLetter_Search("abcdefabc")
	}
}
