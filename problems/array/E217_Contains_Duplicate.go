package array

import (
	"sort"
)

// 哈希方法 复杂度为 O(N)
func containsDuplicate(nums []int) bool {

	m := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = struct{}{}
	}
	return false

}

// 排序方法，复杂度为 O(NlogN)，即为排序算法的复杂度
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}

// 递归
func containsDuplicate3(nums []int) bool {
	var hasDup func([]int) bool
	hasDup = func(nums []int) bool {
		if len(nums) <= 1 {
			return false
		}
		if nums[0] == nums[1] {
			return true
		}
		return false || hasDup(nums[1:])
	}

	sort.Ints(nums)
	return hasDup(nums)
}
