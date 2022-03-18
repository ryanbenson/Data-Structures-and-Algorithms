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
		t.Errorf("When finding if all plants can be planted, incorrect result given, got: %v, expected: %v", result, expected)
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
		t.Errorf("When finding if all plants can be planted, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkCanPlantLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		size := 3
		direction := "left"
		printArrow(direction, size)
	}
}
