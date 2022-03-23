package shortesttimeinterval

import (
	"testing"
)

func TestSmallestTimeInterval(t *testing.T) {
	times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05"}
	result := smallestTimeInterval(times)
	expected := "25 minutes"

	if result != expected {
		t.Errorf("When finding the shortest time interval between a set of times, incorrect result given, got: %v, expected: %v", result, expected)
	}
}


func TestSmallestTimeIntervalLongIntervalNotInOrder(t *testing.T) {
	times := []string{"20:05", "01:00", "11:30", "08:15", "13:45", "23:10"}
	result := smallestTimeInterval(times)
	expected := "2 hours, 15 minutes"

	if result != expected {
		t.Errorf("When finding the shortest time interval between a set of times, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func TestSmallestTimeIntervalSimple(t *testing.T) {
	times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05"}
	result := smallestTimeIntervalSimple(times)
	expected := "25 minutes"

	if result != expected {
		t.Errorf("When finding the shortest time interval between a set of times, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkSmallestTimeInterval(b *testing.B) {
	for i := 0; i < b.N; i++ {
		times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05"}
		smallestTimeInterval(times)
	}
}

func BenchmarkSmallestTimeIntervalSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05"}
		smallestTimeIntervalSimple(times)
	}
}
