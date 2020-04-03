package meanmedianmode

import "testing"

func TestGetMeanMedianMode_evenNumberOfValues(t *testing.T) {
	results := getMeanMedianMode([]int{9, 1, 2, 3, 4, 5})
	expected := meanMedianMode{
		mean:   4,
		median: 3,
		mode:   nil,
	}

	if results.mean != expected.mean {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results.mean, expected.mean)
	}
	if results.median != expected.median {
		t.Errorf("Median int is incorrect, got: %v, want: %v.", results.median, expected.median)
	}
	if results.mode != nil {
		t.Errorf("Mode int is incorrect, got: %v, want: %v.", results.mode, nil)
	}
}

func TestGetMeanMedianMode_oddNumberOfValues(t *testing.T) {
	results := getMeanMedianMode([]int{23, 12, 786, 7, 35, 12, 9})
	expected := meanMedianMode{
		mean:   126,
		median: 12,
		mode:   []int{12},
	}

	if results.mean != expected.mean {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results.mean, expected.mean)
	}
	if results.median != expected.median {
		t.Errorf("Median is incorrect, got: %v, want: %v.", results.median, expected.median)
	}
	if results.mode[0] != expected.mode[0] {
		t.Errorf("Mode is incorrect, got: %v, want: %v.", results.mode, expected.mode[0])
	}
}

func TestMean(t *testing.T) {
	results := mean([]int{23, 12, 786, 7, 35, 12, 9})
	var expected float64
	expected = 126

	if results != expected {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestMedian_EvenNumberItems(t *testing.T) {
	results := median([]int{5, 7, 99, 5, 12, 34, 19, 357})
	var expected float64
	expected = 8

	if results != expected {
		t.Errorf("Median is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestMedian_OddNumberItems(t *testing.T) {
	results := median([]int{5, 7, 99, 5, 12, 34, 19})
	var expected float64
	expected = 5

	if results != expected {
		t.Errorf("Median is incorrect, got: %v, want: %v.", results, expected)
	}
}

func TestMode(t *testing.T) {
	results := mode([]int{1, 5, 7, 9, 12, 5, 1, 9, 9, 5, 7, 9, 9})
	expected := []int{9}

	if results[0] != expected[0] {
		t.Errorf("Mode is incorrect, got: %v, want: %v.", results, expected)
	}
}
