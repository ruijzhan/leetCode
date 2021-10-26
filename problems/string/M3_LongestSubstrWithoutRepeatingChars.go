package string

func lengthOfLongestSubstring(s string) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	mp := make(map[byte]struct{})
	l, res, slow, fast := len(s), 0, 0, 0
	for slow < l && fast < l {
		cFast := s[fast]
		if _, ok := mp[cFast]; !ok {
			mp[cFast] = struct{}{}
			fast++
			res = max(res, fast-slow)
		} else {
			delete(mp, s[slow])
			slow++
		}
	}

	return res
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
