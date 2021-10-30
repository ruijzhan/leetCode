package array

import "sort"

func singleNumber6(nums []int) []int {
	m := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := m[n]; ok {
			delete(m, n)
		} else {
			m[n] = struct{}{}
		}
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

//TODO: 位运算
