package onerow

func isOneRowFilter(words []string) []string {
	wordsWithSameRow := make([]string, 0)

	for _, word := range words {
		isOneRow := isWordOneRow(word)
		if isOneRow == true {
			wordsWithSameRow = append(wordsWithSameRow, word)
		}
	}
	return wordsWithSameRow
}

func isWordOneRow(word string) bool {
	keyMap := map[rune]int {
		'q': 1,
		'w': 1,
		'e': 1,
		'r': 1,
		't': 1,
		'y': 1,
		'u': 1,
		'i': 1,
		'o': 1,
		'p': 1,
		'a': 2,
		's': 2,
		'd': 2,
		'f': 2,
		'g': 2,
		'h': 2,
		'j': 2,
		'k': 2,
		'l': 2,
		'z': 3,
		'x': 3,
		'c': 3,
		'v': 3,
		'b': 3,
		'n': 3,
		'm': 3,
	}

	row := 0
	for i, letter := range word {
		if i == 0 {
			rowNum, ok := keyMap[letter]
			if !ok {
				return false
			}
			row = rowNum
			continue
		}
		rowNum, ok := keyMap[letter]
		if !ok || rowNum != row {
			return false
		}
	}
	return true
}