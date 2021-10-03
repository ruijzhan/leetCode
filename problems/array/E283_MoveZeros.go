package array

// 递归
func moveZeroes(nums []int) {
	l := len(nums)
	if l == 0 {
		return
	}

	moveZeroes(nums[1:])
	if nums[0] == 0 {
		copy(nums[0:], nums[1:])
		nums[l-1] = 0
	}
}

func moveZeroes2(nums []int) {
	for a, b := 0, 0; b < len(nums); b++ {
		if nums[b] != 0 {
			nums[a], nums[b] = nums[b], nums[a]
			a++
		}
	}
}
