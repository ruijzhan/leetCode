package binary_search

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := left + (right-left)>>1
		sq := mid * mid
		if sq == num {
			return true
		} else if sq > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
