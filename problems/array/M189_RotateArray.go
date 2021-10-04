package array

// 递归
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	end := nums[len(nums)-1]
	copy(nums[1:], nums)
	nums[0] = end

	rotate(nums, k-1)
}

// 额外的数组
// 遍历数组，将下标为 i 的元素放到新数组中 (i+k) mod n 的位置
func rotate2(nums []int, k int) {
	n := len(nums)
	roated := make([]int, n)
	for i, v := range nums {
		roated[(i+k)%n] = v
	}
	copy(nums, roated)
}

// 数组翻转
func rotate3(nums []int, k int) {

	reverse := func(nums []int) {
		n := len(nums)
		for i := 0; i < n/2; i++ {
			nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
		}

	}
	k = k % len(nums)

	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}
