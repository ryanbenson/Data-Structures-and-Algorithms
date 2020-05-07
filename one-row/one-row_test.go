package onerow

import "testing"

func TestIsOneRowFilter_HasSome(t *testing.T) {
	words := []string{"candy", "doodle", "pop", "shield", "lag", "typewriter"}
	result := isOneRowFilter(words)
	expected := []string{"pop", "lag", "typewriter"}

	if len(result) != len(expected) {
		t.Errorf("When testing a list of words that is is incorrect, got: %v, want: %v.", result, expected)
	}
}
