package array

/*
	26题：
		遍历数组
			如果当前元素等于上一个元素
				什么都不做
			如果当前元素不等于上一个元素
			    a[l] = 当前元素
				改变当前元素
				l++

		用双指针:
			实际上只需要找出不一样的元素
			用 nums[fast] == nums[fast-1] 找出
			统计唯一元素的个数(slow)
*/

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	slow := 1

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
