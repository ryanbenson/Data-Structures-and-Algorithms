package canplant

import (
	"testing"
)

func TestCanPlant(t *testing.T) {
	garden := []int{1, 0, 0, 0, 1}
	plantsDesired := 1
	result := canPlant(garden, plantsDesired)
	expected := true

	if result != expected {
		t.Errorf("When finding if all plants can be planted, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func TestCanPlantTooManyDesired(t *testing.T) {
	garden := []int{1, 0, 0, 0, 1}
	plantsDesired := 4
	result := canPlant(garden, plantsDesired)
	expected := false

	if result != expected {
		t.Errorf("When finding if all plants can be planted when there are too many, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func TestCanPlantLong(t *testing.T) {
	garden := []int{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1}
	plantsDesired := 4
	result := canPlant(garden, plantsDesired)
	expected := true

	if result != expected {
		t.Errorf("When finding if lots of plants can be planted, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkCanPlantLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		garden := []int{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1}
		plantsDesired := 4
		canPlant(garden, plantsDesired)
	}
}
