package offer

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		want := target - num
		if idx, ok := m[want]; ok {
			return []int{nums[idx], nums[i]}
		} else {
			m[num] = i
		}
	}
	return nil
}
