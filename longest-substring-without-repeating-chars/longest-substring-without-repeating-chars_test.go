package longestsubstringwithoutrepeatingchars

import (
	"testing"
)

func TestGetLongestSubstrLengthWithoutRepeats_NormalEarly(t *testing.T) {
	result := getLongestSubstrLengthWithoutRepeats("abcabcbb")
	expected := 3 // abc

	if result != expected {
		t.Errorf("Length is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestGetLongestSubstrLengthWithoutRepeats_AllRepeats(t *testing.T) {
	result := getLongestSubstrLengthWithoutRepeats("bbbbb")
	expected := 1 // b

	if result != expected {
		t.Errorf("Length is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestGetLongestSubstrLengthWithoutRepeats_NotSubsequence(t *testing.T) {
	result := getLongestSubstrLengthWithoutRepeats("pwwkew")
	expected := 3 // The answer is "wke", with the length of 3. 
	// Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

	if result != expected {
		t.Errorf("Length is incorrect, got: %v, want: %v.", result, expected)
	}
}