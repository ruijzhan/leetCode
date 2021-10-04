package array

import "sort"

// 哈希
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, n := range nums1 {
		m[n] = struct{}{}
	}
	res := make([]int, 0)

	for _, n := range nums2 {
		if _, has := m[n]; has {
			res = append(res, n)
			delete(m, n)
		}
	}

	return res
}

// 排序, 双指针
func intersection2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x < y {
			i++
		} else if x > y {
			j++
		} else {
			if len(res) == 0 || res[len(res)-1] < x {
				res = append(res, x)
			}
			i++
			j++
		}
	}

	return res
}
