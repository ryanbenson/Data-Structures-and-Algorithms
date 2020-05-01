package validsyntax

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

func TestIsValid_Empty(t *testing.T) {
	result := isValid("")
	expected := true

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

// Optimized version tests


func TestIsValidOptimized_SimpleOpenClose(t *testing.T) {
	result := isValidOptimized("()")
	expected := true

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidOptimized_MultipleChainedOpenClose(t *testing.T) {
	result := isValidOptimized("()[]{}")
	expected := true

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidOptimized_Mismatched(t *testing.T) {
	result := isValidOptimized("(]")
	expected := false

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidOptimized_NestedMismatched(t *testing.T) {
	result := isValidOptimized("([)]")
	expected := false

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidOptimized_NestedProperly(t *testing.T) {
	result := isValidOptimized("{[]}")
	expected := true

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidOptimized_Empty(t *testing.T) {
	result := isValidOptimized("")
	expected := true

	if result != expected {
		t.Errorf("isValid boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}


// Benchmarking the unoptimized and optimized
func BenchmarkValidSyntax(b *testing.B) {
    for i := 0; i < b.N; i++ {
        isValid("{[]}")
    }
}

func BenchmarkValidSyntaxAlt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        isValidOptimized("{[]}")
    }
}
