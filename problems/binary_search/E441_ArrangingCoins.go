package binary_search

import "sort"

func arrangeCoins(n int) int {
	low, high := 1, n
	for low < high {
		mid := low + (high-low)>>1 + 1
		if mid*(mid+1) <= 2*n {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func arrangeCoins2(n int) int {
	return sort.Search(n, func(k int) bool { k++; return k*(k+1) > 2*n })
}

//TODO: 数学方法
