package generatearrow

import (
	"testing"
)

func TestPrintArrowRight(t *testing.T) {
	size := 3
	direction := "right"
	result := printArrow(direction, size)
	expected := `*
  *
    *
  *
*`

	if result != expected {
		t.Errorf("When generating an arrow to the right, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func TestPrintArrowLeft(t *testing.T) {
	size := 3
	direction := "left"
	result := printArrow(direction, size)
	expected := `    *
  *
*
  *
    *`

	if result != expected {
		t.Errorf("When generating an arrow to the left, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func TestPrintArrowEmpty(t *testing.T) {
	size := 0
	direction := "left"
	result := printArrow(direction, size)
	expected := ""

	if result != expected {
		t.Errorf("When generating an empty arrow, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func TestPrintArrowSingle(t *testing.T) {
	size := 1
	direction := "left"
	result := printArrow(direction, size)
	expected := "*"

	if result != expected {
		t.Errorf("When generating an single arrow, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkCanPlantLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		size := 3
		direction := "left"
		printArrow(direction, size)
	}
}
