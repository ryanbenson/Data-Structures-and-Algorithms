package validparanthesis

func isValid(sequence string) bool {
	// empty sequences are valid
	if len(sequence) == 0 {
		return true
	}

	paranthesisMap := make(map[string]string)
	paranthesisMap["{"] = "}"
	paranthesisMap["("] = ")"
	paranthesisMap["["] = "]"

	// hold the sequence of characters we're expecting as we move along
	sequenceArr := []string{}

	for i, characterRune := range sequence {
		character := string(characterRune)

		// if we're at the beginning, can move along without checking its closing partner
		// or if we have nothing in our array yet
		if i == 0 || len(sequenceArr) == 0 {
			// if it's not a valid starting character, short circuit since it's an invalid sequence
			expectedCloseChar, ok := paranthesisMap[character]
			if ok == false {
				return false
			}
			sequenceArr = append(sequenceArr, expectedCloseChar)
			continue
		}

		// are we adding to our list?
		expectedCloseChar, ok := paranthesisMap[character]
		if ok == true {
			sequenceArr = append(sequenceArr, expectedCloseChar)
			continue
		}

		// since it's a closing character, it should match the last item if it's valid
		lastCharacterExpected := sequenceArr[len(sequenceArr)-1]
		if character != lastCharacterExpected {
			return false
		}

		// since it's valid, pop the last item off
		sequenceArr = sequenceArr[:len(sequenceArr)-1]
	}

	return true
}
