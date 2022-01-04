package offer

//统计一个数字在排序数组中出现的次数

func search(nums []int, target int) (ans int) {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				for i := mid; i < len(nums) && nums[i] == target; i++ {
					ans++
				}
				return
			} else {
				high = mid
			}
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return
}
