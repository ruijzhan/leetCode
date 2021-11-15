package array

func maxSubArray(nums []int) int {
	l := len(nums)
	max := nums[0]
	for i := 1; i < l; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
