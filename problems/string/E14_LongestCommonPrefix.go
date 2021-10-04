package string

import "strings"

// 横向扫描
func longestCommonPrefix(strs []string) string {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	lcp := func(s1, s2 string) string {
		builder := strings.Builder{}
		for i := 0; i < min(len(s1), len(s2)); i++ {
			if s1[i] == s2[i] {
				builder.WriteByte(s1[i])
			} else {
				break
			}
		}
		return builder.String()
	}

	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

// 纵向扫描
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	builder := strings.Builder{}
	for i := 0; i < len(strs[0]); i++ {
		var char byte
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return builder.String()
			}
			curr := strs[j][i]
			if char == 0 {
				char = strs[j][i]
			} else if char != curr {
				return builder.String()
			}
		}
		builder.WriteByte(char)
	}

	return builder.String()
}

// TODO: 分治

// TODO: 二分查找
