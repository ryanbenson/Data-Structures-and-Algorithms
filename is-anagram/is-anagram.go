package isanagram

import (
	"sort"
	"strings"
)

func isAnagram(str1 string, str2 string) bool {
	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str2)
	return sortedStr1 == sortedStr2
}

func isAnagramLoop(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	if str1 == "" && str2 == "" {
		return true
	}

	strLen := len(str1)
	for i := 0; i < strLen; i++ {
		letter := str1[i]
		str1Index := strings.IndexByte(str1, letter)
		str2Index := strings.IndexByte(str2, letter)

		if str1Index == -1 || str2Index == -1 {
			return false
		}

	}
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
