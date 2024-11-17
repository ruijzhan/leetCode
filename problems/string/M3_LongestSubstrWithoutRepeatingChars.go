package string

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	l, res, slow := len(s), 0, 0

	for fast := 0; fast < l; fast++ {
		cFast := s[fast]
		if i, ok := mp[cFast]; ok {
			slow = max(slow, i+1)
		}
		mp[cFast] = fast
		res = max(res, fast-slow+1)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring2(s string) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	mp := make(map[byte]int)

	l, res := len(s), 0

	for slow, fast := 0, 0; fast < l; fast++ {
		cFast := s[fast]
		if i, ok := mp[cFast]; ok {
			slow = max(i, slow)
		}
		res = max(res, fast-slow+1)
		mp[cFast] = fast + 1
	}

	return res
}
