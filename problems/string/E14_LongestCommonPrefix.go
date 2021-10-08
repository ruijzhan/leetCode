package string

// 横向扫描
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	lcp := func(s1, s2 string) string {
		prefix := []byte{}
		for i := 0; i < min(len(s1), len(s2)); i++ {
			if s1[i] == s2[i] {
				prefix = append(prefix, s1[i])
			} else {
				break
			}
		}
		return string(prefix)
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

// 纵向扫描
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lcp := []byte{}

	for i := 0; i < len(strs[0]); i++ {
		b := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != b {
				return string(lcp)
			}
		}
		lcp = append(lcp, b)
	}
	return string(lcp)
}

// TODO: 分治

// TODO: 二分查找
