package binary_search

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target && (mid == 0 || nums[mid-1] < target) {
			ans[0] = mid
			break
		}
		if nums[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if ans[0] == -1 {
		return ans
	}

	low, high = 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] > target) {
			ans[1] = mid
			break
		}
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}
