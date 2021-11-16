package array

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	for i1, i2, k := m-1, n-1, m+n-1; k >= 0; k-- {
		if i1 < 0 {
			nums1[k] = nums2[i2]
			i2--
		} else if i2 < 0 {
			nums1[k] = nums1[i1]
			i1--
		} else if nums1[i1] < nums2[i2] {
			nums1[k] = nums2[i2]
			i2--
		} else {
			nums1[k] = nums1[i1]
			i1--
		}
	}
}
