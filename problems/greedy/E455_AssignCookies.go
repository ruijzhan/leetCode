package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	// i 是孩子数，j 是饼干数
	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if s[j] >= g[i] {
			i++
			res++
		}
	}
	return res
}
