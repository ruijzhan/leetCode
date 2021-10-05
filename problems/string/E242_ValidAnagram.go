package string

import "sort"

// 哈希
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)

	for i, ss := range []byte(s) {
		tt := t[i]
		m[ss]++
		m[tt]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

// 排序
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss, tt := []byte(s), []byte(t)
	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
	sort.Slice(tt, func(i, j int) bool { return tt[i] < tt[j] })
	return string(ss) == string(tt)
}
