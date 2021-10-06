package array

import (
	"math"
	"sort"
)

// 若第三大的数不存在,则返回最大的数

// 排序
func thirdMax(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}

	for i, diff := 1, 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			diff++
		}
		if diff == 3 {
			return nums[i]
		}
	}
	return nums[0]
}

// 一次遍历
func thirdMax2(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, n := range nums {
		if n > a {
			a, b, c = n, a, b
		} else if n < a && n > b {
			b, c = n, b
		} else if n < b && n > c {
			c = n
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
