package offer

func maxSubArray(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// dp[i]: 以 nums[i] 结尾的最大的连续子数组的和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		// dp[i] 的状态只跟 dp[i-1] 有关
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}

	return ans
}
