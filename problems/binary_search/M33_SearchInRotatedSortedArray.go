package binary_search

func searchInRotatedSortedArray(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] {
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// 递归
func searchInRotatedSortedArray2(nums []int, target int) int {
	var helper func(int, int) int
	helper = func(low, high int) int {
		if low > high {
			return -1
		}
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] {
			if target >= nums[low] && target < nums[mid] {
				return helper(low, mid-1)
			} else {
				return helper(mid+1, high)
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				return helper(mid+1, high)
			} else {
				return helper(low, mid-1)
			}
		}
	}
	return helper(0, len(nums)-1)
}
