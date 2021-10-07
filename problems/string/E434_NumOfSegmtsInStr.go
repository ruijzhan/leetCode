package string

import "unicode"

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	prevSpace := true
	for _, b := range s {
		if unicode.IsSpace(b) {
			prevSpace = true
		} else {
			if prevSpace {
				count++
			}
			prevSpace = false
		}
	}
	return count
}

// 官方方法
func countSegments2(s string) int {
	count := 0
	for i, c := range s {
		if (i == 0 || s[i-1] == ' ') && c != ' ' {
			count++
		}
	}
	return count
}
