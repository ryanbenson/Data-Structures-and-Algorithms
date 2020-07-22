package numbertoroman

func fromInt(number int) string {
	if number <= 0 {
		return ""
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	romanNumeral := ""
	// go through the conversions and decrease from the value from the number
	// until there's no more numbers left (e.g. hit 0)
	for _, conversion := range conversions {
		for number >= conversion.value {
			romanNumeral += conversion.digit
			number -= conversion.value
		}
	}
	return romanNumeral
}
