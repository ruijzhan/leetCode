package array

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i, first := range nums {
		if i == len(nums)-1 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			second, third := nums[j], nums[k]
			sum := first + second + third
			if sum == 0 {
				res = append(res, []int{first, second, third})
				for j < len(nums) && second == nums[j] {
					j++
				}
				for k > i && third == nums[k] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
