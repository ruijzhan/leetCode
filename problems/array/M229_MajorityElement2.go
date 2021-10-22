package array

import "sort"

// 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

func majorityElement(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	res := make([]int, 0)
	count := 0
	for slow, fast := 0, 0; fast < n; fast++ {
		if nums[slow] == nums[fast] {
			count++
		} else {
			count = 1
			slow = fast
		}
		if count > n/3 {
			if len(res) == 0 || res[len(res)-1] != nums[slow] {
				res = append(res, nums[slow])
			}
		}
	}
	return res
}

// 哈希
func majorityElement2(nums []int) []int {
	sort.Ints(nums)
	res := make([]int, 0)
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
		if m[n] > len(nums)/3 {
			if len(res) == 0 || res[len(res)-1] != n {
				res = append(res, n)
			}
		}
	}
	return res
}

//TODO: 官方: 摩尔投票法?
