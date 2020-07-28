package isanagram

import (
	"testing"
)

func TestisAnagram_Success(t *testing.T) {
	result := isAnagram("fried", "fired")
	expected := true

	if result != expected {
		t.Errorf("When determing anagram, incorrect result given, got: %v, expected: %v", result, expected)
	}
}
