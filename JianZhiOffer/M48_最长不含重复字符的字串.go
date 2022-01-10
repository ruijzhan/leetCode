package offer

func lengthOfLongestSubstring(s string) (ans int) {
	dict := make(map[byte]int)
	tmp := 0
	for j := 0; j < len(s); j++ {
		i := -1
		if ii, ok := dict[s[j]]; ok { // 获取索引 i
			i = ii
		}
		dict[s[j]] = j // 更新 hash 表
		if tmp < j-i { // dp[j-1] -> dp[j]
			tmp++
		} else {
			tmp = j - i
		}

		ans = func(a, b int) int {
			if a > b {
				return a
			} else {
				return b
			}
		}(ans, tmp)

	}
	return
}

//todo: 完全没搞懂
