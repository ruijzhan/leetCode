package offer

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	mp := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		profit := prices[i] - minPrice
		if profit > mp {
			mp = profit
		}

	}
	return mp
}
