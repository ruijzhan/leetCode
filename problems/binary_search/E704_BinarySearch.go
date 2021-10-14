package binary_search

import "sort"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 递归
func search2(nums []int, target int) int {
	var helper func(int, int) int
	helper = func(low, high int) int {
		if low > high {
			return -1
		}

		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			return helper(low, mid-1)
		} else {
			return helper(mid+1, high)
		}
	}
	return helper(0, len(nums)-1)
}

// 用 sort.Search
func search3(nums []int, target int) int {
	n := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	if n < len(nums) && nums[n] == target {
		return n
	}
	return -1
}

// 变体一：查找第一个值等于给定值的元素
func search4(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// 变体二：查找最后一个值等于给定值的元素
func search5(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == len(nums) || nums[mid+1] != target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

// 变体三：查找第一个大于等于给定值的元素
func search6(nums []int, target int) int {
	low, high := 0, len(nums)
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 变体四：查找最后一个小于等于给定值的元素
func search7(nums []int, target int) int {
	low, high := 0, len(nums)
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] <= target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}
