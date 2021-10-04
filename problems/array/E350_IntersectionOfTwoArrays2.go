package array

import "sort"

// 排序，双指针
func intersect(nums1 []int, nums2 []int) []int {
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
			res = append(res, x)
			i++
			j++
		}
	}

	return res
}
