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
	return searchHelper(nums, 0, len(nums)-1, target)
}

func searchHelper(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)>>1
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return searchHelper(nums, low, mid-1, target)
	} else {
		return searchHelper(nums, mid+1, high, target)
	}
}

// 用 sort.Search
func search3(nums []int, target int) int {
	n := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	if n < len(nums) && nums[n] == target {
		return n
	}
	return -1
}
