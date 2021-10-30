package array

import "sort"

// 哈希
func singleNumber(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := m[n]; ok {
			delete(m, n)
		} else {
			m[n] = struct{}{}
		}
	}
	for k := range m {
		return k
	}
	return -1
}

// 排序
func singleNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}

// 集合
func singleNumber3(nums []int) int {
	set := make(map[int]struct{})
	sum, sumSet := 0, 0
	for _, n := range nums {
		set[n] = struct{}{}
		sum += n
	}
	for k := range set {
		sumSet += k
	}
	return 2*sumSet - sum
}

// 位运算
func singleNumber4(nums []int) (single int) {
	for _, num := range nums {
		single ^= num
	}
	return
}
