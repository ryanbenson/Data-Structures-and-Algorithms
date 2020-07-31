package isanagram

import (
	"testing"
)

func TestIsAnagram_Success(t *testing.T) {
	result1 := isAnagram("fried", "fired")
	result2 := isAnagram("listen", "silent")
	expected := true

	if result1 != expected {
		t.Errorf("When determing anagram 1, incorrect result given, got: %v, expected: %v", result1, expected)
	}

	if result2 != expected {
		t.Errorf("When determing anagram 2, incorrect result given, got: %v, expected: %v", result2, expected)
	}
}

func TestIsAnagram_Fail(t *testing.T) {
	result1 := isAnagram("bob", "boo")
	result2 := isAnagram("hello", "help")
	expected := false

	if result1 != expected {
		t.Errorf("When determing anagram 1, incorrect result given, got: %v, expected: %v", result1, expected)
	}

	if result2 != expected {
		t.Errorf("When determing anagram 2, incorrect result given, got: %v, expected: %v", result2, expected)
	}
}

func TestSortString(t *testing.T) {
	result := sortString("hello")
	expected := "ehllo"

	if result != expected {
		t.Errorf("When sorting string, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkIsAnagram_SortStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram("listen", "silent")
	}
}
