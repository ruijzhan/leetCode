package array

func numberOfPairs(nums []int) []int {
	res := make([]int, 2)

	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	for _, t := range m {
		res[0] += t / 2
		res[1] += t % 2
	}

	return res
}
