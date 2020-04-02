package meanmedianmode

import "testing"

func TestGetMeanMedianMode_evenNumberOfValues(t *testing.T) {
	results := getMeanMedianMode([]int{9, 1, 2, 3, 4, 5})
	expected := meanMedianMode{
		mean:   4,
		median: 4,
		mode:   1,
	}

	if results.mean != expected.mean {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results.mean, expected.mean)
	}
	if results.median != expected.median {
		t.Errorf("Median int is incorrect, got: %v, want: %v.", results.median, expected.median)
	}
	if results.mode != expected.mode {
		t.Errorf("Mode int is incorrect, got: %d, want: %d.", results.mode, expected.mode)
	}
}

func TestGetMeanMedianMode_oddNumberOfValues(t *testing.T) {
	results := getMeanMedianMode([]int{23, 12, 786, 7, 35, 12, 9})
	expected := meanMedianMode{
		mean:   126,
		median: 12,
		mode:   21,
	}

	if results.mean != expected.mean {
		t.Errorf("Mean is incorrect, got: %v, want: %v.", results.mean, expected.mean)
	}
	if results.median != expected.median {
		t.Errorf("Median int is incorrect, got: %v, want: %v.", results.median, expected.median)
	}
	if results.mode != expected.mode {
		t.Errorf("Mode int is incorrect, got: %d, want: %d.", results.mode, expected.mode)
	}
}
