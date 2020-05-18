package isgoldenratio

import "testing"

// is golden ratio tests
func TestIsGoldenRatio_ValidOne(t *testing.T) {
	result := isGoldenRatio(19.416, 12)
	expected := true

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatio_InvalidOne(t *testing.T) {
	result := isGoldenRatio(1, 2)
	expected := false

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatio_InvalidTwo(t *testing.T) {
	result := isGoldenRatio(2, 5)
	expected := false

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatio_ValidTwo(t *testing.T) {
	result := isGoldenRatio(1, 0.618)
	expected := true

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatio_ValidThree(t *testing.T) {
	result := isGoldenRatio(61.77, 38.176)
	expected := true

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

// is golden ratio (array) tests
func TestIsGoldenRatioArray_ValidOne(t *testing.T) {
	result := isGoldenRatioArray(19.416, 12)
	expected := true

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatioArray_InvalidOne(t *testing.T) {
	result := isGoldenRatioArray(1, 2)
	expected := false

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatioArray_InvalidTwo(t *testing.T) {
	result := isGoldenRatioArray(2, 5)
	expected := false

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatioArray_ValidTwo(t *testing.T) {
	result := isGoldenRatioArray(1, 0.618)
	expected := true

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsGoldenRatioArray_ValidThree(t *testing.T) {
	result := isGoldenRatioArray(61.77, 38.176)
	expected := true

	if result != expected {
		t.Errorf("When determining phi: is incorrect, got: %v, want: %v.", result, expected)
	}
}

// floating remainder tests
func TestFloatRemainder_NoRemainder(t *testing.T) {
	result := floatRemainder(1.0)
	expected := 1.000

	if result != expected {
		t.Errorf("When rounding float with remainder: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestFloatRemainder_RoundDown(t *testing.T) {
	result := floatRemainder(13.1114999)
	expected := 13.111

	if result != expected {
		t.Errorf("When rounding float with remainder: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestFloatRemainder_RoundUpDecimal(t *testing.T) {
	result := floatRemainder(13.444999)
	expected := 13.445

	if result != expected {
		t.Errorf("When rounding float with remainder: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestFloatRemainder_RoundUpEnd(t *testing.T) {
	result := floatRemainder(13.999999)
	expected := 14.000

	if result != expected {
		t.Errorf("When rounding float with remainder: is incorrect, got: %v, want: %v.", result, expected)
	}
}

func BenchmarkIsGoldenRatio(b *testing.B) {
    for i := 0; i < b.N; i++ {
		isGoldenRatio(19.416, 12)
	}
}

func BenchmarkIsGoldenRatioArray(b *testing.B) {
    for i := 0; i < b.N; i++ {
		isGoldenRatioArray(19.416, 12)
	}
}