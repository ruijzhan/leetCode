package string

// 暴力匹配
func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
outer:
	for i := 0; i <= m-n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

// TODO: KMP 算法
