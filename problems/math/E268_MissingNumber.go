package math

func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= i
		res ^= nums[i]
	}
	res ^= len(nums)
	return res
}
