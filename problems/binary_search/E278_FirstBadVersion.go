package binary_search

var isBadVersion func(int) bool

func firstBadVersion(n int) int {
	low, high := 0, n
	for low <= high {
		mid := low + (high-low)>>1
		if isBadVersion(mid) {
			if mid == 0 || !isBadVersion(mid-1) {
				return mid
			} else {
				high = mid
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}
