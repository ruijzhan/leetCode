package binary_search

var guess func(int) int

func guessNumber(n int) int {
	low, high := 1, n
	for low < high {
		mid := low + (high-low)>>1
		if guess(mid) <= 0 {
			high = mid
		} else {
			low = mid + 1
		}

	}
	return low
}
