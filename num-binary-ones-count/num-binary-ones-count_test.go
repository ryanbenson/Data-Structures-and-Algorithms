package numbinaryonescount

import "testing"

func TestGetBinaryOnesCount(t *testing.T) {
	result := getBinaryOnesCount(13)
	// 13 = 1101
	expected := 3

	if result != expected {
		t.Errorf("When getting count of binary 1s from a number, got %v, want, %v", result, expected)
	}
}

func TestGetBinaryOnesCount_LargeNumber(t *testing.T) {
	result := getBinaryOnesCount(123456789)
	// 123456789 = 111010110111100110100010101
	expected := 16

	if result != expected {
		t.Errorf("When getting count of binary 1s from a number, that is a large number, got %v, want, %v", result, expected)
	}
}

func TestGetBinaryOnesCount_Zero(t *testing.T) {
	result := getBinaryOnesCount(0)
	// 0 = 0
	expected := 0

	if result != expected {
		t.Errorf("When getting count of binary 1s from a number, that is zero, got %v, want, %v", result, expected)
	}
}