package binary_search

func mySqrt(x int) int {
	low, high := 0, x
	ans := -1
	for low <= high {
		mid := low + (high-low)>>1
		if mid*mid <= x {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
