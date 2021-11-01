package string

func lengthOfLastWord(s string) int {
	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if l == 0 {
				continue
			}
			break
		} else {
			l++
		}
	}
	return l
}
