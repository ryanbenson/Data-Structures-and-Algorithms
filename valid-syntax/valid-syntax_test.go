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

// Rune vesion tests
func TestIsValidRune_SimpleOpenClose(t *testing.T) {
	result := isValidRune("()")
	expected := true

	if result != expected {
		t.Errorf("isValidRune boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidRune_MultipleChainedOpenClose(t *testing.T) {
	result := isValidRune("()[]{}")
	expected := true

	if result != expected {
		t.Errorf("isValidRune boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidRune_Mismatched(t *testing.T) {
	result := isValidRune("(]")
	expected := false

	if result != expected {
		t.Errorf("isValidRune boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidRune_NestedMismatched(t *testing.T) {
	result := isValidRune("([)]")
	expected := false

	if result != expected {
		t.Errorf("isValidRune boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidRune_NestedProperly(t *testing.T) {
	result := isValidRune("{[]}")
	expected := true

	if result != expected {
		t.Errorf("isValidRune boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidRune_Empty(t *testing.T) {
	result := isValidRune("")
	expected := true

	if result != expected {
		t.Errorf("isValidRune boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

// Optimized version tests
func TestIsValidMagic_SimpleOpenClose(t *testing.T) {
	result := isValidMagic("()")
	expected := true

	if result != expected {
		t.Errorf("isValidOptimized boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidMagic_MultipleChainedOpenClose(t *testing.T) {
	result := isValidMagic("()[]{}")
	expected := true

	if result != expected {
		t.Errorf("isValidOptimized boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidMagic_Mismatched(t *testing.T) {
	result := isValidMagic("(]")
	expected := false

	if result != expected {
		t.Errorf("isValidOptimized boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidMagic_NestedMismatched(t *testing.T) {
	result := isValidMagic("([)]")
	expected := false

	if result != expected {
		t.Errorf("isValidOptimized boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidMagic_NestedProperly(t *testing.T) {
	result := isValidMagic("{[]}")
	expected := true

	if result != expected {
		t.Errorf("isValidOptimized boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsValidMagic_Empty(t *testing.T) {
	result := isValidMagic("")
	expected := true

	if result != expected {
		t.Errorf("isValidOptimized boolean is incorrect, got: %v, want: %v.", result, expected)
	}
}


// Benchmarking the unoptimized and optimized
func BenchmarkValidSyntax(b *testing.B) {
    for i := 0; i < b.N; i++ {
        isValid("{[]}")
    }
}

func BenchmarkValidSyntaxRune(b *testing.B) {
    for i := 0; i < b.N; i++ {
        isValidRune("{[]}")
    }
}

func BenchmarkValidSyntaxMagic(b *testing.B) {
    for i := 0; i < b.N; i++ {
        isValidMagic("{[]}")
    }
}
