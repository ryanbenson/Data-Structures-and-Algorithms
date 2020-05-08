package onerow

import "testing"

func TestIsOneRowFilter_HasSome(t *testing.T) {
	words := []string{"candy", "doodle", "pop", "shield", "lag", "typewriter"}
	result := isOneRowFilter(words)
	expected := []string{"pop", "lag", "typewriter"}

	if len(result) != len(expected) {
		t.Errorf("When testing a list of words that has some valid is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsOneRowFilter_BadCharIncluded(t *testing.T) {
	words := []string{"!candy", "dood...le", "po()p", "shield", "lag", "typewriter"}
	result := isOneRowFilter(words)
	expected := []string{"lag", "typewriter"}

	if len(result) != len(expected) {
		t.Errorf("When testing a list of words that has some valid with bad chars is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsOneRowFilter_Empty(t *testing.T) {
	words := []string{}
	result := isOneRowFilter(words)
	expected := []string{}

	if len(result) != len(expected) {
		t.Errorf("When testing a list of words that is empty is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsWordOneRow_Valid(t *testing.T) {
	word := "typewriter"
	result := isWordOneRow(word)
	expected := true

	if result != expected {
		t.Errorf("When testing a word that is valid is incorrect, got %v, want: %v", result, expected)
	}
}

func TestIsWordOneRow_Invalid(t *testing.T) {
	word := "hello"
	result := isWordOneRow(word)
	expected := false

	if result != expected {
		t.Errorf("When testing a word that is valid is incorrect, got %v, want: %v", result, expected)
	}
}

func BenchmarkIsOneRowFilter(b *testing.B) {
    for i := 0; i < b.N; i++ {
		isOneRowFilter([]string{"candy", "doodle", "pop", "shield", "lag", "typewriter"})
	}
}

func BenchmarkIsWordOneRow(b *testing.B) {
    for i := 0; i < b.N; i++ {
		isWordOneRow("typewriter")
	}
}