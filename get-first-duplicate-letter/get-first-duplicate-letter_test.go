package getfirstduplicateletter

import "testing"

func TestGetFirstDuplicateLetter_Success(t *testing.T) {
	result := getFirstDuplicateLetter("hello")
	expected := "l"

	if result != expected {
		t.Errorf("When the first duplicate letter, incorrect result given, got: %v, expected: %v", result, expected)
	}
}
