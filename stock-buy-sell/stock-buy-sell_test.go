package stockbuysell

import (
	"testing"
)

func TestGetDates(t *testing.T) {
	result := getDates([]int{110, 180, 260, 40, 310, 535, 695})
	expected := "buy on day 4, sell on day 7"

	if result != expected {
		t.Errorf("When getting a days when to buy and sell, result is not what is expected, got: %v, expected: %v.", result, expected)
	}
}
