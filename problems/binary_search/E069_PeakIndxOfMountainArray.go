package binary_search

func peakIndexInMountainArray(arr []int) int {
	low, high := 1, len(arr)-2
	for low <= high {
		mid := low + (high-low)>>1
		l, m, r := arr[mid-1], arr[mid], arr[mid+1]
		if m > l && m > r {
			return mid
		}
		if m < l {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 递归
func peakIndexInMountainArray2(arr []int) int {
	var helper func(int, int) int
	helper = func(low, high int) int {
		if low > high {
			return -1
		}
		mid := low + (high-low)>>1
		l, m, r := arr[mid-1], arr[mid], arr[mid+1]
		if m > l && m > r {
			return mid
		}
		if l > m {
			return helper(low, mid-1)
		} else {
			return helper(mid+1, high)
		}
	}
	return helper(1, len(arr)-2)
}
