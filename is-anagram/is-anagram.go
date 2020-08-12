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

func isAnagramStringsManip(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	if str1 == "" && str2 == "" {
		return true
	}

	for len(str1) > 0 {
		letter := str1[0]
		str1Index := strings.IndexByte(str1, letter)
		str2Index := strings.IndexByte(str2, letter)

		if str1Index == -1 || str2Index == -1 {
			return false
		}

		str1 = str1[:str1Index] + str1[str1Index+1:]
		str2 = str2[:str2Index] + str2[str2Index+1:]
	}

	if len(str1) == 0 && len(str2) == 0 {
		return true
	}

	return false
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
