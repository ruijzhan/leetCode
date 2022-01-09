package offer

import "strconv"

func translateNum(num int) int {
	sNum := strconv.Itoa(num)
	// dp[i]: 字符串 s[0:i] 能翻译成小写字母的种数
	dp := make([]int, len(sNum)+1)
	dp[0], dp[1] = 1, 1
	for i := 0; i < len(sNum); i++ {
		// 当前数字至少有 1 种翻译方法
		dp[i+1] = dp[i]
		if i == 0 {
			continue
		}
		pre := sNum[i-1 : i+1]
		// 当前数字和前一个数字又可以构成一种翻译方法
		if pre <= "25" && pre >= "10" {
			dp[i+1] = dp[i] + dp[i-1]
		}

	}

	return dp[len(dp)-1]
}
