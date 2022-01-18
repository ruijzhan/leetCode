package offer

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		xy, _ := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(y))
		yx, _ := strconv.Atoi(strconv.Itoa(y) + strconv.Itoa(x))
		return xy < yx
	})
	ans := ""
	for _, n := range nums {
		ans = ans + strconv.Itoa(n)
	}
	return ans
}
