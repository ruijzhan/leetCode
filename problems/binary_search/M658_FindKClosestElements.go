package binary_search

func findClosestElements(arr []int, k int, x int) []int {
	low, mid, high := 0, -1, len(arr)-1
	for low <= high {
		mid = low + (high-low)>>1
		if arr[mid] == x || mid == 0 || mid == len(arr)-1 {
			break
		}
		if low == mid && high == mid {
			break
		}
		if arr[mid] > x {
			if arr[mid-1] < x {
				break
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	low, high = mid, mid
	for high-low < k {
		if low == 0 {
			high++
		} else if high == len(arr) {
			low--
		} else {
			if abs(arr[low-1]-x) <= abs(arr[high]-x) {
				low--
			} else {
				high++
			}
		}

	}
	return arr[low:high]
}
