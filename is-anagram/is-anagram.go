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

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
