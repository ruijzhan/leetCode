package string

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for l, r := 0, len(s)-1; l < r; {
		if !unicode.IsDigit(rune(s[l])) && !unicode.IsLetter(rune(s[l])) {
			l++
		} else if !unicode.IsDigit(rune(s[r])) && !unicode.IsLetter(rune(s[r])) {
			r--
		} else if s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}
