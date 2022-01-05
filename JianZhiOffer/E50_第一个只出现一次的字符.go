package offer

func firstUniqChar(s string) byte {
	m := make([]int, 26)
	for _, c := range s {
		m[c-'a']++
	}
	for i, c := range s {
		if m[c-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
