package binary_search

func findMin2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] < nums[low] {
			high = mid
		} else {
			high--
		}
	}
	return nums[low]
}
