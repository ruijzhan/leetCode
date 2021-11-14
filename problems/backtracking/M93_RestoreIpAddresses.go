package bt

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}

	valid := func(num string) bool {
		if num[0] == '0' {
			return len(num) == 1
		}
		n, _ := strconv.Atoi(num)
		return n <= 255
	}
	var helper func(string, []string)
	helper = func(remain string, nums []string) {
		if len(nums) == 4 {
			if len(remain) == 0 {
				ans = append(ans, strings.Join(nums, "."))
			}
			return
		}
		for i := 1; i < 4 && i <= len(remain); i++ {
			num := remain[:i]
			if valid(num) {
				nums2 := nums[:]
				nums2 = append(nums2, num)
				helper(remain[i:], nums2)
			} else {
				break
			}
		}
	}

	helper(s, []string{})
	return ans
}
