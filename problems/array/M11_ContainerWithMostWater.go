package array

// 暴力穷举 超时
func maxArea2(height []int) int {

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	res := 0
	for i := 0; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			cap := (j - i) * min(height[i], height[j])
			if cap > res {
				res = cap
			}
		}

	}
	return res
}

// 双指针
func maxArea(height []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	res := 0
	for i, j := 0, len(height)-1; i < j; {
		hi, hj := height[i], height[j]
		res = max(res, (j-i)*min(hi, hj))
		if hi > hj {
			j--
		} else {
			i++
		}
	}

	return res
}
