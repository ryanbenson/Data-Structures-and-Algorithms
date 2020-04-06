package comparewithbackspace

import "testing"

func TestCompareWithBackspace_MatchesWithBackspace(t *testing.T) {
	result := compareWithBackspace("a##c", "#a#c")
	expected := true

	if result != expected {
		t.Errorf("When given matching strings with backspaces input is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestCompareWithBackspace_MatchesNoBackspace(t *testing.T) {
	result := compareWithBackspace("hello there beautiful day is it not?", "hello there beautiful day is it not?")
	expected := true

	if result != expected {
		t.Errorf("When given matching strings with no backspaces input is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestCompareWithBackspace_MatchesAllBackspace(t *testing.T) {
	result := compareWithBackspace("xy##", "z#w#")
	expected := true

	if result != expected {
		t.Errorf("When given not matching strings with backspaces that empty the string input is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestCompareWithBackspace_NoMatchBackspace(t *testing.T) {
	result := compareWithBackspace("a##c", "#a#cdef#z")
	expected := false

	if result != expected {
		t.Errorf("When given not matching strings with backspaces input is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestCompareWithBackspace_NoMatchNoBackspace(t *testing.T) {
	result := compareWithBackspace("Ég elska", "að læra")
	expected := false

	if result != expected {
		t.Errorf("When given not matching strings with no backspaces input is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestBuildString_WithBackspaces(t *testing.T) {
	result := buildString("a##c")
	expected := "c"

	if result != expected {
		t.Errorf("When building strings with backspaces input is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestBuildString_NoBackspace(t *testing.T) {
	result := buildString("Ég elska")
	expected := "Ég elska"

	if result != expected {
		t.Errorf("When building strings no backspaces input is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestBuildString_TooManyBackspace(t *testing.T) {
	result := buildString("Ég elska######c#####z#####fff#######")
	expected := ""

	if result != expected {
		t.Errorf("When building strings with too many backspaces input is incorrect, got: %v, want: %v.", result, expected)
	}
}