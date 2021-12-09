package binary_search

import "math"

func findPeakElement(nums []int) int {
	at := func(i int) int {
		if i < 0 || i >= len(nums) {
			return math.MinInt
		}
		return nums[i]
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if at(mid) > at(mid-1) && at(mid) > at(mid+1) {
			return mid
		}
		if at(mid) < at(mid+1) {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1
}
