package shortesttimeinterval

import (
	"testing"
)

func TestSmallestTimeInterval(t *testing.T) {
	times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05"}
	result := smallestTimeInterval(times, 60)
	expected := "25 minutes"

	if result != expected {
		t.Errorf("When finding the shortest time interval between a set of times, incorrect result given, got: %v, expected: %v", result, expected)
	}
}


func TestSmallestTimeIntervalLongIntervalNotInOrder(t *testing.T) {
	times := []string{"20:05", "01:00", "11:30", "08:15", "13:45", "23:10"}
	result := smallestTimeInterval(times, 60)
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
		smallestTimeInterval(times, 60)
	}
}

func BenchmarkSmallestTimeIntervalMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05", "08:14", "11:32", "13:47", "14:18", "20:53", "08:43", "11:21", "13:12", "14:43", "20:33", "08:33", "11:31", "13:11", "14:43", "20:54", "08:43", "11:12", "13:43", "14:23", "20:34", "08:43", "11:52", "13:11", "14:43", "20:51"}
		smallestTimeInterval(times, 60)
	}
}

func BenchmarkSmallestTimeIntervalFifteenMinIncr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		times := []string{"01:00", "08:15", "11:30", "13:45", "14:15", "20:00"}
		smallestTimeInterval(times, 900)
	}
}

func BenchmarkSmallestTimeIntervalFiveMinIncr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05"}
		smallestTimeInterval(times, 300)
	}
}

func BenchmarkSmallestTimeIntervalFiveMinIncrMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05", "08:05", "11:50", "13:50", "14:20", "20:15", "08:55", "11:35", "13:35", "14:20", "20:10", "08:00", "11:05", "13:55", "14:25", "20:20", "08:45", "11:05", "13:55", "14:15", "20:15", "08:00", "11:00", "13:00", "14:55", "20:25"}
		smallestTimeInterval(times, 300)
	}
}

func BenchmarkSmallestTimeIntervalSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05"}
		smallestTimeIntervalSimple(times)
	}
}

func BenchmarkSmallestTimeIntervalSimpleMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05", "08:14", "11:32", "13:47", "14:18", "20:53", "08:43", "11:21", "13:12", "14:43", "20:33", "08:33", "11:31", "13:11", "14:43", "20:54", "08:43", "11:12", "13:43", "14:23", "20:34", "08:43", "11:52", "13:11", "14:43", "20:51"}
		smallestTimeIntervalSimple(times)
	}
}
