package numbertoroman

import "testing"

func TestFromInt_Success(t *testing.T) {
	result1 := fromInt(1)
	result2 := fromInt(2)
	result3 := fromInt(4)
	result4 := fromInt(5)
	result5 := fromInt(1993)
	result6 := fromInt(2018)
	result7 := fromInt(1111)
	result8 := fromInt(2222)
	result9 := fromInt(444)
	result10 := fromInt(555)
	result11 := fromInt(666)
	result12 := fromInt(999)

	if result1 != "I" {
		t.Errorf("When converting num to roman numberal, 1, is incorrect, got: %v, want: %v.", result1, "I")
	}

	if result2 != "II" {
		t.Errorf("When converting num to roman numberal, 2, is incorrect, got: %v, want: %v.", result1, "II")
	}

	if result3 != "IV" {
		t.Errorf("When converting num to roman numberal, 3, is incorrect, got: %v, want: %v.", result1, "IV")
	}

	if result4 != "V" {
		t.Errorf("When converting num to roman numberal, 5, is incorrect, got: %v, want: %v.", result1, "V")
	}

	if result5 != "MCMXCIII" {
		t.Errorf("When converting num to roman numberal, 1993, is incorrect, got: %v, want: %v.", result1, "MCMXCIII")
	}

	if result6 != "MMXVIII" {
		t.Errorf("When converting num to roman numberal, 2018, is incorrect, got: %v, want: %v.", result1, "MMXVIII")
	}

	if result7 != "MCXI" {
		t.Errorf("When converting num to roman numberal, 1111, is incorrect, got: %v, want: %v.", result1, "MCXI")
	}

	if result8 != "2222" {
		t.Errorf("When converting num to roman numberal, 2222, is incorrect, got: %v, want: %v.", result1, "2222")
	}

	if result9 != "CDXLIV" {
		t.Errorf("When converting num to roman numberal, 444, is incorrect, got: %v, want: %v.", result1, "CDXLIV")
	}

	if result10 != "DLV" {
		t.Errorf("When converting num to roman numberal, 555, is incorrect, got: %v, want: %v.", result1, "DLV")
	}

	if result11 != "DCLXVI" {
		t.Errorf("When converting num to roman numberal, 666, is incorrect, got: %v, want: %v.", result1, "DCLXVI")
	}

	if result12 != "CMXCIX" {
		t.Errorf("When converting num to roman numberal, 999, is incorrect, got: %v, want: %v.", result1, "CMXCIX")
	}
}

func TestFromInt_Zero(t *testing.T) {
	zero := fromInt(0)

	if zero != "0" {
		t.Errorf("When converting num to roman numberal, 0, is incorrect, got: %v, want: %v.", zero, "0")
	}
}
