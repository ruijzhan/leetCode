package offer

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		if m[n] {
			return n
		}
		m[n] = true
	}
	return -1
}
