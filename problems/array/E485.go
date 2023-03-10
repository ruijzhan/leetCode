// https://leetcode.cn/problems/max-consecutive-ones/

package array

func findMaxConsecutiveOnes(nums []int) int {
	m, max := 0, 0

	for _, n := range nums {
		if n == 1 {
			m++
		}
		if m > max {
			max = m
		}
		if n == 0 {
			m = 0
		}
	}

	return max
}
