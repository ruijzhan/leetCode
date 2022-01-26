package offer

import (
	"math"
	"math/rand"
	"sort"
)

// 哈希
func majorityElement(nums []int) int {
	n := len(nums) / 2
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > n {
			return num
		}
	}
	return math.MinInt64
}

// 排序
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 随机
func majorityElement3(nums []int) int {
	for {
		n := nums[rand.Intn(len(nums))]
		count := 0
		for _, num := range nums {
			if num == n {
				count++
				if count > len(nums)/2 {
					return n
				}
			}
		}
	}
}

//TODO 分治
