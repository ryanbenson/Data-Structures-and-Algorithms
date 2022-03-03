package largestisland

import (
	"testing"
)

func TestLargestRect(t *testing.T) {
	arr := [][]int{
		{0,0,0,1,0},
		{1,1,0,0,1},
		{1,1,0,0,0},
		{0,0,1,0,0},
	}
	result := largestRect(arr)

	expected := "2x2"
	// anything not found? something went wrong
	if result != expected {
		t.Errorf("When finding the largest rectangle, incorrect rectangle given, got: %v, expected: %v", result, expected)
	}
}

func TestLargestRect_Multiple(t *testing.T) {
	arr := [][]int{
		{0,0,0,1,0,1,1,1},
		{1,1,0,0,0,1,1,1},
		{1,1,0,0,0,1,1,1},
		{0,0,1,0,0,1,1,1},
		{0,0,1,0,0,1,1,1},
		{0,0,1,0,0,0,0,0},
		{0,0,1,0,0,0,0,0},
	}
	result := largestRect(arr)

	expected := "3x5"
	// anything not found? something went wrong
	if result != expected {
		t.Errorf("When finding the largest rectangle, incorrect rectangle given, got: %v, expected: %v", result, expected)
	}
}

func TestLargestRect_NotSquareMap(t *testing.T) {
	arr := [][]int{
		{1,1,0,0,0,1,1,1},
		{1,1,0,0,0,1,1,1},
		{0,0,1,0,0,1,1,1},
		{0,0,1,0,0,0,0,0},
	}
	result := largestRect(arr)

	expected := "3x3"
	// anything not found? something went wrong
	if result != expected {
		t.Errorf("When finding the largest rectangle, incorrect rectangle given, got: %v, expected: %v", result, expected)
	}
}

func TestLargestRect_WeirdShapes(t *testing.T) {
	arr := [][]int{
		{1,1,0,0,0,1,1,1},
		{1,1,1,0,0,1,1,1},
		{0,0,1,0,1,1,1,1},
		{0,0,1,0,1,1,1,1},
	}
	result := largestRect(arr)

	expected := "3x4"
	// anything not found? something went wrong
	if result != expected {
		t.Errorf("When finding the largest rectangle, incorrect rectangle given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkLargestRect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := [][]int{
			{0,0,0,1,0},
			{1,1,0,0,1},
			{1,1,0,0,0},
			{0,0,1,0,0},
		}
		_ = largestRect(arr)
	}
}