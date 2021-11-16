package array

func maxProfit(prices []int) int {
	lowest := prices[0]
	max := 0
	for _, price := range prices {
		if price < lowest {
			lowest = price
		}
		profix := price - lowest
		if profix > max {
			max = profix
		}
	}
	return max
}

// 暴力
func maxProfit1(prices []int) int {
	max := 0
	for iBuy, buy := range prices {
		for _, sell := range prices[iBuy:] {
			profit := sell - buy
			if profit > max {
				max = profit
			}
		}
	}
	return max
}
