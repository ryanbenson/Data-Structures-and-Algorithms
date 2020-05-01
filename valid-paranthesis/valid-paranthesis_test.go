package validparanthesis

import "testing"

func TestIsValid_SimpleOpenClose(t *testing.T) {
	result := isValid("()")
	expected := true

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValid_MultipleChainedOpenClose(t *testing.T) {
	result := isValid("()[]{}")
	expected := true

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValid_Mismatched(t *testing.T) {
	result := isValid("(]")
	expected := false

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValid_NestedMismatched(t *testing.T) {
	result := isValid("([)]")
	expected := false

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValid_NestedProperly(t *testing.T) {
	result := isValid("{[]}")
	expected := true

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}
