package walletcombo

import (
	"testing"
)

func TestWalletCombo_Success(t *testing.T) {
	menu := []struct {
		value int
		name  string
	}{
		{1, "taco"},
		{2, "burger"},
		{3, "shawarma"},
	}

	wallet := 3

	result := getCombo(menu, wallet)
	expected := [][]string{[]string{"taco", "taco", "taco"}, []string{"taco", "burger"}, []string{"shawarma"}}

	resultLen := len(result)
	expectedLen := len(expected)

	if resultLen != expectedLen {
		t.Errorf("When getting the length of a combo, length is incorrect, got: %v, want: %v.", resultLen, expectedLen)
	}

}
